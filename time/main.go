package main

import (
	"fmt"
	"time"
)

func main() {

	loc, _ := time.LoadLocation("Asia/Dhaka")
	fmt.Println(time.Now().In(loc))
}
