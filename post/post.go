package post

import (
	"time"

    "github.com/bilginyuksel/simple-blog/user"
)

type PostRequest struct {
    Title string
    SubTitle string
    Content string

    //Category *Category
    //tags []Tag
    Author *user.User
}

func (pr *PostRequest) CreatePost() *Post{

    return &Post{
        id: 1,
        uuid: "randomuuid",
        title: pr.Title,
        subTitle: pr.SubTitle,
        content: pr.Content,
        tags: nil,
        createTime: time.Now(),
        updateTime: time.Now(),
        publishTime: time.Now(),
        author: pr.Author,
        published: false,
        viewed: 0,
        liked: 0,
    }
}

// Post ...
type Post struct {
	id          int64
	uuid        string
	title       string
	category    *Category
	subTitle    string
	content     string
	tags        []Tag
	createTime  time.Time
	updateTime  time.Time
	publishTime time.Time
	author      *user.User
	published bool
	viewed    int64
	liked     int64
}

// Category ...
type Category struct {
	id          int
	title       string
	description string
}

// Tag ....
type Tag struct {
	id          int
	title       string
	description string
}

