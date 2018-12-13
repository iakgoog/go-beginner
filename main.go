/*
VSCODE shortcut
option + command + arrow => create multiple cursors
shift + command + enter => create a blank line above cursor then move cursor up
*/
// 'pgkm' shortcut
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// var errAgeTooLow = errors.New("age too low")
// var errAgeTooHigh = errors.New("age too high")
var ( // alternative way of variables declaration
	errAgeTooLow  = errors.New("age too low")
	errAgeTooHigh = errors.New("age too high")
)

func main() {
	tryFatal()
}

func stringPlayground() {
	var gopher string
	gopher = "Gopher"
	fmt.Printf("Hello, %s.\n", gopher)

	var name = "Sutthinart Khunvadhana"
	fmt.Printf("My name is %s.\n", name)

	nickName := "iakgoog"
	fmt.Printf("My nickname is %s.\n", nickName)

	// nickName := // you can't do it again
	_ = nickName // use when you don't want to use this variable
}

func fruitPicker() {
	fmt.Printf("Input a fruit: ")
	var fruit string
	fmt.Scanln(&fruit)

	if fruit == "" {
		fmt.Println("meh! 🤨")
		return
	}

	switch fruit {
	case "apple":
		fmt.Println("🍎")
	case "banana":
		fmt.Println("🍌")
	case "orange":
		fmt.Println("🍊")
	default:
		fmt.Println("🍷")
	}
}

func gradeCalculator() {
	fmt.Printf("Input score: ")
	var score int
	fmt.Scanln(&score)

	if score < 50 {
		fmt.Println("F")
	} else if score < 60 {
		fmt.Println("D")
	} else if score < 70 {
		fmt.Println("C")
	} else if score < 80 {
		fmt.Println("B")
	} else {
		fmt.Println("A")
	}
}

func arrayPlayground() {
	var a [5]int
	a[0] = 10
	a[2] = 30
	a[3] = 40
	fmt.Println(a)
	fmt.Println(len(a))
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("________________")

	for i := range a {
		fmt.Println(a[i])
	}
	fmt.Println("________________")

	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println("________________")

	// 'forr', 'for' shorcut
	b := [3]int{1, 2, 3}
	for _, v := range b {
		fmt.Println(v)
	}
	fmt.Println("________________")

	var c [2][3]int
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			c[i][j] = j
		}
	}
	fmt.Println(c)
	fmt.Println("________________")

	d := make([]int, 5) // declare Slices in Go (dynamically-sized data structure)
	d[0] = 10
	d[2] = 30
	d[3] = 40
	d = append(d, 90)
	fmt.Println(d)
	fmt.Println("________________")

	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(e)
	fmt.Println(e[2:4])
}

func hashPlayground() {
	// declare dictionary
	a := make(map[string]string)
	a["hello"] = "gopher"
	a["name"] = "iakgoog"

	a["x"] = "iakgoog" // try commenting this line and run again
	delete(a, "x")     // try commenting this line and run again

	x, ok := a["x"]
	fmt.Println(ok)
	fmt.Println(x)

	if y, ok := a["x"]; ok {
		// y is only accessible under this scope
		fmt.Println(y)
	}

	for key, value := range a {
		fmt.Println(key, ":", value)
	}
	fmt.Println("________________")

	// another way to declare dictionary
	b := map[string]string{
		"hello": "gopher",
		"name":  "iakgoog",
	}
	fmt.Println(b)
}

func pointerPlayground() {
	a := 10
	fmt.Println(a)

	ptrA := &a
	fmt.Println(ptrA)
	fmt.Println(*ptrA)

	*ptrA = 20
	fmt.Println(a)
}

/*================================ Function ================================*/
// declare function
func add(x int, y int) int { // can omit 'int' like (x, y int)
	p := 1
	return x + y + p
}

func tryAdd() {
	a, b := 1, 2
	r := add(a, b)
	fmt.Println(r)
}

/*================================ Array Mutation ================================*/
func mutateArray(a []int) { // accept Slices as parameter
	a[0] = 10 // only mutate in this function scope
}

func tryMutateArray() {
	a := [5]int{}
	fmt.Println(a)
	mutateArray(a[0:len(a)]) // convert Array to Slices or a[:]
	fmt.Println(a)
	mutateArray(a[1:])
	fmt.Println(a)
}

/*================================ Struct ================================*/
// 'tys' shorcut
type person struct {
	Name     string
	NickName string
}

func tryStruct() {
	p1 := person{Name: "Sutthinart", NickName: "iakgoog"}
	fmt.Println(p1)
	p1.mutate()
	fmt.Println(p1)

	p2 := person{"Akin", "Kinnie"} // not recommended
	fmt.Println(p2)

	p3 := struct {
		Name     string
		NickName string
	}{
		"Akira",
		"Keira",
	}
	fmt.Println(p3)

	var p4 person
	p4.Name = "Ariya"
	p4.NickName = "Fai"
	mutatePerson(&p4)
	fmt.Println(p4)

}

func mutatePerson(p *person) { // take pointer as a parameter
	p.Name = "Hacker"
	fmt.Println("inside mutate:", p)
}

func (p *person) mutate() { // method of Struct person
	p.Name = "Hacker"
	fmt.Println("inside mutate:", p)
}

/*================================ Interface ================================*/
// 'tyi' shortcut
type talkable interface {
	Talk()
}

type minPerson struct {
	Name string
}

func (p minPerson) Talk() {
	fmt.Println("Hello, I'm", p.Name)
}

type cat struct{}

func (cat) Talk() {
	fmt.Println("Nyaa nyaa")
}

type dog struct{}

func (*dog) Talk() {
	fmt.Println("Woof woof")
}

// in Golang, we don't have to declare 'implements', but declare only Talk() function

func talkWith(t talkable) {
	t.Talk()
}

func tryMinPerson() {
	p := minPerson{"Gopher"}
	talkWith(p)
	talkWith(cat{})
	talkWith(&dog{})
}

func checkType(v interface{}) {
	switch p := v.(type) {
	case string:
		fmt.Println("String:", "hello "+p)
	case int:
		fmt.Println("Int:", p+10)
	case bool:
		fmt.Println("Bool:", !p)
	}
}

func testCheckType(v interface{}) {
	p, ok := v.(string) // p, _ := v.(string) is fine
	if ok {
		fmt.Println(p)
	} else {
		fmt.Println("v is not string")
	}
}

func tryCheckType() {
	checkType("Gopher")
	checkType(10)
	checkType(true)

	testCheckType("Gopher")
	testCheckType(10)
	testCheckType(true)
}

func mapWithInterface() {
	a := make(map[interface{}]interface{})
	a[1] = "Hello"
	a["name"] = "Gopher"

	p := minPerson{"TT"}
	a[p] = "iakgoog"
	if x, ok := a[p].(string); ok {
		fmt.Println(x)
	}
}

/*================================ Error Handling ================================*/

func validateAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age too low")
	} else if age > 60 {
		return fmt.Errorf("age too high")
	}
	return nil
}

func tryValidateAge() {
	// fmt.Println(validateAge(70))
	err := validateAge(10)
	if err == errAgeTooLow {
		fmt.Println("CAN NOT ENTER")
		return
	}
	if err == errAgeTooHigh {
		fmt.Println(":D")
		return
	}
	// iferr shortcut
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*================================ Defer ================================*/

func tryDefer() {
	f, err := os.Create("file.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close() // Will do this at the end of function (after WriteString)
	f.WriteString("Hello")
}

/*================================ Fatal ================================*/
func tryFatal() {
	fmt.Println("Start...")
	log.Fatal("fatal error, program can not run")
	fmt.Println("Hello")
}
