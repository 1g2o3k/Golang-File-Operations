package go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Define file path
	filePath := "example.txt"

	// Create a new file (or truncate if it exists)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Ensure the file is closed after all operations are done
	defer file.Close()

	// Write "hello" 10 times to the file and also print to terminal
	for i := 0; i < 10; i++ {
		_, err := file.WriteString("hello \n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Print("hello \n")
	}

	fmt.Println("-------------------------------------")

	// Read the contents of the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the file content
	fmt.Println(string(data))

	// Delete the file
	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
}
