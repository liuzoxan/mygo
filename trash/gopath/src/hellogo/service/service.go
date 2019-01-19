package service

import (
	//"hellogo/util"
	"encoding/json"
	"net/http"
	"runtime"
)

type Ping struct {
	Pong string
}

type Status struct {
	GOOS         string
	GORoot       string
	NumGoroutine int
}

func DoAdmin(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	//user, hasUser := query["user"]
	//pass, hasPass := query["password"]
	//if !hasUser || !hasPass || util.Base64([]byte(user[0])) != "emV1c3JlZGlz" ||
	//	util.Base64([]byte(pass[0])) != "YmVDYXJlZnVs" {
	//	res.WriteHeader(http.StatusUnauthorized)
	//	res.Write([]byte("auth failed\n"))
	//	return
	//}

	action, hasAction := query["action"]
	if !hasAction {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("unknown action\n"))
		return
	}
	switch action[0] {
	case "ping":
		ping(query, res)
	case "status":
		status(res)
	default:
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("unknown action\n"))
	}
	return
}

func ping(query map[string][]string, res http.ResponseWriter) {
	data := Ping{Pong: "pong"}
	json.NewEncoder(res).Encode(data)
	return
}

func status(res http.ResponseWriter) {
	data := Status{
		GOOS:         runtime.GOOS,
		GORoot:       runtime.GOROOT(),
		NumGoroutine: runtime.NumGoroutine(),
	}
	json.NewEncoder(res).Encode(data)
	return
}
