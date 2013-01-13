
package main

import (
	"os"
	"fmt"
	"log"
	"bytes"
	"math/rand"
)

func main() {
	names := map[int]bool{}
	for len(names) < 5000 {
		names[rand.Int()] = true
	}

	count := 0
	for name, _ := range names {
		path := fmt.Sprint("f", name, ".txt")
		f, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		} else if _, err := f.Write(getData(count)); err != nil {
			log.Fatal(err)
		}
		f.Close()
		count++
	}
}

func getData(i int) []byte {
	data := lines()
	if i == 3785 {
		data[4444] = []byte("he11o")
	} else if i == 2433 {
		data[4994] = []byte("bannana")
	}

	return bytes.Join(data, []byte("\n"))
}

func lines() [][]byte {
	data := [][]byte{}
	for i := 0; i < 4997; i++ {
		data = append(data, []byte("hello"))
	}

	data = append(data, []byte("banana"))
	data = append(data, []byte("apple"))
	data = append(data, []byte("orange"))
	return data
}
