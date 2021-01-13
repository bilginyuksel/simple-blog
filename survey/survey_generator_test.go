package survey

import "testing"

func failIfError(err error, t *testing.T) {
	if err != nil {
		t.Fail()
	}
}

func failIfNotError(err error, t *testing.T) {
	if err == nil {
		t.Fail()
	}
}

func TestGenerate_NoTitleNoQuestionsNoResults(t *testing.T) {

	questions := []Question{}
	results := []Result{}

	reqNoTitle := SaveRequest{Questions: questions, Results: results}
	reqNoQuestions := SaveRequest{Results: results, Title: "Title"}
	reqNoResults := SaveRequest{Questions: questions, Title: "Title"}

	_, err := Generate(&reqNoTitle)
	failIfNotError(err, t)
	_, err = Generate(&reqNoQuestions)
	failIfNotError(err, t)
	_, err = Generate(&reqNoResults)
	failIfNotError(err, t)
}

func TestGenerate_ValidRequest(t *testing.T) {

	questions := []Question{}
	results := []Result{}
	request := SaveRequest{Questions: questions, Results: results, Title: "Title", Description: "Description"}

	survey, err := Generate(&request)
	failIfError(err, t)
	if survey.Title != "Title" || survey.Description != "Description" {
		t.Fail()
	}
}
