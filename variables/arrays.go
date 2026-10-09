// in this file, we learn about arrays 

package main 

import (
	"fmt"
	"slices"
)

func main() {

	var a [5] int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	b := [5] int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3] int

	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)

	two2D := [2][3] int{
		{1, 2,0},
		{0,9,3},
	}
	fmt.Println("2D:", two2D)
	



// slices are quite different from arrays

var s []string 
fmt.Println("uninit:", s, s == nil, len(s) == 0)


s = make([] string, 3)
fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))


s[0] = "a"
s[1] = "b"
s[2] = "c"
fmt.Println("set:", s)
fmt.Println("get:", s[2])

fmt.Println("length:", len(s))

s = append(s, "d")
s = append(s, "e", "f")
fmt.Println("apd:", s)


c := make([] string, len(s))
copy(c, s)
fmt.Println("cpy:", c)

l := s[2:5]

fmt.Println("sl1:", l)

l = s[:5]
fmt.Println("sl2:", l)

t := []string{"g", "h", "t"}
if slices.Equal(t, c) {
	fmt.Println("they are. equal.")
}


}