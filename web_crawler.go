/*Exercise: Web Crawler
In this exercise you'll use Go's concurrency features to parallelize a web crawler.
Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.
*/
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan string) {
	defer close(ch)
	if depth <= 0 {
		return
	}

	if cache.Cached(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ch <- err.Error()
		return
	}
	ch <- fmt.Sprintf("found: %s %q\n", url, body)
	
	result := make([]chan string, len(urls))
	for i, u := range urls {
		result[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, result[i])
	}
	
	for i := range result {
        for s := range result[i] {
            ch <- s
        }
    }	
	return
}

// SafeCache is safe to use concurrently.
type SafeCache struct {
	mu        sync.Mutex
	cacheURLs map[string]struct{}
}

func (c *SafeCache) Cached(url string) bool {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.cacheURLs.
	defer c.mu.Unlock()
	_, ok := c.cacheURLs[url]
	if ok {
		return true
	} else {
		c.cacheURLs[url] = struct{}{}
		return false
	}
}

var cache = SafeCache{cacheURLs: make(map[string]struct{})}

func main() {
	c := make(chan string)
	go Crawl("https://golang.org/", 4, fetcher, c)
	
	for i := range c {		
		fmt.Println(i)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
