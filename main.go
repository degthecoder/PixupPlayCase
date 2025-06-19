package main

import (
	"PixupPlayCaseTrial/src/app"
	"PixupPlayCaseTrial/src/handlers/event_handler"
	"PixupPlayCaseTrial/src/handlers/player_handler"
	"PixupPlayCaseTrial/src/lib/make_handle"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/event", make_handle.MakeHandle(event_handler.EventHandler))
	http.HandleFunc("/players", make_handle.MakeHandle(player_handler.GetAllPlayers))
	http.HandleFunc("/wallet/", make_handle.MakeHandle(player_handler.GetPlayerWallet))

	srv := &http.Server{
		Addr:        app.Settings.Host + ":" + app.Settings.Port,
		ReadTimeout: 30 * time.Second,
		//Handler:      myCors.corsWrapper(),
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Backend server starting server on port: %s\n", app.Settings.Port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Println(err.Error())
	}

	defer app.DisconnectDb()
}
