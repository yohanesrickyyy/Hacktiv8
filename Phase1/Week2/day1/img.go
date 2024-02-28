package main

import (
	"fmt"
	"time"
)

func processImage(imageUrl string) {
	time.Sleep(3 * time.Second)
	fmt.Printf("Processing image: https: %s\n", imageUrl)
	fmt.Printf("Image processing completed: https: %s\n", imageUrl)
}

func main() {
	images := []string{
		"//example.com/images1.jpg",
		"//example.com/images2.jpg",
		"//example.com/images3.jpg",
		"//example.com/images4.jpg",
	}
	for _, image := range images {
		processImage(image)
	}
}
