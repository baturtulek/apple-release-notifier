package types

// Release Single Release
type Release struct {
	Platform string `json:"platform"`
	Version  string `json:"version"`
	Code     string `json:"code"`
}
