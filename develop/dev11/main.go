package main

import (
	"log"
	"main/handler"
	"main/middleware"
	"main/service"
	"net/http"
	"sync"
)

func main() {

	serveMux := http.NewServeMux()
	store := service.NewStore(&sync.Mutex{}, make(map[int]service.EventCalendar))
	storeServer := handler.NewStoreServer(store)

	serveMux.HandleFunc("/create_event", storeServer.HandlerCreateEvent)
	serveMux.HandleFunc("/update_event", storeServer.HandlerUpdateEvent)
	serveMux.HandleFunc("/delete_event", storeServer.HandlerDeleteEvent)
	serveMux.HandleFunc("/events_for_day", storeServer.HandlerEventsForDay)
	serveMux.HandleFunc("/events_for_week", storeServer.HandlerEventsForWeek)
	serveMux.HandleFunc("/events_for_month", storeServer.HandlerEventsForMonth)

	login := middleware.Logging(serveMux)

	err := http.ListenAndServe("localhost:8080", login)
	if err != nil {
		log.Fatal(err)
	}
}
