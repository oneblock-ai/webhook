package config

// Options for the webhook server
type Options struct {
	Namespace       string
	Threadiness     int
	HTTPSListenPort int
	DevMode         bool
	DevURL          string

	ControllerUsername        string
	GarbageCollectionUsername string
}
