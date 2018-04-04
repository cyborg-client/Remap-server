package tcphttpclient

import "github.com/cyborg-client/client/config"

type TcpDataStream [60 * config.SegmentLength]int32

type Status int

const (
	Start = iota
	Stop
	Stimulate
)

type statusTcp int

const (
	start statusTcp = iota
	stop
)

type startStopTcp int

// TODO: Fill inn proper statuses
const (
	status1 startStopTcp = iota
)
