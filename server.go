package main

import (
	"encoding/json"
	"github.com/madhurjain/notifyhub/broker"
	apn "github.com/madhurjain/notifyhub/services/apn"
	gcm "github.com/madhurjain/notifyhub/services/gcm"
	"log"
	"net/http"
	"os"
)

const (
	endpoint = "/notify"
)

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	result, err := broker.Push(r.Body)
	if err != nil {
		errorJSON, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorJSON)
		return
	}
	resultJSON, err := json.Marshal(result)
	w.Write(resultJSON)
}

// Initialize the http server which accepts
// incoming notification JSON
func main() {

	apnService := apn.Initialize(apn.SANDBOX_GATEWAY, os.Getenv("APN_CERT_PATH"), os.Getenv("APN_KEY_PATH"))
	gcmService := gcm.Initialize(os.Getenv("GCM_API_KEY"))
	broker.AddServices(apnService, gcmService)

	httpHost := os.Getenv("HOST")
	httpPort := os.Getenv("PORT")

	if httpPort == "" {
		httpPort = "8080"
	}

	http.HandleFunc(endpoint, notificationHandler)
	log.Printf("NotifyHub listening on %s:%s\n", httpHost, httpPort)
	http.ListenAndServe(httpHost+":"+httpPort, nil)
}
