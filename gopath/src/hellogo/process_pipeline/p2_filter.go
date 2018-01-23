package process_pipeline

import (
	"errors"
)

type P2_Filter struct {
}

func (self *P2_Filter) Process(data interface{}) (interface{}, error) {
	in, ok := data.(*string)
	if !ok {
		return nil, errors.New("P2_Filter input type should be http.Request")
	}
	ret := *in + "_P2"
	return &ret, nil
}

