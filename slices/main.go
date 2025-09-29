package main

import (
	"fmt"
)

func main(){
	/* the purpose of this project is to practice slices. I was trying to make something with slices in the todo project, but was confused.
	I'm taking a detour here to try and understand the datatype before moving on with the project.*/
	fmt.Println(" ** SLICES ** ")
	fmt.Println("First, let's just make a basic slice.")
	var s []string // init a string array.
	 fmt.Println("uninit:", s, s == nil, len(s) == 0) // this is kind of fancy. it comes from https://gobyexample.com/slices
	 // what the above is doing is printing the value of s, which is nil, then the value of s == nil, which is true.
	// then it calls the len() function which determines the length of s which is 0. So 0 == 0 is true.
	s = make([]string,3)
	fmt.Println("empty set:", s, "len:", len(s), "cap:", cap(s)) // this line will be the set, the length, the capacity.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("the entire set:", s)
	fmt.Println("the index 0 of s:", s[0])
	fmt.Println("the index 1 of s:", s[1])
	fmt.Println("the index 2 of s:", s[2])
	fmt.Println("the length of s:", len(s))

	s = append(s, "d")
    s = append(s, "e", "f")
	s = append(s, "g")
    fmt.Println("appended letters:", s)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied slice, c:", len(c))
	fmt.Println("original copy, s:", len(s))
	fmt.Println("contents of c:", c)
	fmt.Println("contents of s:", s)
	fmt.Println("contents of c[3]:", c[3])
	fmt.Println("contents of s[3]:", s[3])

	l := s[2:5]
	fmt.Println("l is the slice from s[2] to s[5].")
	fmt.Println("here is s[2],s[5]then s[2:5]", s[2], s[5],s[2:5])
	fmt.Println("here is l which is s[2:5]", l)

	fmt.Println("Notice how it got c, but not f. it starts at 2 but goes UP TO but not TO f.")
}
