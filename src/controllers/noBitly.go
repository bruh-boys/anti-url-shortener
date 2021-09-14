package controllers

import (
	"fmt"
	"net/http"
	"net/url"
)

func NoBitly(w http.ResponseWriter, r *http.Request) {
	GetProxy()
	r.ParseForm()
	urlVal := r.FormValue("url")
	if proxy != nil {

		u := fmt.Sprintf("%s:%s/", proxy.IP, proxy.Port)
		urlProxy, _ := url.Parse(u)
		http.DefaultClient.Transport = &http.Transport{Proxy: http.ProxyURL(urlProxy)}
	}
	req, err := http.NewRequest("GET", string(urlVal), nil)
	if err != nil {
		http.Error(w, "the url doesnt exist", http.StatusBadRequest)
		return
	}
	resp, _ := http.DefaultTransport.RoundTrip(req)
	fmt.Println(resp.Header.Get("Location"))

}
