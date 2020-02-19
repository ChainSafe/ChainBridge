package msg

type ChainId uint8

const EthereumId = ChainId(0)
const CentrifugeId = ChainId(1)

func (id ChainId) String() string {
	switch id {
	case EthereumId:
		return "ethereum"
	case CentrifugeId:
		return "centrifuge"
	default:
		return "unknown"
	}
}
