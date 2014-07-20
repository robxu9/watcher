package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"
)

const (
	TokenLength = 15
)

type WebHook struct {
	Token      string
	Repository string
	Log        map[time.Time][]byte
}

func NewWebHook(repo *git.Repository) *WebHook {
	return &WebHook{
		Token:      GenerateRandomString(TokenLength),
		Repository: repo,
		Log:        make(map[time.Time][]byte),
	}
}

func NewWebHookWithToken(token string, repo *git.Repository) *WebHook {
	return &WebHook{
		Token:      token,
		Repository: repo,
		Log:        make(map[time.Time][]byte),
	}
}

func (w *WebHook) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	bte, err := ioutil.ReadAll(req)
	if err != nil {
		log.Printf("[warn] failed to read request body for logging: %s", err)
		w.Log[time.Now()] = []byte{}
	} else {
		w.Log[time.Now()] = bte
	}

	cmd := exec.Command("git", "pull")
	cmd.Dir = w.Repository

	err = cmd.Run()
	if err != nil {
		log.Printf("[warn] failed to git pull %s: %s", w.Repository, err)
		http.Error(rw, "failed to git pull", http.StatusInternalServerError)
		return
	}
}
