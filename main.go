package main

import (
	"log"
	"net/http"
)

func main() {
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

	//Starting Web-Server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
