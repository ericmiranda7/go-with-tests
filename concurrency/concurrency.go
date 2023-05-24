package concurrency

type WebsiteChecker func(string) bool

type Result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resChan := make(chan Result)

	for _, url := range urls {
		go func(u string) {
			resChan <- Result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resChan
		results[r.string] = r.bool
	}

	return results
}
