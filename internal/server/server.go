package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diazharizky/go-graphql-bootstrap/config"
	"github.com/diazharizky/go-graphql-bootstrap/internal/app"
	"github.com/diazharizky/go-graphql-bootstrap/internal/handlers"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func init() {
	config.Global.SetDefault("server.host", "0.0.0.0")
	config.Global.SetDefault("server.port", "8080")
}

func Run(appCtx *app.Ctx) {
	schema := graphql.MustParseSchema(handlers.New(appCtx))

	http.Handle("/", &relay.Handler{
		Schema: schema,
	})

	host := config.Global.GetString("server.host")
	port := config.Global.GetString("server.port")
	addr := fmt.Sprintf("%s:%s", host, port)

	log.Printf("Server is listening on %s!", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error unable to run the server: %s", err.Error())
	}
}
