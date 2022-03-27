package question

import (
	"context"

	"github.com/castlemilk/qanda/server/internal/store/question"
	"github.com/castlemilk/qanda/server/pkg/log"
	questionpb "github.com/castlemilk/qanda/server/pkg/question/v1alpha1"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

// questionServer ...
type questionServer struct {
	config Config
	store  question.Store
	questionpb.UnimplementedQuestionServiceServer
	*grpcweb.WrappedGrpcServer
}

// New creates instance of the Server
func New(ctx context.Context, cfg Config) (*questionServer, error) {
	var store question.Store
	var err error
	if cfg.Credentials != "" {
		store, err = question.NewStore(question.FireStore)
		log.Infof("successfully loaded firestore client")
	} else {
		store, err = question.NewStore(question.MemoryStorage)
	}
	if err != nil {
		return nil, err
	}
	return &questionServer{
		config: cfg,
		store:  store,
	}, nil
}

func (s *questionServer) WrapGrpcServer(g *grpc.Server) {
	options := []grpcweb.Option{
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	}
	s.WrappedGrpcServer = grpcweb.WrapServer(g, options...)
}
