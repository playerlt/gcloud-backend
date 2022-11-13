package logic

import (
	"context"

	"gcloud/fileud/rpc/fileud"
	"gcloud/fileud/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(in *fileud.FileUploadPrepareRequest) (*fileud.FileUploadPrepareReply, error) {
	// todo: add your logic here and delete this line

	return &fileud.FileUploadPrepareReply{}, nil
}
