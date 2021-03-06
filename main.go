package main

import (
	"time"
	"net/http"
	"fmt"
)

func main(){
	links := []string{
		"https://google.com",
		"https://flipkart.com",
		"https://amazon.com",
		"https://stackoverflow.com",
		"https://netflix.com",
	}
	c := make(chan string)
	for _, link := range links{
		go checkLink(link, c)
	}

	// Reading a  message from channel is blocking operation
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	for l := range c{
		// go checkLink(l, c)
		go func (l string)  {
			time.Sleep(time.Second*2)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string){
	if _, err := http.Get(link); err != nil{
		fmt.Println(link, "is down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}