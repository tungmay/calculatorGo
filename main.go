package main

import(
	"fmt"
	"time"
)
				
func main(){
	
	fmt.Println("hello world", validate("TT", 5))
	fmt.Println("time", time.Now())
}
func validate(input string, number int) int {
	if input == "TT"{
		return 4 * number
	}
	return 0 * number
}