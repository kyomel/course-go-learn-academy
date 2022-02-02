package main

import (
	"Web-Application/rps"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// html := `<strong>Hello, world</strong>`
	// w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, html)
	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)
	out, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	// log.Println(winner, computerChoice, roundResult)
}

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
