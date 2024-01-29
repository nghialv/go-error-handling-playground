package rpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GRPCErrorConversionUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err == nil {
			return resp, nil
		}

		// Find gRPC error which was included at the handler layer.
		if s, ok := status.FromError(err); ok {
			return nil, s.Err()
		}

		// Show log to notify developer to include the appropriate gRPC error at the handler layer.
		log.Println("warn: detected an error that was not wrapped with a gRPC error")
		return nil, status.Error(codes.Internal, "")
	}
}
