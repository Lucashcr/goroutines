package entities

import (
	"log"
	"time"
)

type Consumer struct {
	id            int
	inputChannel  chan Task
	outputChannel chan Task
}

func NewConsumer(id int, inputChannel chan Task, outputChannel chan Task) *Consumer {
	return &Consumer{id: id, inputChannel: inputChannel, outputChannel: outputChannel}
}

func (c *Consumer) Run() {
	for t := range c.inputChannel {
		time.Sleep(time.Duration(t.Value) * time.Second)
		log.Printf("Task %s finished by consumer %d in %d seconds!\n", t.Id, c.id, t.Value)
	}
}
