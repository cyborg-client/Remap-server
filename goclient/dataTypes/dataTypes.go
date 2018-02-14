package dataTypes

type TcpDataStream []int
// TODO: This might be a struct instead of int
type TcpHttpClientStatus int

const
(
	Start = iota
	Stop
	ChangeFrequency
)
type ClientRequest struct
{
	Request int
	payload []byte
}