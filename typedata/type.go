package typedata

import "fmt"

func Typedata() {

	p := struct {
		name string
		age  int
	}{
		name: "Gopher",
		age:  10}

	fmt.Print(p.name)

}
