package chaolesson

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
	"unsafe"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(3)
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
}

type intFn func(op int) int

func timeSpent(inner intFn) intFn {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 2)
	return op
}
func TestTimeSpent(t *testing.T) {
	fn := timeSpent(slowFunc)
	result := fn(1)
	t.Log("result is", result)
}

func logFilter(inner func(op string) bool) func(op string) bool {
	return func(op string) bool {
		if op != "NOOP" {
			log.Print("action: ", op)
		}
		return true
	}
}

func logMethod(op string) bool {
	if op == "NOOP" {
		return false
	}
	return true
}

func TestFilter(t *testing.T) {
	newFn := logFilter(logMethod)
	newFn("one")
	newFn("two")
	newFn("NOOP")
	newFn("three")

}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarVariable(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
	t.Log(Sum(2, 3, 4, 5))
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("clear resource")
	}()

	t.Log("server start")
	panic("error")
}

type Employee struct {
	Id   int
	Name string
}

func (e Employee) Hello() {
	fmt.Printf("id %x\n", unsafe.Pointer(&e.Id))
	fmt.Printf("name %x\n", e.Name)
}

func (e *Employee) World() {
	fmt.Printf("id %x\n", unsafe.Pointer(&e.Id))
	fmt.Printf("name %x\n", e.Name)
}

func TestEmployee(t *testing.T) {
	em := Employee{Id: 1, Name: "leon"}
	fmt.Printf("orign id %x\n", unsafe.Pointer(&em.Id))
	em.Hello()
	em.World()
}

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (gp *GoProgrammer) WriteHelloWorld() Code {
	return Code("hello world")
}

func TestProgrammer(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Printf("%s\n", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("wang!")
}

func TestPet(t *testing.T) {
	var d *Dog = new(Dog)
	d.Speak()
	d.SpeakTo("name")
}

type Pet2 interface {
	Hello()
}

type Host struct {
	pet Pet2
}

type Cat struct {
}

func (c *Cat) Hello() {
	fmt.Println("hello cat")
}

func TestCat(t *testing.T) {
	host := Host{
	}
	host.pet = &Cat{}

}
