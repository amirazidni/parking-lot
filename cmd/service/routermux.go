package service

import (
	"log"
	"net/http"
	"parking-lot/pkg/manage"
	"parking-lot/pkg/server"
	"parking-lot/pkg/util"

	"github.com/gorilla/mux"
)

type GatewayServer struct {
	Router *mux.Router
	Server *server.Server
}

func (g *GatewayServer) route() {
	g.Router.HandleFunc("/create_parking_lot/{value}", g.Server.CreateHandler()).Methods(http.MethodPost)
	g.Router.HandleFunc("/park/{value}/{attribute}", g.Server.ParkingHandler()).Methods(http.MethodPost)
	g.Router.HandleFunc("/status", g.Server.GetStatusHandler()).Methods(http.MethodGet)
	g.Router.HandleFunc("/leave/{value}", g.Server.LeaveParkHandler()).Methods(http.MethodPost)
	g.Router.HandleFunc("/cars_registration_numbers/colour/{value}", g.Server.GetCarsPlateHandler()).Methods(http.MethodGet)
}

func (g *GatewayServer) SetupServer() http.Handler {
	g.Router = mux.NewRouter()
	g.Server = server.NewServer(manage.NewManager())
	g.route()
	return g.Router
}

func (g *GatewayServer) StartServer() {
	log.Default().Println("Starting API server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", g.SetupServer()); err != nil {
		util.ErrorHandlerFatal(err, "failed to start server")
	}
}
