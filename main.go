package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// Create a new file
func createFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Leak file descriptors by opening the same file repeatedly without closing it
func leakFileDescriptors(fileName string) {
	for {
		for i := 0; i < 100; i++ {
			_, err := os.Open(fileName) // Removed 'file,' from here
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			// Intentionally do not close the file to leak a file descriptor
		}
		time.Sleep(1 * time.Second)
	}
}

// Leak memory by continuously allocating space without freeing it
func leakMemory() {
	leakyData := make([][]byte, 0)
	for {
		data := make([]byte, 100*1024*1024) // Allocate 100MB
		leakyData = append(leakyData, data)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// CLI flags
	fdsFlag := flag.Bool("fds", false, "Leak file descriptors")
	memoryFlag := flag.Bool("memory", false, "Leak memory")

	flag.Parse()

	if !*fdsFlag && !*memoryFlag {
		fmt.Println("Specify either --fds or --memory to demonstrate resource leak.")
		return
	}

	if *fdsFlag {
		fmt.Println("Starting to leak file descriptors...")

		fileName := "leak.txt"
		file, err := createFile(fileName)
		if err != nil {
			fmt.Println("Could not create file:", err)
			return
		}
		file.Close() // Close the file after creating it

		go leakFileDescriptors(fileName)
	}

	if *memoryFlag {
		fmt.Println("Starting to leak memory...")
		go leakMemory()
	}

	// Run indefinitely
	select {}
}
