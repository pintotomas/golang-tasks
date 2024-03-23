package task3

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// Crawl function crawls the given URL and extracts links recursively
func Crawl(ctx context.Context, url string, depth int, client *http.Client) {
	if depth <= 0 {
		return
	}

	// Fetch the web page with a timeout
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request for", url, ":", err)
		return
	}

	// Attach the context to the request to handle cancellation
	req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching", url, ":", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Encountered an error while closing body: %s\n", err)
		}
	}(resp.Body)

	// Parse HTML
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		select {
		case <-ctx.Done(): // Check if the context is canceled
			return
		default:
			tokenType := tokenizer.Next()
			if tokenType == html.ErrorToken {
				// There are no more tokens to process
				return
			}
			token := tokenizer.Token()

			// Links are usually on <a> tags in the href attribute
			if tokenType == html.StartTagToken && token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						link := attr.Val
						if !strings.HasPrefix(link, "http") {
							link = url + link // Convert relative URL to absolute URL
						}
						fmt.Println("Found link:", link)
						Crawl(ctx, link, depth-1, client) // Recursively crawl the link
						break
					}
				}
			}
		}
	}
}

func Run(url string, depth, t int) {
	startURL := url
	maxDepth := depth
	timeout := time.Duration(t) * time.Second // Set the timeout duration
	clientTimeout := 5 * time.Second

	// Create a context with a deadline
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Create an HTTP client with a timeout
	client := &http.Client{
		Timeout: clientTimeout, // Set the client timeout
	}

	// Call the Crawl function with the context
	Crawl(ctx, startURL, maxDepth, client)
}
