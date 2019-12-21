package main

import (
	"fmt"
	"github.com/jteng/distributed-cluster/pub/Address"
)
func main() {
	fmt.Print("dfdfdf")
	a := Address{
		City:"asdfaf"
	}
	fmt.Printf("asdf is %s",a.FullAddress())
}
