package gapi

import (
	"context"
	"database/sql"

	"github.com/gitsuki/simplebank/pb"
	"github.com/gitsuki/simplebank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "account does not exist: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to find account: %s", err)
	}

	err = util.CheckPassword(req.GetPassword(), user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "invalid password: %s", err)
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %s", err)
	}

	rsp := &pb.LoginUserResponse{
		User:                  convertUser(user),
		AccessToken:           accessToken,
		RefreshToken:          "",
		SessionId:             "",
		AccessTokenExpiresAt:  timestamppb.Now(),
		RefreshTokenExpiresAt: timestamppb.Now(),
	}
	return rsp, nil
}
