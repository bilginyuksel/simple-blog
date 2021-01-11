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

func TestSaveSurvey_(t *testing.T) {

}
