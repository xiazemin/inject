package main
import (
	"fmt"
	"github.com/codegangsta/inject"
)

type SpecialString interface{}
type TestStruct struct {
	Name   string `inject`
	Nick   []byte
	Gender SpecialString `inject`
	uid    int           `inject`
	Age    int           `inject`
}

func main() {
	s := TestStruct{}
	inj := inject.New()
	inj.Map("陈一回")
	inj.MapTo("男", (*SpecialString)(nil))
	inj2 := inject.New()
	inj2.Map(20)
	inj.SetParent(inj2)
	inj.Apply(&s)
	fmt.Println("s.Name =", s.Name)
	fmt.Println("s.Gender =", s.Gender)
	fmt.Println("s.Age =", s.Age)
}