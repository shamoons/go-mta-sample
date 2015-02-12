package main

import (
  // "fmt"
  "io/ioutil"
  "log"
  "encoding/xml"
  "net/http"
  "net/url"
)

var MTAURL = "http://web.mta.info/status/serviceStatus.txt"

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
  
}

type Service struct {
  RequestTimestamp string `xml:"timestamp"`
  Subway struct {
    Line []struct {
      Name string `xml:"name"`
      Status string `xml:"status"`
      Date string `xml:"date"`
      Time string `xml:"string"`
    } `xml:"line"`
  } `xml:"subway"`
}

// Main function for getting MTA service status
func main() {
  URL, err := url.Parse(MTAURL)
  check(err)

  resp, err := http.Get(URL.String())
  check(err)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  check(err)

  service := &Service{}
  err = xml.Unmarshal(body, service)
  check(err)

  log.Println(service)
}