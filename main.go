package main

import "fmt"

func main() {

	// /*****FUNCTIONS*****/
	// hello("dolbaeb")
	// fmt.Print(sum(1, 100, 2))
	// fmt.Print(swap("хуй", "залупа"))
	// fmt.Print(divide(10, 0))
	// fmt.Print(greet("dolbaeb"))

	// /*****POINTERS*****/
	// a, b := 1, 2
	// fmt.Println(a, b)
	// Pswap(&a, &b)
	// fmt.Println(a, b)

	// u := User{Age: 15}
	// fmt.Print(u.Age)
	// birthday(&u)
	// fmt.Print(u.Age)

	// a := 5
	// //ptr := &a
	// var ptr *int
	// isNil(ptr)

	/*****SLICES*****/
	// //1
	// s := []int{1, 2, 3, 4, 5}
	// fmt.Println(s)
	// s = append(s, 6, 7)
	// fmt.Println(s)
	// s = append(s[:2], s[3:]...)
	// fmt.Println(s)

	// //2
	// m := map[string]string{
	// 	"розы":     "roses",
	// 	"паравозы": "locomotives",
	// 	"металл":   "metal",
	// }

	// fmt.Println(m)
	// m["срал"] = "shit"
	// fmt.Println(m)
	// _, key := m["розы"]

	// if key {
	// 	fmt.Println("вы любите розы")
	// } else {
	// 	fmt.Println("а я на них срал")
	// }

	// delete(m, "розы")
	// _, key = m["розы"]

	// if key {
	// 	fmt.Println("вы любите розы")
	// } else {
	// 	fmt.Println("а я на них срал")
	// }

	// //3
	m := map[int]int{}

	for i := 1; i <= 10; i++ {
		m[i] = i * i
	}
	fmt.Print(m)

	//4
	s := []string{"go", "is", "awesome"}

	for _, i := range s {
		fmt.Println(i)
	}

}
