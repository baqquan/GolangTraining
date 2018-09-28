package main

import (
	"fmt"

	dataAgg "github.com/baqquan/GolangTraining/CodingChallege/WithPackages/data_aggregator"

	d "github.com/baqquan/GolangTraining/CodingChallege/WithPackages/datasources"
)

func main() {
	c1 := make(chan d.DataStreamBool)
	c2 := make(chan d.DataStreamFloat)
	stopChannel := make(chan string)
	go d.DataSourceBool(c1)
	go d.DataSourceFloat(c2)

	go func() {
		for i := true; i; {
			fmt.Println(`Enter "stop" to exit: `)
			var quit string
			fmt.Scanln(&quit)
			if quit == "stop" {
				i = false
				stopChannel <- "stop message"
				close(c1)
				close(c2)
			}
		}
	}()

	dataAgg.DataAggregator(c1, c2, stopChannel, "dummyFile.txt")
}
