package handler

import (
	"net/http"

	"EBook/internal/logic"
	"EBook/internal/svc"
	"EBook/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func uploadBookHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BookUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUploadBookLogic(r.Context(), ctx)
		resp, err := l.UploadBook(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
