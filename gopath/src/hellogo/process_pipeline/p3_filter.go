package process_pipeline

import (
	"errors"
)

type P3_Filter struct {
}

func (self *P3_Filter) Process(data interface{}) (interface{}, error) {
	in, ok := data.(*string)
	if !ok {
		return nil, errors.New("P3_Filter input type should be http.Request")
	}
	ret := *in + "_P3"
	return &ret, nil
}

