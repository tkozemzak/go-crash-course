package main

import (
	"fmt"
	"math"

	"github.com/tkozemzak/go_crash_course/03_packages/strutil"
)


func main() {
	fmt.Println(math.Floor(2.563456))
	fmt.Println(math.Ceil(2.563456))
	fmt.Println(math.Sqrt(2.563456))
	fmt.Println(strutil.Reverse("Hello"))
}