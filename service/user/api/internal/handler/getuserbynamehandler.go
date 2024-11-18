package handler

import (
	"net/http"

	"LibSystem/service/user/api/internal/logic"
	"LibSystem/service/user/api/internal/svc"
	"LibSystem/service/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserByNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserByNameRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetUserByNameLogic(r.Context(), svcCtx)
		resp, err := l.GetUserByName(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
