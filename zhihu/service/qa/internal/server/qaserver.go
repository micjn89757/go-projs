// Code generated by goctl. DO NOT EDIT.
// Source: qa.proto

package server

import (
	"context"

	"qa/internal/logic"
	"qa/internal/svc"
	"qa/qa"
)

type QaServer struct {
	svcCtx *svc.ServiceContext
	qa.UnimplementedQaServer
}

func NewQaServer(svcCtx *svc.ServiceContext) *QaServer {
	return &QaServer{
		svcCtx: svcCtx,
	}
}

func (s *QaServer) Ping(ctx context.Context, in *qa.Request) (*qa.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
