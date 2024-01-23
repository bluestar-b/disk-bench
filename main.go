package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	filePath := "benchmark_file.txt"
	fileSizeMB := 100

	// Measure write speed
	writeSpeedMBps, writeTime := measureWriteSpeed(filePath, fileSizeMB)
	fmt.Printf("Write speed: %.2f MB/s\n", writeSpeedMBps)
	fmt.Printf("Write time taken: %v\n", writeTime)

	// Measure read speed
	readSpeedMBps, readTime := measureReadSpeed(filePath, fileSizeMB)
	fmt.Printf("Read speed: %.2f MB/s\n", readSpeedMBps)
	fmt.Printf("Read time taken: %v\n", readTime)

	if err := os.Remove(filePath); err != nil {
		fmt.Println("Error removing file:", err)
	}
}

func measureWriteSpeed(filePath string, fileSizeMB int) (float64, time.Duration) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return 0, 0
	}
	defer file.Close()

	// Prepare data to be written
	data := make([]byte, 1024*1024) // 1 MB
	for i := 0; i < len(data); i++ {
		data[i] = byte(i % 256)
	}

	startTime := time.Now()
	for i := 0; i < fileSizeMB; i++ {
		_, err := file.Write(data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return 0, 0
		}
	}
	elapsedTime := time.Since(startTime)

	writeSpeedMBps := float64(fileSizeMB) / elapsedTime.Seconds()
	return writeSpeedMBps, elapsedTime
}

func measureReadSpeed(filePath string, fileSizeMB int) (float64, time.Duration) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, 0
	}
	defer file.Close()

	buffer := make([]byte, 1024*1024) // 1 MB

	startTime := time.Now()
	for i := 0; i < fileSizeMB; i++ {
		_, err := io.ReadFull(file, buffer)
		if err != nil {
			fmt.Println("Error reading from file:", err)
			return 0, 0
		}
	}
	elapsedTime := time.Since(startTime)

	readSpeedMBps := float64(fileSizeMB) / elapsedTime.Seconds()
	return readSpeedMBps, elapsedTime
}
