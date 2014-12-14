package main
import "fmt"

func hello(c chan string) {
  c <- "hello"
}

func world(c chan string) {
  c <- " world!"
}

func main() {
  c := make(chan string)
  go hello(c)
  go world(c)

  fmt.Println(<-c, <-c)
}
