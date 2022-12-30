package client

type Client interface {
	List() []string
	Update(errCh chan<- error)
}
