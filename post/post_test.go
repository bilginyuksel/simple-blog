package post

import (
	"testing"
	"time"

	util "github.com/bilginyuksel/simple-blog/utils"
)

func assertPostEquals(t *testing.T, given *Post, expected *Post) {
	util.AssertEquals(t, given.content, expected.content)
	util.AssertEquals(t, given.title, expected.title)
	util.AssertEquals(t, given.subTitle, expected.subTitle)
	util.AssertEquals(t, given.liked, expected.liked)
	util.AssertEquals(t, given.viewed, expected.viewed)
	util.AssertEquals(t, given.published, expected.published)

	util.AssertTimePrecision(t, given.createTime, expected.createTime)
	util.AssertTimePrecision(t, given.updateTime, expected.updateTime)
	util.AssertTimePrecision(t, given.publishTime, expected.publishTime)
}

func TestCreatePost_GetPost(t *testing.T) {
	pr := CreateRequest{Title: "my-title", SubTitle: "my-sub-title", Content: "big content", Author: nil}
	expected := &Post{title: "my-title", subTitle: "my-sub-title", content: "big content", createTime: time.Now()}
	given := pr.CreatePost()

	assertPostEquals(t, given, expected)
}
