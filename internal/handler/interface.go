package handler

type StreamHandler interface {
	Receive() (interface{}, error)
	Perform(interface{}) (interface{}, error)
	Send(interface{}) error
}

type LocationProvider interface {
	GetCurrentLocation() (Location, error)
}

type Location interface {
	Latitude() float64
	Longitude() float64
}

type PulseRateProvider interface {
	GetCurrentPulseRate() (PulseRate, error)
}

type PulseRate interface {
	Pulse() float32
}
