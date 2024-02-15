package main

import (
	"fmt"
	"os"
	"strconv"
	"web-crawler/impl"
)

const SimpleCrawlingExecutionMode = 0
const ConcurrentCrawlingExecutionMode = 1

func main() {
	argsWithoutProg := os.Args[1:]
	var seedUrl = "https://www.wikipedia.org/"

	if len(argsWithoutProg) != 0 {
		seedUrl = argsWithoutProg[0]
	}

	if len(argsWithoutProg) != 2 {
		fmt.Println("please run the command as : go run app.go https://www.wikipedia.org 1")
		os.Exit(1)
	}

	ExecutionMode, err := strconv.Atoi(argsWithoutProg[1])
	if err != nil {
		os.Exit(1)
	}

	if ExecutionMode == SimpleCrawlingExecutionMode {
		impl.StartCrawlingSimple(seedUrl)
	} else if ExecutionMode == ConcurrentCrawlingExecutionMode {
		impl.StartCrawlingConcurrent(seedUrl)
	}
}
