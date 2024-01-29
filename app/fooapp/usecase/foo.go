package usecase

import "context"

type FooUsecase struct {
}

type GetFooInput struct {
	Id string
}

type GetFooOutput struct {
	Id   string
	Name string
}

func (u *FooUsecase) GetFoo(ctx context.Context, input *GetFooInput) (*GetFooOutput, error) {
	return nil, ErrNotFound
}
