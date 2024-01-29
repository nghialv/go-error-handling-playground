package rpc

import (
	"context"
	"log"

	"github.com/nghialv/go-error-handling-playground/pkg/cerror"
	"google.golang.org/grpc"
)

func LogUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err == nil {
			return resp, nil
		}

		if cerr, ok := cerror.As(err); ok {
			log.Printf("log the request with cerror: %v", cerr)
			return nil, err
		}

		log.Println("warn: detected an error that was not wrapped with a cerror")
		return nil, err
	}
}
