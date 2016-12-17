package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/valyala/fasthttp"
)

func convertPDF(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != "POST" {
		ctx.Error("post only", 405)
		return
	}
	if ctx.PostBody() == nil {
		ctx.Error("Body is missing", 400)
		return
	}
	tmpFileName := "/tmp/pdftotext.tmp" + time.Now().String()
	bodyBytes := ctx.PostBody()
	err := ioutil.WriteFile(tmpFileName, bodyBytes, 0600)
	if err != nil {
		ctx.Error("Failed to open the file for writing", 500)
		return
	}
	defer os.Remove(tmpFileName)
	// log.Printf("File uploaded successfully.")

	body, err := exec.Command("pdftotext", "-nopgbrk", "-enc", "UTF-8", tmpFileName, "-").Output()
	if err != nil {
		log.Printf("pdftotext error: %s", err)
		ctx.Error(err.Error(), 500)
	}
	fmt.Fprintf(ctx, string(body))
	// log.Printf("File successfully converted.")
}

func main() {
	h := convertPDF
	if err := fasthttp.ListenAndServe(":5000", h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
