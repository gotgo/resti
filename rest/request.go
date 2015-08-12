package rest

import (
	"io/ioutil"
	"net/http"

	"github.com/gotgo/fw/util"
	"github.com/gotgo/retrace"
)

type Request struct {
	Raw        *http.Request
	Context    *RequestContext
	Definition ServerResource
	Args       interface{} //map?
	Body       interface{}
	//header????
	bodyBytes []byte
}

func NewRequest(raw *http.Request, ctx *RequestContext, spec ServerResource) *Request {
	r := &Request{
		Raw:        raw,
		Context:    ctx,
		Definition: spec,
	}
	return r
}

func (r *Request) ContentType() string {
	if r.Raw != nil {
		ct := r.Raw.Header["Content-Type"]
		if len(ct) > 0 {
			return ct[0]
		}
	}
	return ""
}

func (r *Request) Annotate(f retrace.From, k string, v interface{}) {
	r.Context.Trace.Annotate(f, k, v)
}

// Bytes returns the body of the request as a []byte
func (r *Request) Bytes() ([]byte, error) {
	if r.bodyBytes == nil {
		defer r.Raw.Body.Close()

		if bts, err := ioutil.ReadAll(r.Raw.Body); err != nil {
			return nil, err
		} else {
			r.bodyBytes = bts
		}
	}
	return r.bodyBytes, nil
}

func (r *Request) DecodeArgs(argValues map[string]string) error {
	args := r.Definition.ResourceArgs()
	r.Args = args

	if args != nil && len(argValues) > 0 {
		return util.MapToStruct(argValues, &args)
	}
	return nil
}
