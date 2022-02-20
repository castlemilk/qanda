package question

import (
	"fmt"

	questionsv1alpha1 "github.com/qanda/server/pkg/question/v1alpha1"
)

type firestoreStore struct {
	questions map[string]questionsv1alpha1.questionDetails
}

func newFireStoreStorage() Store {
	return firestoreStore{questions: map[string]questionsv1alpha1.questionDetails{}}
}

func (s firestoreStore) Get(id string) (*questionsv1alpha1.questionDetails, error) {
	if val, ok := s.questions[id]; ok {
		return &val, nil
	}

	return &questionsv1alpha1.questionDetails{}, fmt.Errorf("no question question found with id: %s", id)
}

func (s firestoreStore) Create(question questionsv1alpha1.questionDetails) (*questionsv1alpha1.questionDetails, error) {
	s.questions[question.Metadata.Id] = question
	return &question, nil
}

func (s firestoreStore) Delete(id string) (*questionsv1alpha1.questionDetails, error) {
	question, err := s.Get(id)
	if err != nil {
		return &questionsv1alpha1.questionDetails{}, err
	}
	delete(s.questions, id)
	return question, nil
}
