package main

import (
	"errors"
	"fmt"
)

// func main(){

// 	var password string
// 	fmt.Scanln(&password)

// 	if valid, err := validPassword(password); err != nil {
// 		panic(err.Error())
// 	}else{
// 		fmt.Println(valid)
// 	}
// }

// func validPassword(password string) (string, error){
// 	pl := len(password)

// 	if pl <5{
// 		return "", errors.New("password harus lebih dari 4 char")
// 	}

// 	return "Valid password", nil
// }


func main() {
	defer catchErr()

	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func catchErr(){
	if r := recover(); r != nil{
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running")
	}
}

func validPassword(password string) (string, error){
	pl := len(password)

	if pl <5{
		return "", errors.New("password harus lebih dari 4 char")
	}

	return "Valid password", nil
}
