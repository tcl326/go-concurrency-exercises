//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(res chan<- *Tweet, stream Stream) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(res)
			return
		}

		res <- tweet
	}
}

func consumer(res <-chan *Tweet) {
	for t := range res {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	res := make(chan *Tweet)

	// Producer
	go producer(res, stream)

	// Consumer
	consumer(res)

	fmt.Printf("Process took %s\n", time.Since(start))
}
