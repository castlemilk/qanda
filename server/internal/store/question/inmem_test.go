package question

import (
	"testing"

	"github.com/qanda/server/internal/store/questiontest"
	questionsv1alpha1 "github.com/qanda/server/pkg/question/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func TestNewStore(t *testing.T) {
	store := newMemoryStorage()
	assert.Equal(t, store, memoryStore{questions: map[string]questionsv1alpha1.questionDetails{}}, "should be equal")
}

func TestCreatequestion(t *testing.T) {
	store := newMemoryStorage()
	want := questiontest.Newquestion()
	store.Create(want)
	got, _ := store.Get(want.Metadata.Id)
	assert.Equal(t, got, &want, "producs creation works")
}

func TestDeletequestion(t *testing.T) {
	store := newMemoryStorage()
	want := questiontest.Newquestion()
	store.Create(want)
	got, _ := store.Delete(want.Metadata.Id)
	assert.Equal(t, got, &want, "producs creation works")
}
