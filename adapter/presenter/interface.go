package presenter

type Presenter interface {
	Bind(interface{}) error
	Show() ([]byte, error)
}
