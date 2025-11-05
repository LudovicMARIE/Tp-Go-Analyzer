package checker

import (
	"net/http"
	"time"
)

type CheckResult struct {
	Target string
	Status string
	Err    error
}

func CheckUrl(url string) CheckResult {
	client := http.Client{
		Timeout: 6 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return CheckResult{
			Target: url,
			Err: &UnreachableError{
				URL: url,
				Err: err,
			},
		}
	}
	defer resp.Body.Close()

	return CheckResult{
		Target: url,
		Status: resp.Status,
	}
}
