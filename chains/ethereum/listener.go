package ethereum

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ChainSafe/ChainBridgeV2/chains"
	"github.com/ChainSafe/ChainBridgeV2/constants"
	emitter "github.com/ChainSafe/ChainBridgeV2/contracts/Emitter"
	msg "github.com/ChainSafe/ChainBridgeV2/message"
	"github.com/ChainSafe/log15"
	eth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
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
		sub := subscription
		switch sub {
		case constants.ErcTransferSignature, constants.NftTransferSignature:
			err := l.RegisterEventHandler(sub, l.handleTransferEvent)
			if err != nil {
				log15.Error("failed to register event handler", "err", err)
			}
		case constants.DepositAssetSignature:
			err := l.RegisterEventHandler(sub, l.handleTestDeposit)
			if err != nil {
				log15.Error("failed to register event handler", "err", err)
			}
		default:
			return fmt.Errorf("Unrecognized event: %s", sub)
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
func (l *Listener) RegisterEventHandler(subscription string, handler chains.EvtHandlerFn) error {
	evt := EventSig(subscription)
	query := l.buildQuery(l.cfg.emitter, evt)
	subscriptionEvent, err := l.conn.subscribeToEvent(query)
	if err != nil {
		return err
	}
	l.subscriptions[EventSig(subscription)] = subscriptionEvent
	go l.watchEvent(subscriptionEvent, handler)
	log15.Debug("Registered event handler", "chainID", l.cfg.id, "contract", l.cfg.emitter, "sig", subscription)
	return nil
}

// watchEvent will call the handler for every occurrence of the corresponding event. It should be run in a separate
// goroutine to monitor the subscription channel.
func (l *Listener) watchEvent(subscriptionEvent *Subscription, handler func(interface{}) msg.Message) {
	for {
		select {
		case evt := <-subscriptionEvent.ch:
			m := handler(evt)
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

func (l *Listener) handleTransferEvent(eventI interface{}) msg.Message {
	event := eventI.(ethtypes.Log)

	contractAbi, err := abi.JSON(strings.NewReader(string(emitter.EmitterABI)))
	if err != nil {
		log15.Error("Err", err)
	}

	var NftEvent NftTransfer

	err = contractAbi.Unpack(&NftEvent, "LogFill", event.Data)
	if err != nil {
		log15.Error("err", err)
	}
	fmt.Printf("%+v\n", NftEvent)

	msg := msg.Message{
		Type: msg.CreateDepositProposalType,
		Data: event.Topics[0].Bytes(),
	}
	log15.Info("event", "data", event.Topics)
	msg.EncodeCreateDepositProposalData(big.NewInt(0), l.cfg.chainID)
	return msg
}

func (l *Listener) handleTestDeposit(eventI interface{}) msg.Message {
	event := eventI.(ethtypes.Log)
	data := ethcrypto.Keccak256Hash(event.Topics[0].Bytes()).Bytes()
	return msg.Message{
		Type: msg.DepositAssetType,
		Data: data,
	}
}

type NftTransfer struct {
	DestinationChain *big.Int
	DepositId        *big.Int
	To               common.Address
	TokenAddress     common.Address
	TokenId          *big.Int
	Data             []byte
}
