# A library for developing Mail.ru Marusia external skill

## How to install
```shell script
go get github.com/sergei-svistunov/marusia
```

## How to use
```go
package main

import (
	"log"
	"net/http"

	"github.com/sergei-svistunov/marusia"
)

type exampleDialog struct{}

func (d exampleDialog) GetResponse(r marusia.Request) marusia.Response {
	if r.Session.New {
		return marusia.Response{
			Text: "Hello",
			Tts:  "hello",
		}
	}

	switch r.Request.Command {
	case "ping":
		return marusia.Response{
			Text: "Pong",
			Tts:  "pong",
		}
	default:
		return marusia.Response{
			Text: "Unknown command",
			Tts:  "unknown command",
		}
	}
}

func main() {
	m := marusia.New(exampleDialog{})

	if err := http.ListenAndServe(":8080", m); err != nil {
		log.Fatalf("Cannot serve HTTP: %v", err)
	}
}

```