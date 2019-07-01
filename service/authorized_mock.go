package service

import "context"

type AuthorizedMock struct {
	StartFunc  func(context.Context) <-chan error
	VerifyRoleTokenFunc func(ctx context.Context, tok, act, res string) error
}

func (am *AuthorizedMock) Start(ctx context.Context) <-chan error {
	return am.StartFunc(ctx)
}

func (am *AuthorizedMock) VerifyRoleToken(ctx context.Context, tok, act, res string) error {
	return am.VerifyRoleTokenFunc(ctx, tok, act, res)
}
