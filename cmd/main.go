package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Lucashcr/goroutine/internal/entities"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must provide the number of tasks to be executed!\n")
		os.Exit(1)
	}

	workersCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("The number of workers must be a number!\n")
		os.Exit(1)
	}

	log.Println("Starting the producer and consumers...")
	inputChannel := make(chan entities.Task)
	outputChannel := make(chan entities.Task)

	for i := 1; i <= workersCount; i++ {
		consumer := entities.NewConsumer(i, inputChannel, outputChannel)
		go consumer.Run()
	}

	producer := entities.NewProducer(inputChannel)
	log.Println("Starting to execute tasks...")
	producer.Run()
}
