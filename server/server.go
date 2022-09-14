package server

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var (
	wsUpgrader = websocket.Upgrader{
		WriteBufferSize: 1024,
		ReadBufferSize:  1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func Run() {
	addr, hasAddr := os.LookupEnv("HEARTS_ADDR")
	if !hasAddr {
		addr = ":8080"
	}

	router := mux.NewRouter()
	router.Methods(http.MethodPost).Subrouter().HandleFunc("/game/", createGame)
	router.Methods(http.MethodGet).Subrouter().HandleFunc("/game/{id}/", playGame)

	err := http.ListenAndServe(addr, handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		log.Fatalln(err)
	}
}

func createGame(w http.ResponseWriter, _ *http.Request) {
	id := uuid.New()
	g := game{inboundEvents: make(chan inboundEvent, 1)}
	games[id] = g

	go g.run()

	_, err := fmt.Fprintln(w, id.String())
	logNonFatal(err)
}

func playGame(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintln(w, err)
		logNonFatal(err)
		return
	}
	g, gameFound := games[id]
	if !gameFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer func() { logNonFatal(conn.Close()) }()

	p := newPlayer(conn, uuid.New())
	g.connectPlayer(p)

	handleWebsocketMessages(p, g)
}
