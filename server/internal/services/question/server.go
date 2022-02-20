package question

import (
	"context"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/qanda/server/internal/store/question"
	questionpb "github.com/qanda/server/pkg/question/v1alpha1"
	"google.golang.org/grpc"
)

// questionServer ...
type questionServer struct {
	config Config
	store  question.Store
	questionpb.UnimplementedquestionServiceServer
	*grpcweb.WrappedGrpcServer
}

// New creates instance of the Server
func New(ctx context.Context, cfg Config) (*questionServer, error) {

	return &questionServer{
		config: cfg,
		store:  question.NewStore(question.MemoryStorage),
	}, nil
}

func (s *questionServer) WrapGrpcServer(g *grpc.Server) {
	options := []grpcweb.Option{
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	}
	s.WrappedGrpcServer = grpcweb.WrapServer(g, options...)
}
