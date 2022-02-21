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
	Create(questionsv1alpha1.Question) (*questionsv1alpha1.Question, error)
	Delete(string) (*questionsv1alpha1.Question, error)
}

func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage( /*...*/ )
	case FireStore:
		return newFireStoreStorage( /*...*/ )
	default:
		return newMemoryStorage( /*...*/ )
	}
}
