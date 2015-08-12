package rest_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/gotgo/resti/rest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request", func() {

	var request *rest.Request

	BeforeEach(func() {
		def := &rest.ResourceDef{
			ResourceArgs: reflect.TypeOf(rest.IdIntArg{}),
		}
		ct := []string{"application/json"}
		request = rest.NewRequest(&http.Request{}, &rest.RequestContext{}, rest.NewServerResource(def, ct, ct))
	})

	It("should DecodeArgs", func() {
		v := make(map[string]string)
		v["id"] = "9"
		err := request.DecodeArgs(v)
		Expect(err).To(BeNil())

		idArg, ok := request.Args.(*rest.IdIntArg)
		Expect(ok).To(BeTrue())
		Expect(idArg.Id).To(Equal(9))
	})

	It("should get bytes", func() {
		orig := []byte{1, 3, 2, 4}
		request.Raw.Body = ioutil.NopCloser(bytes.NewBuffer(orig))
		result, err := request.Bytes()
		Expect(err).To(BeNil())
		Expect(result).To(Equal(orig))
	})

})
