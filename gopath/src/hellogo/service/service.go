package service

import (
	//"hellogo/util"
	"net/http"
)

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
	default:
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("unknown action\n"))
	}
	return
}

func ping(query map[string][]string, res http.ResponseWriter) {
	res.Write([]byte("Pong\n"))
	return
}
