package channel

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	resultChan := make(chan result, 1)
	sumAndMultiply(2, 3, resultChan)

	res := <-resultChan
	fmt.Printf("Sum Value: %d\n", res.sumValue)
	fmt.Printf("Multiply Value: %d\n", res.multiplyValue)
	close(resultChan)
}

type result struct {
	sumValue      int
	multiplyValue int
}

func sumAndMultiply(a, b int, resultChan chan result) {
	sumValue := a + b
	multiplyValue := a * b
	res := result{sumValue: sumValue, multiplyValue: multiplyValue}
	time.Sleep(time.Second * 2)
	resultChan <- res
	return
}
