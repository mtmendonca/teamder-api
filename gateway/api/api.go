package api

import (
	"log"
	"net/http"
	"time"

	"github.com/mtmendonca/teamder-api/common/grpc/account"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/mtmendonca/teamder-api/gateway/middleware"
	"github.com/urfave/negroni"
)

// Service encapsulates the api and its dependencies
type Service struct {
	GRPCAccountClient account.AccountServiceClient
	GraphqlHandler    *relay.Handler
	Cfg               *Config
	Log               *logrus.Logger
}

// Start fires up the graphql server
func Start(s *Service, addr string) {
	// Router
	r := mux.NewRouter()

	// Handlers
	r.HandleFunc("/healthcheck", s.healthcheck).Methods("GET")
	r.HandleFunc("/login", s.login).Methods("POST")
	r.Handle("/graphql", negroni.New(
		negroni.HandlerFunc(middleware.JWTMiddleware(s.Cfg.JWTSecret).HandlerWithNext),
		negroni.HandlerFunc(middleware.ExtractUser),
		negroni.Wrap(s.GraphqlHandler),
	))

	//Negroni
	n := negroni.Classic()
	n.UseFunc(middleware.AddHeaders)
	n.UseHandler(middleware.Cors(r))

	// Http server
	srv := &http.Server{
		Handler:      n,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
