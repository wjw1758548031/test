package logic

import (
	"context"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"test/internal/svc"
	"test/pb/test"
)

type TestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Test 测试接口
func (l *TestLogic) Test(in *test.TestRequest) (*test.TestResponse, error) {
	l.Logger.WithFields(logx.LogField{Key: "id:", Value: in.Id}, logx.LogField{Key: "node_id:", Value: l.svcCtx.Config.NodeId}).Info("Test")
	return &test.TestResponse{Result: cast.ToString(l.svcCtx.Config.NodeId)}, nil
}
