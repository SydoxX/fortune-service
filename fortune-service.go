package main

import (
  //"fmt"
  "net/http"
  "os/exec"
  "encoding/json"
)

func main() {
  http.HandleFunc("/", FortuneServer)
  http.ListenAndServe(":80", nil)
}

func FortuneServer(w http.ResponseWriter, r *http.Request) {
  fortuneRaw, _ := exec.Command("fortune").Output()
  fortune := map[string]string{
    "fortune": string(fortuneRaw),
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(fortune)
}
