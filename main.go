package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var jIDCookie, _ = r.Cookie("JSESSIONID")
		if jIDCookie == nil {
			cookie := &http.Cookie{
				Name:     "JSESSIONID",
				Value:    "abc123",
				Secure:   true,
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			}
			http.SetCookie(w, cookie)
			w.Write([]byte(fmt.Sprintf("set cookie: %s", cookie)))
		}
	})

	http.ListenAndServe(":8080", mux)

}
