package handling

import "github.com/gotgo/resti/rest"

func AnonymousHandler(handler rest.HandlerFunc) func(*rest.Request, rest.Responder) {
	return handler
}
