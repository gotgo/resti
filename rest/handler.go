package rest

type HandlerFunc func(*Request, Responder)

// RestHandler is a marker interface
type Handler interface {
}

type GetHandler interface {
	Get(*Request, Responder)
}

type PostHandler interface {
	Post(*Request, Responder)
}

type PutHandler interface {
	Put(*Request, Responder)
}

type DeleteHandler interface {
	Delete(*Request, Responder)
}

type PatchHandler interface {
	Patch(*Request, Responder)
}

type HeadHandler interface {
	Head(*Request, Responder)
}
