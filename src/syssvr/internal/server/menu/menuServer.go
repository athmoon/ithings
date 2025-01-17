// Code generated by goctl. DO NOT EDIT!
// Source: sys.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/logic/menu"
	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"
)

type MenuServer struct {
	svcCtx *svc.ServiceContext
	sys.UnimplementedMenuServer
}

func NewMenuServer(svcCtx *svc.ServiceContext) *MenuServer {
	return &MenuServer{
		svcCtx: svcCtx,
	}
}

func (s *MenuServer) MenuCreate(ctx context.Context, in *sys.MenuCreateReq) (*sys.Response, error) {
	l := menulogic.NewMenuCreateLogic(ctx, s.svcCtx)
	return l.MenuCreate(in)
}

func (s *MenuServer) MenuIndex(ctx context.Context, in *sys.MenuIndexReq) (*sys.MenuIndexResp, error) {
	l := menulogic.NewMenuIndexLogic(ctx, s.svcCtx)
	return l.MenuIndex(in)
}

func (s *MenuServer) MenuUpdate(ctx context.Context, in *sys.MenuUpdateReq) (*sys.Response, error) {
	l := menulogic.NewMenuUpdateLogic(ctx, s.svcCtx)
	return l.MenuUpdate(in)
}

func (s *MenuServer) MenuDelete(ctx context.Context, in *sys.MenuDeleteReq) (*sys.Response, error) {
	l := menulogic.NewMenuDeleteLogic(ctx, s.svcCtx)
	return l.MenuDelete(in)
}
