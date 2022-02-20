package question

import (
	"context"
	"net/http"
	"strings"

	questionpb "github.com/castlemilk/qanda/server/pkg/question/v1alpha1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/qanda/server/internal/statik"
	"github.com/qanda/server/pkg/log"
	statikfs "github.com/rakyll/statik/fs"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *questionServer) Serve(ctx context.Context, logger *log.Logger, address string) error {

	gRPCServer := s.newgRPCServer(ctx, logger)
	mux, err := s.newHTTPHandler(ctx, address)
	if err != nil {
		return err
	}
	registerSwaggerUI(mux)
	s.WrapGrpcServer(gRPCServer)
	server := &http.Server{
		Addr:    address,
		Handler: s.dispatcher(ctx, gRPCServer, mux),
	}
	return server.ListenAndServe()
}

func registerSwaggerUI(mux *http.ServeMux) {

	statikFS, err := statikfs.New()
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(statikFS)
	sh := http.StripPrefix("/api/docs/", staticServer)
	mux.Handle("/api/docs/", sh)
}

// dispatch enables support for grpc and rest on same port
func (s *questionServer) dispatcher(ctx context.Context, grpcHandler http.Handler, httpHandler http.Handler) http.Handler {
	hf := func(w http.ResponseWriter, r *http.Request) {

		if s.IsAcceptableGrpcCorsRequest(r) || s.IsGrpcWebRequest(r) {
			log.Infof("grpcweb request")
			s.ServeHTTP(w, r)
		}
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			log.Infof("grpc request")
			grpcHandler.ServeHTTP(w, r)
		} else {
			log.Infof("http request")
			httpHandler.ServeHTTP(w, r)
		}
	}
	return allowCORS(h2c.NewHandler(http.HandlerFunc(hf), &http2.Server{}))
}

func (s *questionServer) newgRPCServer(ctx context.Context, logger *log.Logger) *grpc.Server {
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []grpc_zap.Option{
		grpc_zap.WithLevels(log.LevelFunc),
	}
	zapLogger := logger.Desugar()
	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLogger(zapLogger)
	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(zapLogger, opts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(zapLogger, opts...),
		),
	)
	questionpb.RegisterquestionServiceServer(grpcServer, s)
	reflection.Register(grpcServer)
	return grpcServer
}

func (s *questionServer) newHTTPHandler(ctx context.Context, address string) (*http.ServeMux, error) {
	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption("application/json+pretty", &runtime.JSONPb{MarshalOptions: protojson.MarshalOptions{Indent: " "}}))
	if err := questionpb.RegisterquestionServiceHandlerFromEndpoint(ctx, gwmux, address, []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		return nil, err
	}
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	return mux, nil
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	// headers := []string{"Content-Type", "Accept", ""}
	w.Header().Set("Access-Control-Allow-Headers", "*")
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Infof("preflight request for %s", r.URL.Path)
	return
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in questionion systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
