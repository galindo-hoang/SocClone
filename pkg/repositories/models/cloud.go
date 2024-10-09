package models

type CloudEntity struct {
	ContentType string `json:"content_type"`
	ObjectName  string `json:"object_name"`
	FilePath    string `json:"prefix_path"`
	Size        int64  `json:"size"`
	Retry       int    `json:"retry"`
}
