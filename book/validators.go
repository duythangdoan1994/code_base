package book

import (
	"net/http"
	"encoding/xml"
	"github.com/luciandd/core/common"
	"encoding/json"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	var res common.ResponseData

	res.Status = status
	res.Message = common.ResponseMessage(status)
	res.Data = payload

	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondXML(w http.ResponseWriter, status int, payload interface{}) {
	var res common.ResponseData

	res.Status = status
	res.Message = common.ResponseMessage(status)
	res.Data = payload

	response, err := xml.MarshalIndent(res, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, status int, message string) {
	var res common.ResponseData
	rescode := common.ResponseMessage(status)
	res.Status = status
	res.Message = rescode
	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}


