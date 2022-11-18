package logic

import (
	"context"
	"gcloud/core/helper"
	"gcloud/core/models"

	"gcloud/fileud/rpc/fileud"
	"gcloud/fileud/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileUploadLogic) FileUpload(in *fileud.FileUploadRequest) (*fileud.FileUploadReply, error) {
	rp := &models.RepositoryPool{
		Identity: helper.UUID(),
		Name:     in.Name,
		Hash:     in.Hash,
		Path:     in.Path,
		Ext:      in.Ext,
		Size:     in.Size,
	}
	resp := new(fileud.FileUploadReply)
	err := l.svcCtx.Engine.
		Select("identity", "name", "hash", "path", "ext", "size", "created_at", "updated_at").
		Create(rp).Error
	if err != nil {
		resp.Msg = "error"
		return nil, err
	}
	resp.Identity = rp.Identity
	resp.Ext = rp.Ext
	resp.Name = rp.Name
	resp.Msg = "success"
	return resp, nil
}
