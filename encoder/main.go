package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Something noteworthy happened!")

	// concurrency := 10
	// in := make(chan int)
	// done := make(chan []byte)

	// // auto exec func
	// go func() { // simulando um rabbitmq
	// 	i := 0
	// 	for { // for infinito
	// 		in <- i
	// 		i++
	// 	}
	// }()

	// for x := 0; x < concurrency; x++ {
	// 	go ProcessWorker(in, x)
	// }

	// <-done

}

// caiu um valor, será processado
// func ProcessWorker(in chan int, worker int) {
// 	for x := range in {
// 		time.Sleep(time.Duration(rand.Intn(4) * int(time.Second)))
// 		fmt.Println("worker ", worker, int(x))
// 	}
// }
