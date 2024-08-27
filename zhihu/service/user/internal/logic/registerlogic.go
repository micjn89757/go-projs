package logic

import (
	"context"
	"os"
	"time"
	"user/internal/code"
	"user/internal/model"
	"user/internal/svc"
	"user/service"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 注册逻辑
type RegisterLogic struct {
	ctx   	context.Context
	svcCtx 	*svc.ServiceContext  // 一些依赖，例如usermodel的操作
	logger  *zap.Logger
}


func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	//TODO: logger配置分离出去
	core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(os.Stderr), zapcore.DebugLevel)
	logger := zap.New(core)
	return &RegisterLogic{
		ctx: ctx,
		svcCtx: svcCtx,
		logger: logger,
	}
}


func (l *RegisterLogic) Register(in *service.RegisterRequest) (*service.RegisterResponse, error) {
	// TODO: 校验
	if len(in.Username) == 0 {
		return nil, code.RegisterNameEmpty
	}


	ret, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username: in.Username,
		Mobile: in.Mobile,
		Avatar: in.Avatar,
		CreateTime:	time.Now(),
		UpdateTime: time.Now(),
	})

	if err != nil {
		l.logger.Error("register req err:", zap.Error(err))
		return nil, err
	}

	userId, err := ret.LastInsertId()
	if err != nil {
		l.logger.Error("LastInsertId error", zap.Error(err))
		return nil, err 
	}

	return &service.RegisterResponse{UserId: userId}, nil
}


