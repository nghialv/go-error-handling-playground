package handler

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nghialv/go-error-handling-playground/app/fooapp/usecase"
	"github.com/nghialv/go-error-handling-playground/service"
)

type Handler struct {
	fooUC *usecase.FooUsecase
}

func (h *Handler) GetFoo(ctx context.Context, req *service.GetFooRequest) (*service.GetFooResponse, error) {
	out, err := h.fooUC.GetFoo(ctx, &usecase.GetFooInput{Id: req.Id})
	if err != nil {
		if errors.Is(err, usecase.ErrNotFound) {
			return nil, errors.Join(err, status.Error(codes.NotFound, "foo not found"))
		}

		return nil, errors.Join(err, status.Error(codes.Internal, ""))
	}

	return &service.GetFooResponse{
		Name: out.Name,
	}, nil
}
