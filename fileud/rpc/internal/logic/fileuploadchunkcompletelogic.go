package logic

import (
	"context"

	"gcloud/fileud/rpc/fileud"
	"gcloud/fileud/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkCompleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileUploadChunkCompleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkCompleteLogic {
	return &FileUploadChunkCompleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileUploadChunkCompleteLogic) FileUploadChunkComplete(in *fileud.FileUploadChunkCompleteRequest) (*fileud.FileUploadChunkCompleteReply, error) {
	// todo: add your logic here and delete this line

	return &fileud.FileUploadChunkCompleteReply{}, nil
}
