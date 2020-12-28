package main

import (
	"fmt"
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, r *http.Request) {
	write(writer, "Hello, web!")
}

func frenchHandler(writer http.ResponseWriter, r *http.Request) {
	write(writer, "Salut web!")
}

func hindiHandler(writer http.ResponseWriter, r *http.Request) {
	write(writer, "Namaste, web!")
}

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Byte")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func main() {
	//http.HandleFunc("/hello", englishHandler)
	//http.HandleFunc("/salut", frenchHandler)
	//http.HandleFunc("/namaste", hindiHandler)
	//err := http.ListenAndServe("localhost:8080", nil)
	//log.Fatal(err)

	var greeterFunctino func()
	var mathFunction func(int, int) float64
	greeterFunctino = sayHi
	mathFunction = divide

	greeterFunctino()
	fmt.Println(mathFunction(5, 2))
}
