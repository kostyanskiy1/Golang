package event

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Event struct {
	EventID int       `json:"event_id"`
	UserID  int       `json:"user_id"`
	Date    time.Time `json:"date"`
	Text    string    `json:"text"`
}

func GetUserIntFromURL(r *http.Request, key string) (int, error) {
	userID, err := strconv.Atoi(r.URL.Query().Get(key))
	if err != nil {
		log.Println(fmt.Sprintf("Error in getting parameter from URL: %s", err))
		return 0, err
	}
	return userID, nil
}

func GetDateFromURL(r *http.Request) (time.Time, error) {
	timeFromURl, err := time.Parse("02-01-2006", r.URL.Query().Get("date"))
	if err != nil {
		log.Println(fmt.Sprintf("Error in in time parcing: %s", err))
		return time.Date(2000, 01, 01, 0, 0, 0, 0, time.Local), err
	}
	return timeFromURl, nil
}
