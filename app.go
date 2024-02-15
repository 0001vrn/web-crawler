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
	fmt.Println(argsWithoutProg)
	if len(argsWithoutProg) != 0 {
		seedUrl = argsWithoutProg[0]
	}

	if len(argsWithoutProg) != 0 && len(argsWithoutProg[1]) != 0 {
		ExecutionMode, err := strconv.Atoi(argsWithoutProg[1])
		if err != nil {
			panic(err)
		}

		if ExecutionMode == SimpleCrawlingExecutionMode {
			impl.StartCrawlingSimple(seedUrl)
		} else if ExecutionMode == ConcurrentCrawlingExecutionMode {
			impl.StartCrawlingConcurrent(seedUrl)
		} else {
			impl.StartCrawlingConcurrent(seedUrl)
		}

	} else {
		impl.StartCrawlingSimple(seedUrl)
	}
}
