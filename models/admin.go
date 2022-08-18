package models

// AddResourceRequest add resource request.
type AddResourceRequest struct {
	DagId int    `json:"dag_id"`
	Name  string `json:"name"`
}
