package datatypes


// TODO: This might be a struct instead of int

type ClientRequest struct
{
	Request int
	Options map[string]int
	payload []byte
}