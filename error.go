package easycron

// Error response.
type Error struct {
	Status string `json:"status"`
	Err    struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// Error string.
func (e *Error) Error() string {
	return "Easycron Error Code " + e.Err.Code + ": " + e.Err.Message
}
