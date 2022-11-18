package private_repo

import (
	"context"
	"gcloud/core/internal/svc"
	"gcloud/core/internal/types"
	"gcloud/private_repo/rpc/private_repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {
	resp = new(types.UserFileListReply)
	uflRet, err := l.svcCtx.PrivateRepoRPC.UserFileList(l.ctx, &private_repo.UserFileListRequest{
		Id:           req.Id,
		UserIdentity: userIdentity,
		Page:         int32(req.Page),
		Size:         int32(req.Size),
	})
	if err != nil {
		resp.Msg = "error"
		return
	}
	for i := 0; i < len(uflRet.List); i++ {
		uf := new(types.UserFile)
		uf.Identity = uflRet.List[i].Identity
		uf.ParentId = uflRet.List[i].ParentId
		uf.Name = uflRet.List[i].Name
		uf.Ext = uflRet.List[i].Ext
		uf.Id = uflRet.List[i].Id
		uf.RepositoryIdentity = uflRet.List[i].RepositoryIdentity
		uf.Size = uflRet.List[i].Size
		uf.Path = uflRet.List[i].Path
		uf.UpdatedAt = uflRet.List[i].UpdatedAt
		resp.List = append(resp.List, uf)
	}

	for i := 0; i < len(uflRet.DeletedList); i++ {
		uf := new(types.DeletedUserFile)
		uf.Identity = uflRet.DeletedList[i].Identity
		uf.ParentId = uflRet.DeletedList[i].ParentId
		uf.Name = uflRet.DeletedList[i].Name
		uf.Ext = uflRet.DeletedList[i].Ext
		uf.Id = uflRet.DeletedList[i].Id
		uf.RepositoryIdentity = uflRet.DeletedList[i].RepositoryIdentity
		uf.Size = uflRet.DeletedList[i].Size
		uf.Path = uflRet.DeletedList[i].Path
		uf.DeletedAt = uflRet.DeletedList[i].DeletedAt
		resp.DeletedList = append(resp.DeletedList, uf)
	}

	resp.Count = uflRet.Count
	resp.Msg = "success"
	return
}
