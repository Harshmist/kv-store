package network

import (
	"fmt"
	"kv-store/store"
	"net/http"
	"strconv"
	"strings"
)

func HttpHandleFuncs() {

	http.HandleFunc("/list", ListData)
	http.HandleFunc("/get/", GetData)
	http.HandleFunc("/post/", Post)
	http.HandleFunc("/delete/", Delete)
	http.ListenAndServe(":8001", nil)

}

//HTTP HandleFuncs

func ListData(w http.ResponseWriter, r *http.Request) {

	for i := 1; i < len(store.Data)+1; i++ {
		if store.Data[i] != "" {
			fmt.Fprintf(w, strconv.FormatInt(int64(i), 10)+": ")
			fmt.Fprintf(w, store.Data[i]+"\n")
		}
	}

}

func GetData(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")

	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	key, err := strconv.Atoi(parts[2])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if store.Data[key] == "" {
		fmt.Fprintf(w, fmt.Sprint(key)+" empty")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, store.Data[key]+"\n")

	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")

	if len(parts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Allowing POST to take multiple words seperated by underscores "_"

	var value string
	valArr := strings.Split(parts[2], "_")
	if len(valArr) > 1 {
		value = strings.Join(valArr, " ")
	} else {
		value = parts[2]
	}
	store.PostChannel <- value
}

func Delete(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")

	if len(parts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key, err := strconv.Atoi(parts[2])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, store.Data[key]+" deleted")
	store.DelChannel <- key

}
