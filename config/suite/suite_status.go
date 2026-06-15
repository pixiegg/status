package suite

// Status represents the status of a suite
type Status struct {
	// Name of the suite
	Name string `json:"name,omitempty"`

	// Group the suite is a part of. Used for grouping multiple suites together on the front end.
	Group string `json:"group,omitempty"`

	// Key of the Suite
	Key string `json:"key"`

	// Results is the list of suite execution results
	Results []*Result `json:"results"`

	// Uptime30d is the uptime percentage over the last 30 days (0.0 to 1.0)
	Uptime30d *float64 `json:"uptime30d,omitempty"`
}

// NewStatus creates a new Status for a given Suite
func NewStatus(s *Suite) *Status {
	return &Status{
		Name:    s.Name,
		Group:   s.Group,
		Key:     s.Key(),
		Results: []*Result{},
	}
}