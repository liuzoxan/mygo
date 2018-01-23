package process_pipeline

import (
	"errors"
)

type P1_Filter struct {
}

func (self *P1_Filter) Process(data interface{}) (interface{}, error) {
	in, ok := data.(*string)
	if !ok {
		return nil, errors.New("P1_Filter input type should be http.Request")
	}
	ret := *in + "_P1"
	return &ret, nil
}


