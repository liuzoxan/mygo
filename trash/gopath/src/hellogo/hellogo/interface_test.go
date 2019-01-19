package hellogo

import (
	"fmt"
	"testing"
)

//1
type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func(s S) Get()int {
	return s.Age
}

func(s *S) Set(age int) {
	s.Age = age
}

//3
func f(i I){
	i.Set(10)
	fmt.Println(i.Get())
}

func TestInterface(t *testing.T) {
	s := S{}
	f(&s)  //4
}
