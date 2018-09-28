package datasources

import (
	"math/rand"
	"time"
)

// DataStreamBool is a struct that we will pass through to whoever consumes the datastream
type DataStreamBool struct {
	Description string
	Data        bool
}

// function to randomly set a boolean. Not very perfomant but works.
func randbool() bool {
	return rand.Int()%2 == 0
}

// DataSourceBool outputs a data stream that has description and data (boolean)
func DataSourceBool(c chan DataStreamBool) {
	rand.Seed(time.Now().UnixNano())
	dataOut := DataStreamBool{
		Description: "random boolean datasource",
		Data:        randbool(),
	}
	for i := 0; ; i++ {
		dataOut.Data = randbool()
		c <- dataOut
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}
