package goduckgo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

type Gogo struct {
	Ads          interface{} `json:"ads"`
	Next         string      `json:"next"`
	Query        string      `json:"query"`
	QueryEncoded string      `json:"queryEncoded"`
	ResponseType string      `json:"response_type"`
	Results      []struct {
		Height    int    `json:"height"`
		Image     string `json:"image"`
		Source    string `json:"source"`
		Thumbnail string `json:"thumbnail"`
		Title     string `json:"title"`
		URL       string `json:"url"`
		Width     int    `json:"width"`
	} `json:"results"`
}

type Query struct {
	Keyword, P, S string
}

type regex_r struct {
	Regex, Body string
}

// Hit duckduckgo for results
// NOTE: this module for now can only get p=1 s=0 results
func Search(keyword Query) Gogo {

	if keyword.Keyword == "" {
		log.Fatal("No Query!")
		os.Exit(3)
	}
	if keyword.P == "" {
		keyword.P = "1"
	}

	if keyword.S == "" {
		keyword.S = "0"
	}

	url := "https://duckduckgo.com/"

	log.Print("Hitting DuckDuckGo for Token")

	//   First make a request to above URL, and parse out the 'vqd'
	//   This is a special token, which should be used in the subsequent request

	netreq := http.Client{
		Timeout: time.Second * 20, // Timeout is 20 Seconds
	}

	requrl := url + "?q=" + keyword.Keyword
	req, err := http.NewRequest(http.MethodGet, requrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:93.0) Gecko/20100101 Firefox/93.0")

	res, err := netreq.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	hunsen := regexfind(regex_r{Regex: `vqd=([\d-]+)\&`, Body: string(body)})

	if hunsen == "" {
		log.Fatal("Token Parsing Failed !")
		os.Exit(3)
	} else {
		log.Println("Obtained Token")
	}

	requrl = url + "i.js" + "?l=us-en&o=json&q=" + keyword.Keyword + "&" + hunsen + "f=,,,&v7exp=a&p=" + keyword.P + "&s=" + keyword.S
	log.Print("Hitting Url : " + requrl)

	resq, err := http.NewRequest(http.MethodGet, requrl, nil)
	resq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:93.0) Gecko/20100101 Firefox/93.0")
	resq.Header.Add("accept", "application/json, text/javascript, */* q=0.01")
	resq.Header.Add("authority", "duckduckgo.com")
	resq.Header.Add("sec-fetch-dest", "Empty")
	resq.Header.Add("x-requested-with", "XMLHttpRequest")
	resq.Header.Add("sec-fetch-site", "same-origin")
	resq.Header.Add("sec-fetch-mode", "cors")
	resq.Header.Add("referer", "https://duckduckgo.com/")
	resq.Header.Add("accept-language", "en-US,enq=0.9")
	if err != nil {
		log.Fatal(err)
	}
	resqs, getErr := netreq.Do(resq)
	if getErr != nil {
		log.Fatal(getErr)
	}

	brudy, readErr := ioutil.ReadAll(resqs.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	duckduck := Gogo{}
	jsonErr := json.Unmarshal(brudy, &duckduck)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	log.Println("\nHitting Url Success : " + requrl)
	return duckduck
}

func regexfind(hitin regex_r) string {
	gexgex := regexp.MustCompile(hitin.Regex)
	hunsen := gexgex.FindString(string(hitin.Body))
	return hunsen
}
