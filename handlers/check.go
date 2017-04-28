package handlers

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"
)

var domain = flag.String("domain", "goreportcard.com", "Domain used goreportcard installation")

// ReportHandler handles the request for checking a repo
func ReportHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	repo := r.FormValue("repo")
	if repo == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Could not get the repository"))
		return
	}

	log.Printf("Checking repo %q...", repo)

	resp, err := newChecksResp(repo)
	if err != nil {
		log.Println("ERROR: from newChecksResp:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Could not download the repository.`))
		return
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Println("ERROR: could not marshal json:", err)
		http.Error(w, err.Error(), 500)
		return
	}

	t := template.Must(template.New("report.html").Delims("[[", "]]").ParseFiles("templates/report.html"))
	t.Execute(w, map[string]interface{}{
		"repo":     repo,
		"response": string(respBytes),
		"loading":  false,
		"domain":   domain,
		"badge":    badgePath(resp.Grade, ""),
	})
}
