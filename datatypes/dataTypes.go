// Package datatypes provides datatypes used in more then two packages
package datatypes

// ClientRequest represents a request from a user, such as startMEA and stopMEA server
type ClientRequest struct {
	Request int
	Options map[string]int
	Payload []byte
}

// Enum defining a global Start and Stop
const (
	Start = iota
	Stop
)