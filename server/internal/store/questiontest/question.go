package questiontest

import (
	"github.com/google/uuid"
	apiv1alpha1 "github.com/qanda/server/pkg/api/v1alpha1"
	questionsv1alpha1 "github.com/qanda/server/pkg/question/v1alpha1"
)

func Newquestion() questionsv1alpha1.questionDetails {
	return questionsv1alpha1.questionDetails{Metadata: &apiv1alpha1.Metadata{
		Id:   uuid.NewString(),
		Name: "test",
	}, question: &apiv1alpha1.question{Category: "test", Description: "test"}}
}
