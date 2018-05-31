// Are my brother and I siblings? See below to find out.
package main

import "fmt"

type person struct {
	name      string
	age       int
	interests []string
	sibling   *person
}

func main() {
	jonathan := &person{name: "Jonathan", age: 28, interests: []string{"golang", "baseball", "beer"}}
	andy := &person{name: "Andy", age: 37, interests: []string{"css", "banking", "scotch"}}
	AddSibling(jonathan, andy)
	siblings := AreSiblings(jonathan, andy)
	if siblings == true {
		fmt.Printf("%s and %s are siblings", jonathan.name, andy.name)
	} else {
		fmt.Printf("%s and %s are not siblings", jonathan.name, andy.name)
	}
}

func AddSibling(p1, p2 *person) {
	p1.sibling = p2
	p2.sibling = p1
}

func AreSiblings(p1, p2 *person) bool {
	if p1.sibling == p2 && p2.sibling == p1 {
		return true
	} else {
		return false
	}
}
