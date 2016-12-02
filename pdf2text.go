package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func convertPDF(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "post only", http.StatusMethodNotAllowed)
		return
	}
	if r.Body == nil {
		http.Error(w, "Body is missing", http.StatusBadRequest)
		return
	}
	tmpFileName := "/tmp/pdftotext.tmp" + time.Now().String()
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	err := ioutil.WriteFile(tmpFileName, bodyBytes, 0600)
	if err != nil {
		http.Error(w, "Failed to open the file for writing", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tmpFileName)
	// log.Printf("File uploaded successfully.")

	body, err := exec.Command("pdftotext", "-nopgbrk", "-enc", "UTF-8", tmpFileName, "-").Output()
	if err != nil {
		log.Printf("pdftotext error: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, string(body))
	// log.Printf("File successfully converted.")
}

func main() {
	http.HandleFunc("/", convertPDF)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
