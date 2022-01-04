package main

import "fmt"

type notifier interface {
	notify()
}

type no struct {
	name string
}

func (n *no) notify() {
	fmt.Printf("notify %s", n.name)
}

func (n no) test() {
	fmt.Printf("notify %s", n.name)
}

func note(n notifier) {
	n.notify()
}

func main() {
	n := no{name: "benjamin"}

	// note(n)
	// cannot use n (variable of type no) as notifier value in argument to note: missing method notify (notify has pointer receiver)

	n.test()
}
