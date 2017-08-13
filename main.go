package main

import (
  _"context"
  "encoding/json"
  "fmt"
  _"log"
  "net/http"

  _"golang.org/x/oauth2"
  "github.com/google/go-github/github"
  "github.com/rjz/githubhook"
)

func event(w http.ResponseWriter, req *http.Request) {
  secret := []byte("githubsecret")
  hook, err := githubhook.Parse(secret, req)
  if err != nil {
    fmt.Println("Request error", err)
  }

  evt := github.PullRequestReviewCommentEvent{}
  if err := json.Unmarshal(hook.Payload, &evt); err != nil {
    log.Debug("Invalid JSON?", err)
  }

  fmt.Println(evt.Comment.GetBody())
}

func main() {
  http.HandleFunc("/ev", event)
  http.ListenAndServe(":8090", nil)
}