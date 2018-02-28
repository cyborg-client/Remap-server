package tcphttpclient

type TcpDataStream []int

type TcpHttpClientStatus int

const
(
	Start = iota
	Stop
	Stimulate
)


type statusTcp int
const
(
	start statusTcp = iota
	stop
)

type startStopTcp int
// TODO: Fill inn proper statuses
const
(
	status1 startStopTcp = iota
)