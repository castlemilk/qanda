package question

import (
	"fmt"

	"github.com/castlemilk/qanda/server/pkg/log"
	questionsv1alpha1 "github.com/castlemilk/qanda/server/pkg/question/v1alpha1"
)

type memoryStore struct {
	questions map[string]questionsv1alpha1.Question
}

func newMemoryStorage() (Store, error) {
	return memoryStore{questions: map[string]questionsv1alpha1.Question{}}, nil
}

func (s memoryStore) Get(id string) (*questionsv1alpha1.Question, error) {
	if val, ok := s.questions[id]; ok {
		return &val, nil
	}

	return &questionsv1alpha1.Question{}, fmt.Errorf("no question question found with id: %s", id)
}

func (s memoryStore) List() ([]*questionsv1alpha1.Question, error) {

	values := make([]*questionsv1alpha1.Question, 0, len(s.questions))
	log.Infof("values: %+v", s.questions)
	for _, v := range s.questions {
		values = append(values, &v)
	}

	return values, nil
}

func (s memoryStore) Create(question questionsv1alpha1.Question) (*questionsv1alpha1.Question, error) {
	s.questions[question.Metadata.Id] = question
	return &question, nil
}

func (s memoryStore) Delete(id string) (string, error) {
	_, err := s.Get(id)
	if err != nil {
		return "", err
	}
	delete(s.questions, id)
	return id, nil
}
