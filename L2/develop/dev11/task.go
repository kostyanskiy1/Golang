package main

import (
	"L2/develop/dev11/event"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

/*
=== HTTP server ===
Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.
В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

var eventMap = make(map[int]event.Event)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", createEvent)
	mux.HandleFunc("/update_event", updateEvent)
	mux.HandleFunc("/delete_event", deleteEvent)
	mux.HandleFunc("/events_for_day", eventsForDay)
	mux.HandleFunc("/events_for_week", eventsForWeek)
	mux.HandleFunc("/events_for_month", eventsForMonth)

	log.Printf("Server launched")
	err := http.ListenAndServe(":8181", mux)
	if err != nil {
		log.Fatal(err)
	}
}

// createEvent - GET request to server
func createEvent(w http.ResponseWriter, r *http.Request) {
	var e event.Event
	var err error

	e.EventID, err = event.GetUserIntFromURL(r, "event_id")
	if err != nil {
		errorHandler(w, "400")
		return
	}

	if _, ok := eventMap[e.EventID]; ok {
		log.Printf("Event with id %v already exis", e.EventID)
		errorHandler(w, "503")
		return
	}

	e.UserID, err = event.GetUserIntFromURL(r, "user_id")
	if err != nil {
		errorHandler(w, "400")
		return
	}

	e.Date, err = event.GetDateFromURL(r)
	if err != nil {
		errorHandler(w, "400")
		return
	}

	e.Text = r.URL.Query().Get("text")
	eventMap[e.EventID] = e

	log.Printf("Event %v created", e.EventID)
	outputAllUserEvents(w, e.UserID)
}

// updateEvent - update event data if it exists
func updateEvent(w http.ResponseWriter, r *http.Request) {
	var e event.Event
	var err error

	e.EventID, err = event.GetUserIntFromURL(r, "event_id")
	if err != nil {
		errorHandler(w, "400")
		return
	}

	if _, ok := eventMap[e.EventID]; !ok {
		log.Printf("Event with id %v doesn't exist", e.EventID)
		errorHandler(w, "503")
		return
	}

	e.UserID, err = event.GetUserIntFromURL(r, "user_id")
	if err != nil {
		errorHandler(w, "400")
		return
	}

	e.Date, err = event.GetDateFromURL(r)
	if err != nil {
		errorHandler(w, "400")
		return
	}

	e.Text = r.URL.Query().Get("text")

	eventMap[e.EventID] = e

	log.Printf("Event %v updated", e.EventID)
	outputAllUserEvents(w, e.UserID)
}

// deleteEvent - delete event by event_id
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	var e event.Event
	var err error

	e.EventID, err = event.GetUserIntFromURL(r, "event_id")
	if err != nil {
		return
	}

	e.UserID = eventMap[e.EventID].UserID

	delete(eventMap, e.EventID)
	log.Printf("Event %v deleted\n", e.EventID)
	outputAllUserEvents(w, e.UserID)
}

// eventsForDay - output event that planning on current day
func eventsForDay(w http.ResponseWriter, r *http.Request) {
	eventForHours(w, 24)
}

// eventsForWeek - event that planning on one week
func eventsForWeek(w http.ResponseWriter, r *http.Request) {
	eventForHours(w, 24*7)
}

// eventsForWeek - event that planning on one week
func eventsForMonth(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	eventForHours(w, 24*7*float64(t.Day()))
}

func eventForHours(w http.ResponseWriter, hours float64) {
	var eventArray []event.Event
	for _, v := range eventMap {
		if v.Date.Sub(time.Now()).Hours() < hours {
			eventArray = append(eventArray, v)
		}
	}

	unmarshalledEvent, err := json.Marshal(eventArray)
	if err != nil {
		log.Printf("Error in unmarshalling event: %s", err)
		errorHandler(w, "500")
	}

	_, err = w.Write(unmarshalledEvent)
	if err != nil {
		log.Printf("Error in responce: %s", err)
		errorHandler(w, "500")
	}
}

// outputAllUserEvents - output all event for user from get request
func outputAllUserEvents(w http.ResponseWriter, userID int) {
	var eventArray []event.Event

	for _, v := range eventMap {
		if v.UserID == userID {
			eventArray = append(eventArray, v)
		}
	}

	unmarshalledEvent, err := json.Marshal(eventArray)
	if err != nil {
		log.Printf("Error in unmarshalling event: %s", err)
		errorHandler(w, "500")
	}

	_, err = w.Write(unmarshalledEvent)
	if err != nil {
		log.Printf("Error in responce: %s", err)
		errorHandler(w, "500")
	}
}

func errorHandler(w http.ResponseWriter, errorCode string) {
	unmarshalledCode, _ := json.Marshal(errorCode)
	w.Write(unmarshalledCode)
}
