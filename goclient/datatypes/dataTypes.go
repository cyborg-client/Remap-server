package datatypes

type TcpDataStream []int
// TODO: This might be a struct instead of int
type TcpHttpClientStatus int

const
(
	Start = iota
	Stop
	Stimulate
)

type ClientRequest struct
{
	Request int
	Options map[string]int
	payload []byte
}