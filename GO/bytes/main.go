package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	seekerTest()
	//readerTest()
}

func seekerTest() {
	reader := strings.NewReader("the quick brown fox jumps over the lazy dog")
	reader1, _ := ioutil.ReadAll(reader)
	reader.Seek(2, io.SeekStart)
	reader2, _ := ioutil.ReadAll(reader)
	fmt.Println(string(reader1))
	fmt.Println(string(reader2))
}

func readerTest() {
	someString := "hello world\nand hello go and more"
	myReader := strings.NewReader(someString)

	buffer := make([]byte, 200)

	log.Println(myReader.Read(buffer)) // *strings.Reader

	for {
		count, err := myReader.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Printf("Count: %v\n", count)
		fmt.Printf("Data: %v\n", string(buffer))
	}
}
