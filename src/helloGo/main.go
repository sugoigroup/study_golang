package main

import (
	"fmt"
	"math"

	"./strutils"
)

const username = "kim"
const (
	ca = 10
	cd
	da = iota
	ea = "earth"
)

func main() {
	fname, lname := "fuck", "you"
	var name string
	var num int = 5
	var pnum = &num

	fmt.Println(fname, name)
	fmt.Println(strutils.ToUpperCase(lname))
	fmt.Println(pnum)
	fmt.Println(*pnum)

	fmt.Println(add(5, 6))
	fmt.Println(asSlice())

	fmt.Println(ca, cd, da, ea)
	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	results := <-ch
	fmt.Println(results)

	fmt.Println(reMap())
	remap := reMap()
	elem, ok := remap["st122"]
	fmt.Println(elem, ok)

	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	strutils.Add("hohoho")
	fmt.Println(strutils.GetAll())

	arr := [3]string{"go", "is", "awesome"}
	fmt.Println(arr)

}

func add(x, y int) (returnX int, returnB bool) {
	returnX = x + y
	returnB = x > y
	return
}

func asSlice() (returnX []int) {
	b := make([]int, 5)

	fmt.Println(len(b), cap(b))
	p := []int{2, 3, 5, 7, 11, 13}
	for _, value := range p {

		fmt.Println(value)
	}

	return p[1:4]
}

// Vertex is good
type Vertex struct {
	Lat, Long float64
}

func reMap() (ms map[string]Vertex) {
	ms = map[string]Vertex{
		"st1": {40.0, -30.01},
		"st2": {56.42, 244.22},
	}
	return
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.Lat * v.Long)
}
