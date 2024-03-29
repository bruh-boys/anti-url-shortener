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
	w.Header().Set("allow-access-control-origin", "*")
	w.Header().Set("allow-access-origin", "*")
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

	req, _ := http.NewRequest("GET", string(urlVal), nil)

	respReq, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, "the url doesnt exist", http.StatusBadRequest)
		return
	}
	response := map[string]string{"url": respReq.Header.Get("Location")}
	api := r.URL.Query().Get("api")
	if api == "true" {
		json.NewEncoder(w).Encode(response)

		return
	}
	temp, _ := template.ParseFiles("src/view/index.html")
	if err := temp.Execute(w, response); err != nil {
		log.Println(err)
	}

}
