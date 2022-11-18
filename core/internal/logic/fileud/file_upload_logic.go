package fileud

import (
	"context"
	"gcloud/fileud/rpc/fileud"

	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadReply, err error) {
	uploadRet, err := l.svcCtx.FileudRPC.FileUpload(l.ctx, &fileud.FileUploadRequest{
		Name: req.Name,
		Hash: req.Hash,
		Path: req.Path,
		Ext:  req.Ext,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.FileUploadReply)
	resp.Identity = uploadRet.Identity
	resp.Ext = uploadRet.Ext
	resp.Name = uploadRet.Name
	resp.Msg = "success"
	return
}
