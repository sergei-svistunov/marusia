package marusia_test

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sergei-svistunov/marusia"
)

type testDialog struct{}

func (d testDialog) GetResponse(r marusia.Request) marusia.Response {
	return marusia.Response{
		Text: "Test",
		Tts:  "test",
	}
}

func TestMarusia_ServeHTTP(t *testing.T) {
	s := httptest.NewServer(marusia.New(testDialog{}))
	defer s.Close()

	data := bytes.NewBufferString(`{
  "meta": {
    "client_id": "MailRu-VC/1.0",
    "locale": "ru_RU",
    "timezone": "Europe/Moscow",
    "interfaces": {
      "screen": {}
    },
    "_city_ru": "Москва"
  },
  "request": {
    "command": "test",
    "original_utterance": "test",
    "type": "SimpleUtterance",
    "nlu": {
      "tokens": [
        "test"
      ],
      "entities": []
    }
  },
  "session": {
    "session_id": "168c7fd0-abd4-4981-87f2-bede9b088b33",
    "user_id": "12d892b013ddfac570f58a042144315d84019a387e504b862b46ca758048e9e3",
    "skill_id": "2f7a8484-28ab-11eb-adc1-0242ac120002",
    "new": true,
    "message_id": 0
  },
  "version": "1.0"
}`)

	resp, err := s.Client().Post(s.URL+"/", "application/json;charset=UTF-8", data)
	if err != nil {
		t.Fatalf("Cannot get response: %v", err)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Cannot read response body: %v", err)
	}
	defer resp.Body.Close()

	got := strings.TrimSpace(string(respData))
	expected := `{"response":{"text":"Test","tts":"test","end_session":false},"session":{"session_id":"168c7fd0-abd4-4981-87f2-bede9b088b33","message_id":0,"user_id":"12d892b013ddfac570f58a042144315d84019a387e504b862b46ca758048e9e3"},"version":"1.0"}`
	if got != expected {
		t.Fatalf("Invalid response data:\n     got: %s\nexpected: %s", got, expected)
	}
}
