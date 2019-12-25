package models

import (
	"math/rand"
	"time"
)


type model struct{
}

type Model interface {
	// Return read only channel that has information about % current cpu
	GetLiveCpuUsage()( <- chan Cpu , error)
}

func NewModel() Model {
	return &model{}
}


type Cpu struct{
	PercentageUsage int 	`json:"cpuPercentageUsage"`
}

// Consider this model responsible for puling data from backend and serving it to controller
func (m *model) GetLiveCpuUsage() (<- chan Cpu , error){
	ch := make(chan Cpu)
	go WriteValues(ch)

	return ch, nil
}

// Generates random values and writes to passed on channel
func WriteValues(ch chan <- Cpu ) error {
	ticker := time.NewTicker(1 * time.Second)
	for{
		<- ticker.C
		newVal := Cpu{PercentageUsage:rand.Intn(100)}
		ch <- newVal
	}
}