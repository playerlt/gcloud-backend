package logic

import (
	"context"

	"gcloud/fileud/rpc/fileud"
	"gcloud/fileud/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDownloadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDownloadLogic {
	return &FileDownloadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDownloadLogic) FileDownload(in *fileud.FileDownloadRequest) (*fileud.FileDownloadReply, error) {
	// todo: add your logic here and delete this line

	return &fileud.FileDownloadReply{}, nil
}
