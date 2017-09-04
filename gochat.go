package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "The address of the app")
	flag.Parse()
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "gochat.html"})
	http.Handle("/room", r)

	go r.run()

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`
				<html>
					<head>
						<title>gochat</title>
					</head>
					<body>
						Let's do chat!!:D
					</body>
				</html>
			`))
		})
	*/

	//Starting the Web-Server
	log.Println("Starting the Web-Server:D Port: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
