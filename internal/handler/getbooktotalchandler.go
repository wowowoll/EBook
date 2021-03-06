package handler

import (
	"net/http"

	"EBook/internal/logic"
	"EBook/internal/svc"
	"EBook/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func getBookTotalCHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetBookTotalcReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetBookTotalCLogic(r.Context(), ctx)
		resp, err := l.GetBookTotalC(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
