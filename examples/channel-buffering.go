package examples

import "fmt"

func init() {
	// 固定buffer的channel
	messages := make(chan string, 3)

	messages <- "Hello"
	messages <- "Channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
