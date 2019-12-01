package msg

type ChainId uint8

const EthereumId = ChainId(0)
const CentrifugeId = ChainId(1)

func (id ChainId) String() string {
	switch id {
	case 0:
		return "ethereum"
	case 1:
		return "centrifuge"
	default:
		return "unknown"
	}
}
