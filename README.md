# shortcut

command line

go run main.go

**pack** = package main

**fmain** = func main

**pk** = magic

**pkgm**

**fp** = fmt.Println("")

log.Println("Hello, from log")  
2019/09/17 08:30:34 Hello, from log

var gopher string  
gopher = "Gopher"  
fmt.Printf("Hello, %s.\n", gopher)

**ff** for shortcut fmt.Printf("")

var name = "Sutthinart Khunvadhana"  
fmt.Printf("My name is **%s**.\n", name)

var name = "Sutthinart Khunvadhana" (can be reduced to)  
**name := "Sutthinart Khunvadhana"**

how to declare var but will not use  
**_ = name**

**for** = for index := 0; index < count; index++

**forr** = for _, var := range var

**tys** = type name struct {}

**tyi** = type name interface {}

**iferr** = if err != nil {}

**hand** = func (w http.ResponseWriter, r *http.Request) {}

r.URL.Path vs r.RequestURI  
when request = <http://localhost:8080/about?name=hello>  
r.RequestURI = /about?name=hello  
r.URL.Path = /about  
try **r.URL.Query()**
