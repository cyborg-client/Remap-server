package datatypes

type ClientRequest struct {
	Request int
	Options map[string]int
	Payload []byte
}
