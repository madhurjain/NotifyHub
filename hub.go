package main

import (
	"errors"
)

type Hub struct {
	apn chan *APN
	gcm chan *GCM
	ws  chan *WS
}

var hub *Hub

func (hub *Hub) run() {
	for {
		select {
		case apn := <-hub.apn:
			SendAPN(apn)
		case gcm := <-hub.gcm:
			SendGCM(gcm)
		case ws := <-hub.ws:
			SendWS(ws)
		}
	}
}

func InitHub() {
	hub = &Hub{
		apn: make(chan *APN),
		gcm: make(chan *GCM),
		ws:  make(chan *WS),
	}
	go hub.run()
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
	if notification.Ws != nil {
		hub.ws <- notification.Ws
	}
	return
}
