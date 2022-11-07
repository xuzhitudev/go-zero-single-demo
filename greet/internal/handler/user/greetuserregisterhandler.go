package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-single-demo/greet/internal/logic/user"
	"go-zero-single-demo/greet/internal/svc"
	"go-zero-single-demo/greet/internal/types"
)

func GreetUserRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewGreetUserRegisterLogic(r.Context(), svcCtx)
		resp, err := l.GreetUserRegister(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
