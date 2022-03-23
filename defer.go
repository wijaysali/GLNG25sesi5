// //Defer #1
// package main

// import "fmt"

// func main() {
// 	defer fmt.Println("defer function starts to execute")
// 	fmt.Println("Hai everyone")
// 	fmt.Println("Welcome back to Go learning center")
// }

//Defer #2
// package main

// import "fmt"

// func main() {
// 	callDeferFunc(false)
// 	fmt.Println("Hai everyone!!")
// }

// func callDeferFunc(condition bool) {
// 	if condition == false {
// 		fmt.Println("Don't call deferFunc!!")
// 		return
// 	}
// 	defer deferFunc()
// }

// func deferFunc() {
// 	fmt.Println("Defer func starts to execute")
// }

// func main() {
// 	callDeferFunc()
// 	fmt.Println("Hai everyone!!")
// }

// func callDeferFunc(){
// 	defer deferFunc()
// }

// func deferFunc() {
// 	fmt.Println("Defer func starts to execute")
// }

// OS.Exit
package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}

