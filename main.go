package main

import (
  // "fmt"
  "io/ioutil"
  "log"
  // "encoding/xml"
  "net/http"
  "net/url"
)

var MTAURL = "http://web.mta.info/status/serviceStatus.txt"

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
  
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

  log.Println(string(body))
}