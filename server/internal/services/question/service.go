package question

import (
	"context"
	"fmt"
	"time"

	"github.com/castlemilk/qanda/server/pkg/log"
	questionpb "github.com/castlemilk/qanda/server/pkg/question/v1alpha1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// validate ServerServer implements questionpb.ServerService
var _ questionpb.QuestionServiceServer = (*questionServer)(nil)

func (s *questionServer) GetQuestion(ctx context.Context, msg *questionpb.GetQuestionRequest) (*questionpb.GetQuestionResponse, error) {
	log.Infof("get question")
	question, err := s.store.Get(msg.GetId())
	if err != nil {
		return &questionpb.GetQuestionResponse{}, status.Errorf(codes.NotFound, "error get question, id: %s", msg.GetId())
	}
	return &questionpb.GetQuestionResponse{Question: question}, nil
}

func (s *questionServer) CreateQuestion(ctx context.Context, msg *questionpb.CreateQuestionRequest) (*questionpb.CreateQuestionResponse, error) {
	log.Infof("create question")
	msg.Question.Metadata.Id = uuid.NewString()
	msg.Question.Metadata.CreatedAt = time.Now().Unix()
	question, err := s.store.Create(*msg.GetQuestion())
	if err != nil {
		return &questionpb.CreateQuestionResponse{}, fmt.Errorf("error get question, id: %s", question)
	}
	return &questionpb.CreateQuestionResponse{Question: question}, nil
}
