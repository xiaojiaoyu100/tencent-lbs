package txlbs

// Meta includes common response info.
type Meta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Success means a valid response.
func (r *Meta) Success() bool {
	return r != nil && r.Status == 0
}
