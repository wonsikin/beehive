package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wonsikin/beehive/src/db/mongo"
	"github.com/wonsikin/beehive/src/scheme"
)

// PostMessage handles message posting
func (e *Entry) PostMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	msg := new(scheme.Message)
	err = json.Unmarshal(data, msg)
	if err != nil {
		http.Error(w, "Must be JSON Data", http.StatusNotAcceptable)
		return
	}

	// TODO xxxx
	err = mongo.NewMsgCollection().Insert(msg)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OK"))
}
