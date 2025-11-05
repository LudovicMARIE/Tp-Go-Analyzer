package checker

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ludovicMARIE/Tp-Go-Analyzer/internal/config"
)

type CheckResult struct {
	InputTarget config.InputTarget
	Status      string
	Err         error
}

type ReportEntry struct {
	Name   string
	URL    string
	Owner  string
	Status string // "OK", "Ineccessible", ou "Erreur"
	ErrMsg string // Message d'erreur, omis si vide
}

func CheckUrl(target config.InputTarget) CheckResult {
	client := http.Client{
		Timeout: 6 * time.Second,
	}
	resp, err := client.Get(target.URL)
	if err != nil {
		return CheckResult{
			InputTarget: target,
			Err: &UnreachableError{
				URL: target.URL,
				Err: err,
			},
		}
	}
	defer resp.Body.Close()

	return CheckResult{
		InputTarget: target,
		Status:      resp.Status,
	}
}

func ConvertToReportEntry(res CheckResult) ReportEntry {
	report := ReportEntry{
		Name:   res.InputTarget.Name,
		URL:    res.InputTarget.URL,
		Owner:  res.InputTarget.Owner,
		Status: res.Status,
	}

	if res.Err != nil {
		var unreachable *UnreachableError
		if errors.As(res.Err, &unreachable) {
			report.Status = "Unreachable"
			report.ErrMsg = fmt.Sprintf("unreachable URL: %v", unreachable.URL)
		} else {
			report.Status = "Error"
			report.ErrMsg = fmt.Sprintf("erreur générique: %v", res.Err)
		}
	}
	return report
}
