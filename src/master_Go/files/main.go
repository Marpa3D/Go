// Работа с файлами
package main

import (
	"fmt"
	"os"
)

func main() {
	pf := fmt.Printf

	var newFile *os.File
	pf("%T\n", newFile)
}
