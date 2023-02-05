package post

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
)

type GetPostListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetPostListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetPostListLogic {
	return &GetPostListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetPostListLogic) GetPostList(req *types.PostListReq) (resp *types.PostListResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetPostList(l.ctx,
		&core.PostListReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Name:     req.Name,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.PostListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.PostInfo{
				BaseInfo: types.BaseInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Trans:  l.svcCtx.Trans.Trans(l.lang, v.Name),
				Status: v.Status,
				Sort:   v.Sort,
				Name:   v.Name,
				Code:   v.Code,
				Remark: v.Remark,
			})
	}
	return resp, nil
}