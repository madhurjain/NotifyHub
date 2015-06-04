package main

import (
	"errors"
)

type Hub struct {
	apn chan *APN
	gcm chan *GCM
}

var hub *Hub

func (hub *Hub) route() {
	for {
		select {
		case apn := <-hub.apn:
			SendAPN(apn)
		case gcm := <-hub.gcm:
			SendGCM(gcm)
		}
	}
}

func InitHub() {
	hub = &Hub{
		apn: make(chan *APN),
		gcm: make(chan *GCM),
	}
	go hub.route()
}

func EnqueueNotification(notification Notification) (err error) {
	if hub == nil {
		err = errors.New("Notification hub not initialized")
		return
	}
	if notification.Apn != nil {
		hub.apn <- notification.Apn
	}
	if notification.Gcm != nil {
		hub.gcm <- notification.Gcm
	}
	return
}
