package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Print("Goodbye World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Whhhoooppss", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Goodbye %s", d)
}
