package impl

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/set"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

func StartCrawlingSimple(seedUrl string) {
	fmt.Println("using simple bfs implementation")
	level := 0
	visitedLinks := set.New()
	q := queue.New()
	q.Enqueue(seedUrl)
	for q.Len() > 0 {
		var curr = q.Dequeue().(string)
		if visitedLinks.Has(curr) {
			continue
		} else {
			visitedLinks.Insert(curr)
		}
		level++
		htmlContent, err := Fetch(curr)
		if err != nil {
			panic(err)
		}
		links, err := Parse(htmlContent)
		if err != nil {
			panic(err)
		}
		fmt.Println(len(links))

		for _, link := range links {
			q.Enqueue(link)
		}

		if level == 10 {
			break
		}
	}

	fmt.Println("unique sites count -> ", visitedLinks.Len())
}

func Fetch(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	return doc, err
}

func Parse(htmlNode *html.Node) ([]string, error) {
	// Parse HTML content
	links := getAllLinks(htmlNode)
	//fmt.Println(links)
	return links, nil
}

func getAllLinks(n *html.Node) []string {
	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" && checkPrefixProtocol(a.Val) {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		exLinks := getAllLinks(c)
		links = append(links, exLinks...)
	}
	return links
}

func checkPrefixProtocol(val string) bool {
	return strings.HasPrefix(val, "http") || strings.HasPrefix(val, "https")
}
