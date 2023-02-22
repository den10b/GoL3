package service

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

type EventCalendar struct {
	Date time.Time
	Mes  string
}

type StoreServer struct {
	m     *sync.Mutex
	store map[int]EventCalendar
}

func NewStore(m *sync.Mutex, store map[int]EventCalendar) *StoreServer {
	return &StoreServer{m: m, store: store}
}

func (ss *StoreServer) CreateEvent(date time.Time, mes string) int {
	event := EventCalendar{date, mes}

	ss.m.Lock()
	defer ss.m.Unlock()

	id := len(ss.store)

	for {
		if reflect.DeepEqual(ss.store[id], EventCalendar{}) {
			ss.store[id] = event
			return id
		}
		id++
	}
}

func (ss *StoreServer) UpdateEvent(id int, date time.Time, mes string) (EventCalendar, error) {
	ss.m.Lock()
	defer ss.m.Unlock()

	if reflect.DeepEqual(ss.store[id], EventCalendar{}) {
		return EventCalendar{}, fmt.Errorf("503: такое событие отсутствует")
	}

	event := EventCalendar{date, mes}

	ss.store[id] = event

	return ss.store[id], nil
}

func (ss *StoreServer) DeleteEvent(id int) error {
	ss.m.Lock()
	defer ss.m.Unlock()

	if reflect.DeepEqual(ss.store[id], EventCalendar{}) {
		return fmt.Errorf("503: такое событие отсутствует")
	}

	delete(ss.store, id)

	return nil
}

func (ss *StoreServer) EventsForDay(date time.Time, days int) ([]EventCalendar, error) {
	var result []EventCalendar

	for _, event := range ss.store {
		if event.Date.Sub(date) >= time.Duration(days*time.Now().Day()) {
			result = append(result, event)
		}
	}
	if len(result) == 0 {
		return []EventCalendar{}, fmt.Errorf("503: такое событие отсутствует")
	}

	return result, nil
}
