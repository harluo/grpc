package kernel

type Router interface {
	Handlers() []Handler
}
