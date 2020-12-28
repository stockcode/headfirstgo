package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL string
	Size int
}

func a() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func greeting(myChannel chan string)  {
	myChannel <- "hi"
}

func abc( channel chan string)  {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def( channel chan string)  {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func odd(channel chan int) {
	channel <- 1
	channel <- 3
}

func even(channel chan int) {
	channel <- 2
	channel <- 4
}

func main() {
	pages := make(chan Page)
	urls := []string{"https://baidu.com", "https://bing.com", "https://qq.com"}
	for _, url := range urls {
		go responseSize(url,pages)
	}

	for i:=0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}

	//myChannel := make(chan string)
	//go greeting(myChannel)
	//fmt.Println(<- myChannel)
	//
	//channel1 := make(chan string)
	//channel2 := make(chan string)
	//go abc(channel1)
	//go def(channel2)
	//
	//fmt.Println(<-channel1)
	//fmt.Println(<-channel2)
	//fmt.Println(<-channel1)
	//fmt.Println(<-channel2)
	//fmt.Println(<-channel1)
	//fmt.Println(<-channel2)
	//fmt.Println()

	//channelA := make(chan int)
	//channelB := make(chan int)
	//
	//go odd(channelA)
	//go even(channelB)
	//fmt.Println(<-channelA)
	//fmt.Println(<-channelA)
	//fmt.Println(<-channelB)
	//fmt.Println(<-channelB)
}

func responseSize(url string, channel chan Page)  {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}

