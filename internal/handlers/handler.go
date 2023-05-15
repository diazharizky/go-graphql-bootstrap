package handlers

import (
	"github.com/diazharizky/go-graphql-bootstrap/internal/app"
)

type handler struct {
	appCtx *app.Context
}

func NewHandler(appCtx *app.Context) (string, *handler) {
	return schema.String(), &handler{
		appCtx: appCtx,
	}
}
