package public_repo

import (
	"context"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/public_repo/rpc/public_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileListLogic {
	return &PublicFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFileListLogic) PublicFileList(req *types.PublicFileListRequest) (resp *types.PublicFileListReply, err error) {
	resp = new(types.PublicFileListReply)
	pflRet, err := l.svcCtx.PublicRepoRPC.PublicFileList(l.ctx, &public_repo.PublicFileListRequest{
		Id:   req.Id,
		Page: int32(req.Page),
		Size: int32(req.Size),
	})
	if err != nil {
		resp.Msg = "error"
		return
	}
	for i := 0; i < len(pflRet.List); i++ {
		uf := new(types.PublicFile)
		uf.Identity = pflRet.List[i].Identity
		uf.ParentId = pflRet.List[i].ParentId
		uf.Name = pflRet.List[i].Name
		uf.Ext = pflRet.List[i].Ext
		uf.Id = pflRet.List[i].Id
		uf.RepositoryIdentity = pflRet.List[i].RepositoryIdentity
		uf.Size = pflRet.List[i].Size
		uf.Path = pflRet.List[i].Path
		uf.UpdatedAt = pflRet.List[i].UpdatedAt
		uf.Owner = pflRet.List[i].Owner
		resp.List = append(resp.List, uf)
	}

	resp.Msg = "success"
	resp.Count = pflRet.Count
	return
}
