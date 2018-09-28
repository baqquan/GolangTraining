package datasources

import (
	"math/rand"
	"time"
)

// DataStreamFloat is a struct that we will pass through to whoever consumes the datastream
type DataStreamFloat struct {
	Description string
	Data        float64
}

// DataSourceFloat outputs a data stream that has description and data(float64)
func DataSourceFloat(c chan DataStreamFloat) {
	rand.Seed(time.Now().UnixNano())
	dataOut := DataStreamFloat{
		Description: "these should be float value",
	}
	for i := 0; ; i++ {
		dataOut.Data = rand.Float64()
		c <- dataOut
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}
