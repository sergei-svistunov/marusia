// A library for developing Mail.ru Marusia external skill.
package marusia

import (
	"encoding/json"
	"log"
	"net/http"
)

var Log Logger = &log.Logger{}

type Marusia struct {
	dialog Dialog
}

type Dialog interface {
	GetResponse(r Request) Response
}

type Logger interface {
	Printf(format string, v ...interface{})
}

func New(dialog Dialog) *Marusia {
	return &Marusia{
		dialog: dialog,
	}
}

func (m *Marusia) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Max-Age", "86400")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var reqData Request
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		Log.Printf("Cannot parse request: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(FullResponse{
		Response: m.dialog.GetResponse(reqData),
		Session: ResponseSession{
			SessionId: reqData.Session.SessionId,
			MessageId: reqData.Session.MessageId,
			UserId:    reqData.Session.UserId,
		},
		Version: "1.0",
	}); err != nil {
		Log.Printf("Cannot encode response: %v", err)
	}
}
