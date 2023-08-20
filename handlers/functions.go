package handlers

import "log"

func apiLog(l *log.Logger, counter *uint, url *string, err error) {
	var status string
	if err == nil {
		status = "SUCCESS"
	} else {
		status = err.Error()
	}

	l.Printf("[%d] [%s] [%s]", *counter, *url, status)
}
