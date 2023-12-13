package main

import (
	condb "SHOP_LEK/database/condb"
	"fmt"
)

func main() {
	condb.NewEnv()
	New_msg()
}

func New_msg() {
	mag := "some message"

	var msgPointer *string = &mag

	fmt.Println(mag)
	fmt.Println(*msgPointer)

	changeMessage(&mag, "new message 1")
	fmt.Println(mag)

	changeMessage(msgPointer, "new message 2")
	fmt.Println(mag)

	changeMessage(msgPointer, "new message 3")
	fmt.Println(*msgPointer)

	E := mag

	fmt.Println(E)
	count := 5 // You need to define the count value
	for i := 0; i < count; i++ {
		fmt.Println("LEK")
	}
}
func changeMessage(aPointer *string, newMessage string) {
	*aPointer = newMessage
}
