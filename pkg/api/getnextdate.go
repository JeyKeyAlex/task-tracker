package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/JKasus/go_final_project/pkg/constants"
)

var daysInterval int

//var weeksInterval int

func checkSymbol(symbol string) bool {
	switch symbol {
	case constants.DaySign, constants.YearSign:
		return true
	default:
		return false
	}
}

func afterNow(date time.Time, now time.Time) bool {
	return date.After(now)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		err := fmt.Errorf("value %s can not be empty", repeat)
		return "", err
	}

	partsRepeat := strings.Split(repeat, " ")
	if len(partsRepeat) > 2 {
		err := errors.New("repeat parameter can not be consists of more than 2 elements")
		return "", err
	}

	if !checkSymbol(partsRepeat[0]) {
		err := fmt.Errorf("value %s can not be used is the rule of repeating", partsRepeat[0])
		return "", err
	}

	if partsRepeat[0] == constants.DaySign {
		if len(partsRepeat) != 2 {
			err := fmt.Errorf("repeat parameter with first value %s can not consists of less than 2 elements", partsRepeat[0])
			return "", err
		}
	} else if partsRepeat[0] == constants.YearSign {
		if len(partsRepeat) != 1 {
			err := fmt.Errorf("repeat parameter with first value %s can not consists of more than 1 element", partsRepeat[0])
			return "", err
		}
	}

	if partsRepeat[0] == constants.DaySign {
		days, err := strconv.Atoi(partsRepeat[1])
		if err != nil {
			err = fmt.Errorf("value %s can not convert to int", partsRepeat[1])
			return "", err
		}
		daysInterval = days
		if daysInterval > constants.DaysMaxValue || daysInterval < constants.DaysMinValue {
			err := fmt.Errorf("%d: ivalid value for days", days)
			return "", err
		}
	}
	//} else if partsRepeat[0] == constants.WeekSign {
	//	weeks, err := strconv.Atoi(partsRepeat[1])
	//	if err != nil {
	//		err = fmt.Errorf("value %s can not convert to int", partsRepeat[1])
	//		return "", err
	//	}
	//	weeksInterval = weeks
	//	if weeksInterval > constants.WeeksMaxValue || weeksInterval < constants.WeeksMinValue {
	//		err := fmt.Errorf("%d: ivalid value for weeks", weeks)
	//		return "", err
	//	}
	//}

	startDate, err := time.Parse(constants.DateFormat, dstart)
	if err != nil {
		err = fmt.Errorf("Failed to parse start date: %v", err)
		return "", err
	}

	switch partsRepeat[0] {
	case constants.DaySign:
		for {
			startDate = startDate.AddDate(0, 0, daysInterval)
			if afterNow(startDate, now) {
				break
			}
		}
	case constants.YearSign:
		for {
			startDate = startDate.AddDate(1, 0, 0)
			if afterNow(startDate, now) {
				break
			}
		}
	}

	return startDate.Format(constants.DateFormat), nil
}

func getNextDayHandler(w http.ResponseWriter, r *http.Request) {

	var nowDate time.Time
	nowParam := r.URL.Query().Get("now")
	if nowParam != "" {
		var err error
		nowDate, err = time.Parse(constants.DateFormat, nowParam)
		if err != nil {
			nowDate = time.Now()
		}
	} else {
		nowDate = time.Now()
	}
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	nextDate, err := NextDate(nowDate, date, repeat)
	if err != nil {
		log.Printf("Failed to run NextDate function: %v", err)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}
