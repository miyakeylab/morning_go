package typedata

import "fmt"

type Qdata struct {
	name string
	age  int
}

func Typedata() {

	p := struct {
		name string
		age  int
	}{
		name: "Gopher",
		age:  10}

	fmt.Print(p.name)

}
