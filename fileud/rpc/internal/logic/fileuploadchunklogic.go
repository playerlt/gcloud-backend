package logic

import (
	"context"

	"gcloud/fileud/rpc/fileud"
	"gcloud/fileud/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadChunkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileUploadChunkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadChunkLogic {
	return &FileUploadChunkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileUploadChunkLogic) FileUploadChunk(in *fileud.FileUploadChunkRequest) (*fileud.FileUploadChunkReply, error) {
	// todo: add your logic here and delete this line

	return &fileud.FileUploadChunkReply{}, nil
}
