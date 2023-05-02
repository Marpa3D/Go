// Запись в файл с помощью пакета bufio
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("textFile.txt", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// пишем, используя буфер
	buferWriter := bufio.NewWriter(file)
	bs := []byte{98, 99, 100}

	bytesWritten, err := buferWriter.Write(bs)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Байты были записаны в буфер (не в файл)! %d\n", bytesWritten)

	// проверяем, насколько заполнен буфер, в байтах, его размер по умолчанию: 4096 байт!
	bufBytesFree := buferWriter.Available()
	log.Printf("Свободных байт в буфере: %d", bufBytesFree)

	// пишем в файл

	bytesWritten, err = buferWriter.WriteString("\nЭто тестовая строка.")

	if err != nil {
		log.Fatal(err)
	}
	unflushedBufferSize := buferWriter.Buffered()

	log.Printf("Байтов буферезировано: %d", unflushedBufferSize)

	// вот этот момент записи в файл)
	buferWriter.Flush()

	// очистка буфера (если надо)
	buferWriter.Reset(buferWriter)

	// проверям буфер
	bufBytesFree = buferWriter.Available()
	log.Printf("Свободных байт в буфере: %d", bufBytesFree)
}
