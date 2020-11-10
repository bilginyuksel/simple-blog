package post

import "testing"

func TestNewPost_ExpectNewPost(t *testing.T) {
	post := NewPost("title", "subtitle", "content")
	if "title" != post.Title {
		t.Error()
	}

	if "subtitle" != post.SubTitle {
		t.Error()
	}

	if "content" != post.Content {
		t.Error()
	}
}

func TestNewPost_NotPublished(t *testing.T) {
	post := NewPost("title", "subtitle", "content")
	if post.Published {
		t.Error()
	}
}

func TestPublishPost_ExpectPublishTimeAndPublishedUpdated(t *testing.T) {
	post := NewPost("title", "subtitle", "content")
	oldTime := post.PublishTime
	if post.Published {
		t.Error()
	}
	post.Publish()
	if !post.Published {
		t.Error()
	}

	if oldTime == post.PublishTime {
		t.Error()
	}

	t.Logf("old time %v, new time %v", oldTime, post.PublishTime)
}

// func TestUpdatePost_Update(t *testing.T) {
// 	post := NewPost("tittltle", "subtitle", "content")
// }
func TestAddTag_ExpectNewTag(t *testing.T) {
	tag := NewTag("tag", "description")
	post := NewPost("title", "subtitle", "content")

	if len(post.Tags) != 0 {
		t.Error()
	}
	post.AddTag(tag)
	if len(post.Tags) == 0 {
		t.Error()
	}
}

func TestAddTagAddSameTag_ExpectNothing(t *testing.T) {
	tag := NewTag("tag", "description")
	post := NewPost("title", "subtitle", "content")
	post.AddTag(tag)
	post.AddTag(tag)

	if len(post.Tags) != 1 {
		t.Error()
	}
}

func TestViewPost_IncreaseView(t *testing.T) {
	post := NewPost("title", "content", "description")
	post.View()
	if 1 != post.Viewed {
		t.Error()
	}

	for i := 0; i < 10; i++ {
		post.View()
	}

	if 11 != post.Viewed {
		t.Error()
	}
}

func TestNewCategory_NewCategory(t *testing.T) {
	cat := NewCategory("category", "description")
	if "category" != cat.Title {
		t.Error()
	}

	if "description" != cat.Description {
		t.Error()
	}
}

func TestAddCategory_AddCatToPost(t *testing.T) {
	cat := NewCategory("category", "description")
	post := NewPost("title", "subtitle", "content")
	post.AddCategory(cat)
	if post.Category.Title != cat.Title {
		t.Error()
	}
	if post.Category.Description != cat.Description {
		t.Error()
	}
}
