package test

import (
	"fmt"
	"time"
)

type myInterface interface {
	Greeting()
}

type myStruct struct {
	Name string
}

func (m *myStruct) Greeting() {
	fmt.Printf("%v, %sさん、こんにちは！\n", time.Now(), m.Name)
}

func (m *myStruct) Message(s string) time.Time {
	fmt.Printf("%s さん、%s\n", m.Name, s)
	return time.Now()
}

func main() {
	ms := myStruct{"山田"}
	ms.Greeting()

	fmt.Println(ms.Message("さようなら"))
}
