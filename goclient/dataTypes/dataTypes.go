package dataTypes

type TcpHttpClientData []int
// TODO: This might be a struct instead of int
type TcpHttpClientStatus int

const
(
	Start = iota
	Stop
	ChangeFrequency
)
type ClientHttpServerRequest struct
{
	Request int
	payload []byte
}
