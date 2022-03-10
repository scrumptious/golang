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
	return &Goodbye{}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Whhhoooppss", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Goodbye %d", d)
}
