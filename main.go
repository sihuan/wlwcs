package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlePostCmd(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	cmd := request.Form.Get("CMD")

	fmt.Printf("POST form-urlencoded: CMD=%s\n", cmd)

	var resp string
	if cmd == "1" {
		resp = `{"CMD":"2","VILL":"BECFC1F5B4E500000000","NUMB":"001","FORM":"1"}`
	} else if cmd == "3" {
		resp = `{"CMD":"4","RETURN":"OK"}`
	}
	fmt.Fprintf(writer, resp)
}

func main() {

	http.HandleFunc("/cmd", handlePostCmd)
	fmt.Println("Running at port 6011 ...")

	err := http.ListenAndServe(":6011", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
