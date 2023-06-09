package starter

import (
	"fmt"
	"strconv"
)

func SayHello(name string) string {
	return fmt.Sprintf("Hello %v. Welcome!", name)
}

func OddOrEven(number int) string {
	oddOrEven := "odd"
	if number%2 == 0 {
		oddOrEven = "even"
	}
	return fmt.Sprintf("%v is an %v number", number, oddOrEven)
}

func FizzBuzz(n int) string {
	result := ""
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result += "FizzBuzz"
		} else if i%3 == 0 {
			result += "Fizz"
		} else if i%5 == 0 {
			result += "Buzz"
		} else {
			result += strconv.Itoa(i)
		}
	}
	return result
}
