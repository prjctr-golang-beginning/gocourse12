package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//1 write
	data := []byte("Hello Bold!")
	text := "Hello Gold!"
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, _ = file.WriteString(text)
	_, _ = file.Write(data)
	fmt.Println("Done.")

	//2.1 read
	content, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(string(content))

	//2.2 read
	file, err = os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data = make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Print(string(data[:n]))
	}

	//3 remove
	err = os.Remove("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = os.RemoveAll("else")
	if err != nil {
		log.Fatal(err)
	}

	// exercises
	// 1. Створити файл з текстом в три рядки, де серед текста є наступні посилання: https://www.google.com/, https://bing.com, chat.openai.com
	// 2. Считати файл порядково
	// 3. Написати регулярний вираз, за допомогою якого знайдуться всі три посилання ( використовуйте https://regex101.com/ )
	// 4. Видалити файл
}
