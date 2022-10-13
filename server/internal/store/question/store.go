package question

import (
	questionsv1alpha1 "github.com/castlemilk/qanda/server/pkg/question/v1alpha1"
)

type StorageType string

const (
	FireStore     StorageType = "firestore"
	MemoryStorage             = "inmemory"
)

type Store interface {
	Get(string) (*questionsv1alpha1.Question, error)
	List() ([]*questionsv1alpha1.Question, error)
	Create(questionsv1alpha1.Question) (*questionsv1alpha1.Question, error)
	Delete(string) (string, error)
}

func NewStore(t StorageType) (Store, error) {
	switch t {
	case MemoryStorage:
		return newMemoryStorage( /*...*/ )
	case FireStore:
		return newFireStoreStorage( /*...*/ )
	default:
		return newMemoryStorage( /*...*/ )
	}
}
