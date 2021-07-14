package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pedrazzani/golang/backend-api/model"
)

func ListOrders(w http.ResponseWriter, r *http.Request) {
	response, _ := json.Marshal(values)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(response))
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	id--
	if id >= 0 && id < len(values) {
		response, _ := json.Marshal(values[id])
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(response))
	} else {
		w.WriteHeader(404)
	}

}

func PostOrder(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var o model.Order
	json.Unmarshal(body, &o)
	values = append(values, o)
	w.WriteHeader(201)
}

var values []model.Order

func init() {
	values = make([]model.Order, 5)
	values[0] = model.Order{1, 100}
	values[1] = model.Order{2, 200}
	values[2] = model.Order{3, 300}
	values[3] = model.Order{4, 400}
	values[4] = model.Order{5, 500}
}
