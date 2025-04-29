package constant

const (
	HandlerTypeGrpc HandlerType = iota + 1
	HandlerTypeHttp
)

type HandlerType uint8
