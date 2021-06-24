package helpers

type UpdateResponse struct {
	Status string `json:"status"`
}

type ReadResponse struct {
	Status string `json:"status"`
	NPM    string `json:"npm"`
	Name   string `json:"name"`
}
