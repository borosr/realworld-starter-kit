package models

import "time"

type image string

type User struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    *image `json:"image"`
}

type Profile struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     *image `json:"image"`
	Following bool   `json:"following"`
}

type Article struct {
	Slug           string    `json:"slug"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Body           string    `json:"body"`
	TagList        []string  `json:"tagList"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Favorited      bool      `json:"favorited"`
	FavoritesCount uint      `json:"favoritesCount"`
	Author         *Profile  `json:"author"`
}

type Articles struct {
	Articles      []Article `json:"articles"`
	ArticlesCount uint      `json:"articlesCount"`
}

type Comment struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Body      string    `json:"body"`
	Author    *Profile  `json:"author"`
}

type Comments struct {
	Comments []Comment `json:"comments"`
}

type Tags struct {
	Tags []string `json:"tags"`
}