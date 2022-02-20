package question

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/qanda/server/pkg/log"
	questionpb "github.com/qanda/server/pkg/question/v1alpha1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// validate ServerServer implements questionpb.ServerService
var _ questionpb.questionServiceServer = (*questionServer)(nil)

func (s *questionServer) Getquestion(ctx context.Context, msg *questionpb.GetquestionRequest) (*questionpb.GetquestionResponse, error) {
	log.Infof("get question")
	question, err := s.store.Get(msg.GetId())
	if err != nil {
		return &questionpb.GetquestionResponse{}, status.Errorf(codes.NotFound, "error get question, id: %s", msg.GetId())
	}
	return &questionpb.GetquestionResponse{Item: question}, nil
}

func (s *questionServer) Createquestion(ctx context.Context, msg *questionpb.CreatequestionRequest) (*questionpb.CreatequestionResponse, error) {
	log.Infof("create question")
	msg.Item.Metadata.Id = uuid.NewString()
	msg.Item.Metadata.CreationTimestamp = time.Now().String()
	question, err := s.store.Create(*msg.GetItem())
	if err != nil {
		return &questionpb.CreatequestionResponse{}, fmt.Errorf("error get question, id: %s", question)
	}
	return &questionpb.CreatequestionResponse{Item: question}, nil
}
