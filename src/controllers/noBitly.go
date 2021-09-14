package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"text/template"
)

func NoBitly(w http.ResponseWriter, r *http.Request) {
	GetProxy()

	if r.Body == http.NoBody {
		log.Println(proxy)
		http.Error(w, "the request is empty >:(", http.StatusBadRequest)
		return
	}

	r.ParseForm()

	urlVal := r.FormValue("url")
	if urlVal == "" {
		var a map[string]string
		json.NewDecoder(r.Body).Decode(&a)
		urlVal = a["url"]
	}
	if proxy != nil {

		u := fmt.Sprintf("%s:%s/", proxy.IP, proxy.Port)
		urlProxy, _ := url.Parse(u)
		http.DefaultClient.Transport = &http.Transport{Proxy: http.ProxyURL(urlProxy)}
	}
	log.Println(urlVal)
	req, err := http.NewRequest("GET", string(urlVal), nil)
	if err != nil {
		http.Error(w, "the url doesnt exist", http.StatusBadRequest)
		return
	}
	resp, _ := http.DefaultTransport.RoundTrip(req)

	temp, _ := template.ParseFiles("src/view/index.html")
	if err := temp.Execute(w, map[string]string{"Url": resp.Header.Get("Location")}); err != nil {
		log.Println(err)
	}

}
