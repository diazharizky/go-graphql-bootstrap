package handlers

import (
	"log"

	"github.com/diazharizky/go-graphql-bootstrap/internal/app"
	"github.com/diazharizky/go-graphql-bootstrap/internal/schema"
)

type handler struct {
	appCtx *app.Context
}

func NewHandler(appCtx *app.Context) (string, *handler) {
	s, err := schema.String()
	if err != nil {
		log.Fatalf("Error unable to generate schema: %v", err)
	}

	return s, &handler{
		appCtx: appCtx,
	}
}
