package checker

import (
	"fmt"
	"net/http"
	"time"
)

type CheckResult struct {
	Target string
	Status string
	Err    error
}

func CheckUrl(url string, results chan<- CheckResult) {
	client := http.Client{
		Timeout: 6 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		results <- CheckResult{
			Target: url,
			Err:    fmt.Errorf("request failed: %w", err),
		}
		return
	}
	defer resp.Body.Close()

	results <- CheckResult{
		Target: url,
		Status: resp.Status,
	}
}
