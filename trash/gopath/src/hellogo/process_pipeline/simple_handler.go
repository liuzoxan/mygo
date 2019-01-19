package process_pipeline

import (
	"fmt"
	"io"
	"net/http"

	"gitlab.mobvista.com/common/pipefiter_framework/pipefilter"
)

type SimpleHandler struct {
	Pipeline pipefilter.Filter
}

func CreateHandler(pipelineName string) (SimpleHandler, error) {
	pl, err := CreatePipeline(pipelineName)
	if err != nil {
		return SimpleHandler{}, err
	}
	handler := SimpleHandler{
		Pipeline: pl,
	}
	return handler, nil
}

func (self *SimpleHandler) ServeHTTP(c http.ResponseWriter, req *http.Request) {
	ret, err := self.Pipeline.Process(req)
	if err == nil {
		fmt.Println(ret)
	}
	io.WriteString(c, *ret.(*string))
}
