package md5_generator

import (
	"log"
	"net/http"
)

func (a Services) Run() error {
	port := a.PORT

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
