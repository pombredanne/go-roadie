package main

import (
  "os"
  "log"
  "net/http"
  "github.com/hailiang/gosocks"
)

func main() {
  dialSocksProxy := socks.DialSocksProxy(socks.SOCKS4A, "127.0.0.1:9050")
  transport := &http.Transport{Dial: dialSocksProxy}
  httpClient := &http.Client{Transport: transport}

  resp, err := httpClient.Get("http://ifconfig.me")
  if err != nil {
    log.Print(err)
    os.Exit(1)
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Print(err)
    os.Exit(1)
  }
  log.Print(body)
}
