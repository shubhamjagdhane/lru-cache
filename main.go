package main

import "fmt"

func main() {
	c := newCache(3)

	c.Set(10, "a")
	c.Set(20, "b")
	c.Set(30, "c")
	c.Set(40, "d")
	c.Set(50, "e")
	c.Set(60, "f")

	c.Get(50)

	fmt.Printf("cache = %+v\n\n", c.records)
	fmt.Println("list: ", c.dll)
}
