package question

import (
	"fmt"

	questionsv1alpha1 "github.com/castlemilk/qanda/server/pkg/question/v1alpha1"
)

type firestoreStore struct {
	questions map[string]questionsv1alpha1.Question
}

func newFireStoreStorage() Store {
	return firestoreStore{questions: map[string]questionsv1alpha1.Question{}}
}

func (s firestoreStore) Get(id string) (*questionsv1alpha1.Question, error) {
	if val, ok := s.questions[id]; ok {
		return &val, nil
	}

	return &questionsv1alpha1.Question{}, fmt.Errorf("no question question found with id: %s", id)
}

func (s firestoreStore) Create(question questionsv1alpha1.Question) (*questionsv1alpha1.Question, error) {
	s.questions[question.Metadata.Id] = question
	return &question, nil
}

func (s firestoreStore) Delete(id string) (*questionsv1alpha1.Question, error) {
	question, err := s.Get(id)
	if err != nil {
		return &questionsv1alpha1.Question{}, err
	}
	delete(s.questions, id)
	return question, nil
}
