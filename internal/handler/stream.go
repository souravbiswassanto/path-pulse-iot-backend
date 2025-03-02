package handler

type StreamHandler interface {
	Receive() (interface{}, error)
	Perform(interface{}) (interface{}, error)
	Send(interface{}) error
}
