package basic

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

// var errAgeTooLow = errors.New("age too low")
// var errAgeTooHigh = errors.New("age too high")
var ( // alternative way of variables declaration
	errAgeTooLow  = errors.New("age too low")
	errAgeTooHigh = errors.New("age too high")
)

// StringPlayground function
func StringPlayground() {
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

// FruitPicker function
func FruitPicker() {
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

// GradeCalculator function
func GradeCalculator() {
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

// ArrayPlayground function
func ArrayPlayground() {
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

// HashPlayground function
func HashPlayground() {
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

// PointerPlayground function
func PointerPlayground() {
	a := 10
	fmt.Println(a)

	ptrA := &a
	fmt.Println(ptrA)
	fmt.Println(*ptrA)

	*ptrA = 20
	fmt.Println(a)
}

/*================================ Function ================================*/
func add(x int, y int) int { // can omit 'int' like (x, y int)
	p := 1
	return x + y + p
}

// TryAdd function
func TryAdd() {
	a, b := 1, 2
	r := add(a, b)
	fmt.Println(r)
}

/*================================ Array Mutation ================================*/
func mutateArray(a []int) { // accept Slices as parameter
	a[0] = 10 // only mutate in this function scope
}

// TryMutateArray function
func TryMutateArray() {
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

// TryStruct function
func TryStruct() {
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

func mutatePerson(p *person) { // accept pointer as a parameter
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

// TryMinPerson function
func TryMinPerson() {
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

// TryCheckType function
func TryCheckType() {
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

// TryValidateAge function
func TryValidateAge() {
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

// TryDefer function
func TryDefer() {
	f, err := os.Create("file.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close() // Will do this at the end of function (after WriteString)
	f.WriteString("Hello")
}

/*================================ Fatal ================================*/

// TryFatal function
func TryFatal() {
	fmt.Println("Start...")
	log.Fatal("fatal error, program can not run")
	fmt.Println("Hello")
}

/*================================ Panic ================================*/

// TryPanic function
func TryPanic() {
	fmt.Println("(1) Start...")
	// panic("panic error, program cannot run") // log Error and END go routine
	doSafeWork()
	fmt.Println("(5) Done")
}

func doFailWork() {
	panic("(3) fail")
}

func doSafeWork() {
	defer func() {
		// this would definitely run before panic
		// ***** to recover panic
		if r := recover(); r != nil {
			fmt.Println("(2) work fail:", r) // r = panic("(3) fail")
		}
		fmt.Println("(4) defer")
	}()
	doFailWork()
	fmt.Println("(6) work success")
}

/*================================ GOROUTINES ================================*/
func say(prefix string) {
	for i := 0; i < 10; i++ {
		fmt.Println(prefix, i)
		time.Sleep(time.Second)
	}
}

func normalLoop() {
	fmt.Println("Start")
	say("hello") // this will run time.Sleep 10 times before running say("hi")
	say("hi")
	fmt.Println("Waiting")
	time.Sleep(5 * time.Second)
	fmt.Println("End")
}

func goRoutineLoop() {
	fmt.Println("Start")
	go say("hello") // this will run say("hi") simultaneously
	go say("hi")    // this also run time.Sleep *1 simultaneously
	fmt.Println("Waiting")
	time.Sleep(10 * time.Second) // *1
	fmt.Println("End")
}

// TryGoRoutine function
func TryGoRoutine() {
	// ***** GOROUTINES != THREADS
	// normalLoop()
	goRoutineLoop()
}

/*================================ Channel ================================*/

var (
	arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
)

// TryChannel function
func TryChannel() {
	withGo()
}

func withGo() {
	defer timer()()
	// create channel that send int
	chRes1 := make(chan int)
	chRes2 := make(chan int)
	go func() {
		chRes1 <- sum(arr1)
	}()
	go func() {
		chRes2 <- sum(arr2)
	}()
	fmt.Println("sum arr1:", <-chRes1)
	fmt.Println("sum arr2:", <-chRes2)
}

func withoutGo() {
	// *1 + *2 cane be rewrite as
	defer timer()()
	// f := timer() // *1
	fmt.Println("sum arr1", sum(arr1))
	fmt.Println("sum arr2", sum(arr2))
	// f() // *2
}

func timer() func() {
	t := time.Now()
	return func() {
		diff := time.Now().Sub(t)
		fmt.Println(diff)
	}
}

func sum(a []int) int {
	s := 0
	for _, x := range a {
		s += x
		time.Sleep(time.Millisecond * 200)
	}
	return s
}

/*================================ Select ================================*/

// TrySelect function
func TrySelect() {
	res, err := doWorkWithLimitTime(5 * time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func doVeryLongWork() int {
	time.Sleep(4 * time.Second)
	return 1
}

func doWorkWithLimitTime(d time.Duration) (int, error) {
	ch := make(chan int)
	go func() {
		ch <- doVeryLongWork()
	}()

	select {
	case r := <-ch: // do work until done
		return r, nil
	case <-time.After(d): // stop work at given time limit
		return 0, fmt.Errorf("timeout")
	}
}
