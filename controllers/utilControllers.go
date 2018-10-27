package controllers

import (
	u "github.com/Kronin-Cloud/estate-inventory/utils"
	"net/http"
)

var Root = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, u.Message(true,"Root"))
	return
}

var Metrics = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, u.Message(true,"Metrics"))
	return
}
var Healthz = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, u.Message(true,"Healthz"))
	return
}