package main

import(
	"encoding/json"
	"fmt"

	//"fmt"
	"io"
	"io/ioutil"
	//"log"
	"net/http"
	"strconv"
	//"time"

	mux "github.com/julienschmidt/httprouter"
)

func GetPerson(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	id, err := strconv.Atoi(ps.ByName("postId"))
	HandleError(err)
	post := GetData(id)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		panic(err)
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
   fmt.Println("vaibhav")
	var post Post
	if err := json.Unmarshal(body, &post); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
	}

	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}

	CreateData(post)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}