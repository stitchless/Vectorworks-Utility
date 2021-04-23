package window

import "fmt"

func SendDataToChannel(ch chan int, value int) {
	ch <- value
}

func main() {
	var v int
	ch := make(chan int)

	go SendDataToChannel(ch, 101)

	v = <-ch

	fmt.Println(v)
}