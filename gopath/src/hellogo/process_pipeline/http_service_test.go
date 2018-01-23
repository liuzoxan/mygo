package process_pipeline

import (
	"net/http"

	"testing"
)

func TestInputDataCheck(t *testing.T) {
	p1 := P1_Filter{}
	if _, err := p1.Process(8); err == nil {
		t.Fail()
	}
	if _, err := p1.Process("Eight"); err == nil {
		t.Fail()
	}
	in := "Eight"
	if _, err := p1.Process(&in); err != nil {
		t.Fail()
	}
}

func TestHttpService(t *testing.T) {
	stdHandle, err := CreateHandler(STANDARD)
	revHandle, err := CreateHandler(REVERSE)
	if err == nil {
		http.Handle("/std", &stdHandle)
		http.Handle("/rev", &revHandle)
		http.ListenAndServe(":8213", nil)
	}
}

