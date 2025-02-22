package entities

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Producer struct {
	inputChannel chan Task
}

func NewProducer(inputChannel chan Task) *Producer {
	return &Producer{inputChannel: inputChannel}
}

func (p *Producer) Run() {
	for {
		time.Sleep(1 * time.Second)
		id := uuid.New()
		p.inputChannel <- Task{Id: id, Value: rand.Intn(5)}
		log.Printf("Task %s added to the queue!\n", id)
	}
}
