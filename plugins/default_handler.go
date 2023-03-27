package plugins

type Handler interface {
	Handle() error
}
