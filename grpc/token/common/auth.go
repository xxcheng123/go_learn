package common

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

/**
* @Author: xxcheng
* @Email developer@xxcheng.cn
* @Date: 2024/1/3 15:04
 */

type UserToken struct {
	Username string
	Password string
}

func (u *UserToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"username": u.Username,
		"password": u.Password,
	}, nil
}
func (u *UserToken) RequireTransportSecurity() bool {
	return false
}
func (u *UserToken) Pass(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("not user")
	}
	if md["username"][0] != u.Username || md["password"][0] != u.Password {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}
	return nil
}
