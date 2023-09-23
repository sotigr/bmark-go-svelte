package store

import (
	"time"
)

type File struct {
	Name        string    `json:"name"`
	FullName    string    `json:"fullName"`
	Size        int64     `json:"size"`
	Created     time.Time `json:"created"`
	ContentType string    `json:"contentType"`
}

type Dir struct {
	Name     string `json:"name"`
	FullName string `json:"fullName"`
}

type DirectoryListing struct {
	Files   *[]File `json:"files"`
	Folders *[]Dir  `json:"folders"`
}

type FileLock struct {
	Name string
}
