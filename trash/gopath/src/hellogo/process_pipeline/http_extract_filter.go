package process_pipeline

import (
	"errors"
	"net/http"
)

type HttpExtractFilter struct {
}

func (self *HttpExtractFilter) Process(data interface{}) (interface{}, error) {
	in, ok := data.(*http.Request)
	if !ok {
		return nil, errors.New("HttpExtractFilter input type should be http.Request")
	}
	ret := (*in).URL.String()
	return &ret, nil
}

