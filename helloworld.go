package main
import (
  "fmt"
  "time"
  "math/rand"
)

type MessageDelay struct {
  amount int
}

func (md MessageDelay) Amount() int {
  if md.amount > 0 {
    return rand.Intn(md.amount)
  } else {
    return 0
  }
}

func hello() string {
  return "hello"
}

func bang() string {
  return "!"
}

func world() string {
  return " world"
}

func message(m string, c chan string, delay MessageDelay) {
  time.Sleep(time.Duration(delay.Amount()) * time.Millisecond)
  c <- m
}

func seedTheRandomNumberGenerator() {
  //rand.Seed( time.Now().UnixNano())
  rand.Seed( time.Now().UTC().UnixNano())
}

func main() {
  seedTheRandomNumberGenerator()
  c := make(chan string)
  go message(hello(), c, MessageDelay{amount: 2})
  go message(world(), c, MessageDelay{amount: 1})
  go message(bang(), c, MessageDelay{}) // ints default to 0

  fmt.Println(<-c, <-c, <-c)
}
