package questiontest

import (
	apiv1alpha1 "github.com/castlemilk/qanda/server/pkg/api/v1alpha1"
	questionsv1alpha1 "github.com/castlemilk/qanda/server/pkg/question/v1alpha1"
	"github.com/google/uuid"
)

func Newquestion() questionsv1alpha1.Question {
	return questionsv1alpha1.Question{Metadata: &apiv1alpha1.Metadata{
		Id:   uuid.NewString(),
		Name: "test",
	}, Spec: &questionsv1alpha1.QuestionSpec{Content: &apiv1alpha1.Content{Content: &apiv1alpha1.Content_Text{Text: &apiv1alpha1.Text{Value: "test"}}}, Tags: []*apiv1alpha1.Tag{}}}
}
