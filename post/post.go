package post

import (
	"time"
)

// NewPost ...
func NewPost(title string, subtitle string, content string) *Post {
	return &Post{Title: title, SubTitle: subtitle, Content: content}
}

// NewTag ...
func NewTag(title string, description string) *Tag {
	return &Tag{Title: title, Description: description}
}

// NewCategory ...
func NewCategory(title string, description string) *Category {
	return &Category{Title: title, Description: description}
}

// Publish ...
func (p *Post) Publish() {
	p.Published = true
	p.PublishTime = time.Now()
}

// AddTag ...
func (p *Post) AddTag(t *Tag) {
	for _, tag := range p.Tags {
		if tag.Title == t.Title {
			return
		}
	}
	p.Tags = append(p.Tags, *t)
}

// AddCategory ...
func (p *Post) AddCategory(c *Category) {
	p.Category = c
}

// View ...
func (p *Post) View() {
	p.Viewed++
}

// Post ...
type Post struct {
	ID          int64
	UUID        string
	Title       string
	Category    *Category
	SubTitle    string
	Content     string
	Tags        []Tag
	CreateTime  time.Time
	UpdateTime  time.Time
	PublishTime time.Time
	// Author      *User
	Published bool
	Viewed    int64
	Liked     int64
}

// Category ...
type Category struct {
	ID          int
	Title       string
	Description string
}

// Tag ....
type Tag struct {
	ID          int
	Title       string
	Description string
}
