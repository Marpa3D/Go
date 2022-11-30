// Узнаю, сколько байт на конкретном компе занимает тип int
package main

import (
	"fmt"
	"unsafe"
)

func main(){

	aInt := 108
	fmt.Printf("Type %T size of %d bytes\n", aInt, unsafe.Sizeof(aInt))
}
