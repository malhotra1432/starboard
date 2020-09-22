package v1alpha1

import "time"

// Scanner is the spec for a scanner generating a security assessment report.
type Scanner struct {
	Name    string `json:"name"`
	Vendor  string `json:"vendor"`
	Version string `json:"version"`
}

type LastUpdated struct {
	Time time.Time `json:"timestamp"`
}
