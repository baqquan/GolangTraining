package dataaggregator

import (
	"log"
	"os"
	"strconv"
	"time"

	ds "github.com/baqquan/GolangTraining/CodingChallege/WithPackages/datasources"
)

//DataAggregator takes in 2 channels, c1 and c2, from datasource package. Will stream data from the channels to a text file based on filepath.
func DataAggregator(c1 chan ds.DataStreamBool, c2 chan ds.DataStreamFloat, stop chan string, filepath string) {
	stringToWrite := ""
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for {
		select {
		case msg1 := <-c1:
			// streaming data to file with a proper time stamp. This specifically with a dataStreamBool data type.
			stringToWrite = "[" + time.Now().Format(time.RFC3339) + "]    " + msg1.Description + "    " + strconv.FormatBool(msg1.Data) + "\n"
			if _, err := f.Write([]byte(stringToWrite)); err != nil {
				log.Fatal(err)
			}
		case msg2 := <-c2:
			stringToWrite = "[" + time.Now().Format(time.RFC3339) + "]    " + msg2.Description + "    " + strconv.FormatFloat(msg2.Data, 'E', -1, 64) + "\n"
			if _, err := f.Write([]byte(stringToWrite)); err != nil {
				log.Fatal(err)
			}
			//third case will stop this loop and exit out of the program once you get a message from the stop channel
		case msg3 := <-stop:
			if msg3 == "stop message" {
				return
			}
		}
	}
}
