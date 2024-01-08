package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.golang.org/",
		"https://googledddd.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
				return nil
			}
			fmt.Println("=====> ", err)
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err != nil {
		fmt.Println("errrrror", err)
		return
	}
	fmt.Println("Successfully fetched all URLs.")
}
