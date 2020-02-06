package ethereum

import (
	"github.com/ChainSafe/ChainBridgeV2/chains"
	"github.com/ChainSafe/ChainBridgeV2/constants"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

var _ chains.Listener = &Listener{}

type Subscription struct {
	ch  <-chan ethtypes.Log
	sub eth.Subscription
}

type Listener struct {
	cfg             Config
	conn            *Connection
	subscriptions   map[EventSig]*Subscription
	router          chains.Router
	emitterContract EmitterContract // instance of bound emitter contract
}

func NewListener(conn *Connection, cfg *Config) *Listener {
	return &Listener{
		cfg:           *cfg,
		conn:          conn,
		subscriptions: make(map[EventSig]*Subscription),
	}
}

func (l *Listener) SetEmitterContract(ec EmitterContract) {
	l.emitterContract = ec
}

func (l *Listener) SetRouter(r chains.Router) {
	l.router = r
}

// Start registers all subscriptions provided by the config
func (l *Listener) Start() error {
	log15.Debug("Starting listener...", "chainID", l.cfg.id, "subs", l.cfg.subscriptions)
	for _, subscription := range l.cfg.subscriptions {
		subscription := subscription
		err := l.RegisterEventHandler(subscription)
		if err != nil {
			log15.Error("failed to register event handler", "err", err)
		}
	}
	return nil
}

// buildQuery constructs a query for the contract by hashing sig to get the event topic
// TODO: Start from current block
func (l *Listener) buildQuery(contract common.Address, sig EventSig) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: nil,
		Addresses: []common.Address{contract},
		Topics: [][]common.Hash{
			{sig.GetTopic()},
		},
	}
	return query
}

// RegisterEventHandler creates a subscription for the provided event on the emitter contract.
// Handler will be called for every instance of event.
func (l *Listener) RegisterEventHandler(subscription string) error {
	evt := EventSig(subscription)
	query := l.buildQuery(l.cfg.emitter, evt)
	subscriptionEvent, err := l.conn.subscribeToEvent(query)
	if err != nil {
		return err
	}
	l.subscriptions[EventSig(subscription)] = subscriptionEvent
	go l.watchEvent(subscription, subscriptionEvent)
	log15.Debug("Registered event handler", "chainID", l.cfg.id, "contract", l.cfg.emitter, "sig", subscription)
	return nil
}

// watchEvent will call the handler for every occurrence of the corresponding event. It should be run in a separate
// goroutine to monitor the subscription channel.
func (l *Listener) watchEvent(subscription string, subscriptionEvent *Subscription) {
	for {
		select {
		case evt := <-subscriptionEvent.ch:
			m := handleEvents(subscription, evt)
			err := l.router.Send(m)
			if err != nil {
				log15.Error("subscription error: cannot send message", "sub", subscriptionEvent, "err", err)
			}
		case err := <-subscriptionEvent.sub.Err():
			if err != nil {
				log15.Error("subscription error", "sub", subscriptionEvent, "err", err)
			}
		}
	}
}

// Unsubscribe cancels a subscription for the given event
func (l *Listener) Unsubscribe(sig EventSig) {
	if _, ok := l.subscriptions[sig]; ok {
		l.subscriptions[sig].sub.Unsubscribe()
	}
}

// Stop cancels all subscriptions. Must be called before Connection.Stop().
func (l *Listener) Stop() error {
	for _, sub := range l.subscriptions {
		sub.sub.Unsubscribe()
	}
	return nil
}

func handleEvents(subscription string, evtI interface{}) msg.Message {
	evt := evtI.(ethtypes.Log)
	log15.Trace("Got event!", "type", subscription)
	log15.Trace("EVT", "type", evt)

	var msgType msg.MessageType
	if subscription == constants.ErcTransferSignature ||
		subscription == constants.NftTransferSignature {

		msgType = msg.CreateDepositProposalType
	} else if subscription == constants.DepositAssetSignature {
		msgType = msg.DepositAssetType
	}

	log15.Trace("Data Dump", "topics", evt.Topics)
	log15.Trace("Data Dump", "length", len(evt.Topics[0].Bytes()))

	return msg.Message{
		Type: msgType,
		Data: evt.Topics[0].Bytes(),
	}
}
