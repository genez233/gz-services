package service

import (
	"context"
	"errors"
	"github.com/genez233/go-utils/md5"
	"gorm.io/gorm"
	pb "gz-services/proto"
	"gz-services/servers/user_server/global"
	"gz-services/servers/user_server/model"
)

type UserService struct {
	ctx *context.Context
	db  *gorm.DB
}

func (s UserService) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if in == nil {
		return nil, errors.New("param error")
	}

	var user model.User
	err := s.db.Where("id = ?", in.Id).Find(&user).Error
	return &pb.GetUserResponse{
		Id:          user.ID,
		Email:       user.Email,
		UserName:    user.UserName,
		Avatar:      user.Avatar,
		CreatedTime: user.CreatedTime.Unix(),
		UpdatedTime: user.UpdatedTime.Unix(),
		State:       user.State,
	}, err
}

func (s UserService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	if in == nil {
		return nil, errors.New("param error")
	}

	pwd := md5.EncodeMD5(in.Password + global.App.PwdSalt)

	var user model.User
	err := s.db.Where("email = ? and password = ?", in.Acc, pwd).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Id:          user.ID,
		Email:       user.Email,
		UserName:    user.UserName,
		Avatar:      user.Avatar,
		CreatedTime: user.CreatedTime.Unix(),
		UpdatedTime: user.UpdatedTime.Unix(),
		State:       user.State,
	}, nil
}
