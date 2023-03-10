package main

//
// start the coordinator process, which is implemented
// in ../mr/coordinator.go
//
// go run mrcoordinator.go pg*.txt
//
// Please do not change this file.
//

import (
	"fmt"
	"os"
	"time"

	"MapReduce/mr"
)

func main() {
	//determine if input's length is valid
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: mrcoordinator inputfiles...\n")
		os.Exit(1)
	}

	//initialize coordinator with (inputFiles, 10 reduce tasks)
	m := mr.MakeCoordinator(os.Args[1:], 10)

	//keep checking if the coordinator is done per second
	for m.Done() == false {
		//stop the latest go-routine for at least the stated duration time.Second
		time.Sleep(time.Second)
	}

	//why we want to sleep before exit?
	time.Sleep(time.Second)
}
