package handler

import (
	"encoding/json"
	"fmt"
	"main/service"
	"net/http"
	"strconv"
	"time"
)

type StoreServer struct {
	store *service.StoreServer
}

func NewStoreServer(store *service.StoreServer) *StoreServer {
	return &StoreServer{store: store}
}

func (ss *StoreServer) HandlerCreateEvent(w http.ResponseWriter, r *http.Request) {
	_, date, mes, err := handlerDataPost(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	id := ss.store.CreateEvent(date, mes)
	RenderJSON(w, id)
}

func (ss *StoreServer) HandlerUpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, date, mes, err := handlerDataPost(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	event, err := ss.store.UpdateEvent(id, date, mes)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	RenderJSON(w, event)
}

func (ss *StoreServer) HandlerDeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, _, _, err := handlerDataPost(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	errDelete := ss.store.DeleteEvent(id)
	if errDelete != nil {
		http.Error(w, errDelete.Error(), 503)
		return
	}
	RenderJSON(w, "Событие удалено!")
}

func (ss *StoreServer) HandlerEventsForDay(w http.ResponseWriter, r *http.Request) {

	date := handlerDataGet(r)

	events, err := ss.store.EventsForDay(date, 0)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	RenderJSON(w, events)
}

func (ss *StoreServer) HandlerEventsForWeek(w http.ResponseWriter, r *http.Request) {
	date := handlerDataGet(r)

	events, err := ss.store.EventsForDay(date, 7)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	RenderJSON(w, events)
}

func (ss *StoreServer) HandlerEventsForMonth(w http.ResponseWriter, r *http.Request) {
	date := handlerDataGet(r)

	events, err := ss.store.EventsForDay(date, 30)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	RenderJSON(w, events)
}

func handlerDataPost(r *http.Request) (int, time.Time, string, error) {
	var id int
	var date time.Time
	var mes string

	idString := r.FormValue("id")
	if idString != "" {
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			return 0, time.Time{}, "", fmt.Errorf("400: некорректное число")
		}
		id = idInt
	}

	dateString := r.FormValue("date")
	if dateString != "" {
		dateString += "T00:00:00Z"
		dateTime, err := time.Parse(time.RFC3339, dateString)
		if err != nil {
			return 0, time.Time{}, "", fmt.Errorf("400: некоррректная дата")
		}
		date = dateTime
	}

	mes = r.FormValue("mes")

	return id, date, mes, nil
}

func handlerDataGet(r *http.Request) time.Time {
	dateF := r.FormValue("date") + "T00:00:00Z"
	date, err := time.Parse(time.RFC3339, dateF)
	if err != nil {
		fmt.Println(err)
	}
	return date
}

func RenderJSON(w http.ResponseWriter, v interface{}) {
	resultJSON := struct {
		Result interface{}
	}{Result: v}

	js, err := json.Marshal(&resultJSON)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
