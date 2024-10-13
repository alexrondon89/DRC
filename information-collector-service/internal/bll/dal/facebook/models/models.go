package models

// Groups section
type Groups struct {
	Data   []Group `json:"data"`
	Paging Paging  `json:"paging"`
}

type Group struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Privacy     string `json:"privacy"`
	UpdatedTime string `json:"updated_time"`
}

// Posts section
type Posts struct {
	Data   []Post `json:"data"`
	Paging Paging `json:"paging"`
}

type Post struct {
	ID        string `json:"id"`
	Message   string `json:"message,omitempty"`
	CreatedAt string `json:"created_time"`
	From      User   `json:"from"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Comments section
type Comments struct {
	Data   []Comment `json:"data"`
	Paging Paging    `json:"paging"`
}

type Comment struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_time"`
	From      User   `json:"from"`
}

// Paging section
type Paging struct {
	Previous string `json:"previous"`
	Next     string `json:"next"`
}
