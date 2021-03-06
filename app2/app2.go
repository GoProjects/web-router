// app2
package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	//"net/http/httputil"
	//"net/url"
	//"io/ioutil"
)

func main() {
	fmt.Println("Version 2")
	ipStr, _ := IP()
	fmt.Println(ipStr)

	handlerMap := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Run App2.")
		printInterfaces(w)
	}
	/*
		handlerMapPlus1 := func(w http.ResponseWriter, r *http.Request) {

			fmt.Fprintln(w, "Run App2.")
			printInterfaces(w)
			fmt.Fprintln(w)

			resp, _ := http.Get("http://localhost:81")
			read, _ := ioutil.ReadAll(resp.Body)
			fmt.Fprintln(w, fmt.Sprintf("%s", read))
		}

		http.HandleFunc("/app2/plus1", handlerMapPlus1) */
	http.HandleFunc("/", handlerMap)

	err := http.ListenAndServe("0.0.0.0:82", nil)
	fmt.Println(err)
}

func printInterfaces(w http.ResponseWriter) {
	n, _ := net.InterfaceAddrs()
	for _, i := range n {
		fmt.Fprintln(w, i.String())
	}
}

func IP() (string, error) {
	r, err := http.Get("http://api.ipify.org/?format=json")
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	info := struct {
		Ip string `json:"ip"`
	}{}

	err = json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		return "", err
	}
	return info.Ip, nil
}
