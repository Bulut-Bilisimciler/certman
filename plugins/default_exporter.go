package plugins

type Exporter interface {
	Export() error
}
