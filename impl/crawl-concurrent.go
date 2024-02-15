package impl

import (
	"fmt"
	"github.com/antigloss/go/container/concurrent/queue"
	"golang.org/x/net/html"
	"sync"
)

func StartCrawlingConcurrent(seedUrl string) {
	fmt.Println("using concurrent bfs implementation")

	level := 0
	q := queue.NewLockfreeQueue[string]()
	q.Push(seedUrl)

	for q != nil {

		var wg sync.WaitGroup
		htmlContentCh := make(chan *html.Node)
		linksCh := make(chan []string)
		curr, err := q.Pop()
		level++
		if !err {
			panic(err)
		}
		// Start a goroutine for fetching the HTML content
		wg.Add(1)
		go fetch(curr, &wg, htmlContentCh)

		// Start a goroutine for parsing the HTML content
		wg.Add(1)
		go parse(htmlContentCh, &wg, linksCh)

		go func() {
			wg.Wait()
			close(htmlContentCh)
			close(linksCh)
		}()

		// Wait for the parsing to complete and collect the links
		links := <-linksCh

		for _, link := range links {
			q.Push(link)
		}

		fmt.Println(links)

		if level == 10 {
			break
		}
	}
}

func fetch(url string, wg *sync.WaitGroup, htmlContentCh chan<- *html.Node) {
	defer wg.Done()
	htmlContent, err := Fetch(url)
	if err != nil {
		panic(err)
	}
	htmlContentCh <- htmlContent
}

func parse(htmlContentCh <-chan *html.Node, wg *sync.WaitGroup, linksCh chan<- []string) {
	defer wg.Done()
	htmlContent := <-htmlContentCh
	links, err := Parse(htmlContent)
	if err != nil {
		panic(err)
	}
	linksCh <- links
}
