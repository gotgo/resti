package handling

import (
	"io"

	"github.com/gotgo/retrace"
)

type ContentTypeDecoder struct {
	ContentType string
	Decode      func(r io.Reader, v interface{}, trace retrace.Tracer) error
}
