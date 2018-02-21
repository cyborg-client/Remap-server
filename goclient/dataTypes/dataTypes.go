package dataTypes

type TcpDataStream []int
// TODO: This might be a struct instead of int
type TcpHttpClientStatus int

const
(
	Start = iota
	Stop
	Stimulate
)

/*
If reqquest = start, post, segment_length and freq

if req = stop, post, no data
 */
type ClientRequest struct
{
	Request int
	Options map[string]int
	payload []byte
}