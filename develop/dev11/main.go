package main

import (
	"log"
	"main/handler"
	"main/middleware"
	"main/service"
	"net/http"
	"sync"
)

func configureRoutes(serveMux *http.ServeMux, storeServer *handler.StoreServer) {
	serveMux.HandleFunc("/create_event", storeServer.HandlerCreateEvent)
	serveMux.HandleFunc("/update_event", storeServer.HandlerUpdateEvent)
	serveMux.HandleFunc("/delete_event", storeServer.HandlerDeleteEvent)
	serveMux.HandleFunc("/events_for_day", storeServer.HandlerEventsForDay)
	serveMux.HandleFunc("/events_for_week", storeServer.HandlerEventsForWeek)
	serveMux.HandleFunc("/events_for_month", storeServer.HandlerEventsForMonth)
}

func main() {

	serveMux := http.NewServeMux()
	store := service.NewStore(&sync.Mutex{}, make(map[int]service.EventCalendar))
	storeServer := handler.NewStoreServer(store)
	configureRoutes(serveMux, storeServer)

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./tmpl/index.html")
	})

	login := middleware.Logging(serveMux)

	log.Fatal(http.ListenAndServe("localhost:8080", login))
}
