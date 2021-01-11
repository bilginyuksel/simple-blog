package survey

import (
	"errors"
	"fmt"
	"time"

	"github.com/bilginyuksel/simple-blog/user"
)

func init() {
	fmt.Println("Hello world")
}

// SaveRequest ...
type SaveRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
	Results     []Result   `json:"results"`
	Author      user.User  `json:"author"`
}

// Generate ...
func Generate(request *SaveRequest) (*Survey, error) {

	if request.Title == "" || request.Questions == nil || request.Results == nil {
		return nil, errors.New("title, question and results fields are mandatory")
	}

	return &Survey{
		id:          111,
		Title:       request.Title,
		Description: request.Description,
		Questions:   request.Questions,
		Results:     request.Results,
		Author:      request.Author,
		CreatedTime: time.Now(),
		ViewCount:   0,
		SolvedCount: 0,
		LikedCount:  0,
	}, nil
}

// Survey ...
type Survey struct {
	id          int64
	Title       string
	Description string
	Questions   []Question
	Results     []Result
	Author      user.User
	CreatedTime time.Time
	ViewCount   int64
	SolvedCount int64
	LikedCount  int64
}

// Question ...
type Question struct {
	ID      int64
	Content string
	Choices []Choice
}

// Result ...
type Result struct {
	Title       string
	Description string
}

// Choice ...
type Choice struct {
	ID            int64
	LinkedResults []Result
	Content       string
}
