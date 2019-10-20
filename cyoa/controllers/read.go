package controllers

import (
	// "html/template"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var story map[string]Chapter

func init() {
	story = getMap()
}

func Read(w http.ResponseWriter, r *http.Request) {
	chapter := getChapter(r.URL)
	// spew.Dump(w)
	if content, ok := story[chapter]; ok {
		t, err := template.ParseFiles("resources/templates/chapter.gohtml")
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, content)
	} else {
		t, err := template.ParseFiles("resources/templates/404.gohtml")
		if err != nil {
			fmt.Println(err)
		}
		redirection := struct {
			Ref  string
			Text string
		}{
			Ref:  "/",
			Text: "Go home",
		}
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, redirection)
	}
}

func getChapter(u *url.URL) string {
	path := u.RequestURI()
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func getMap() map[string]Chapter {
	storyMap := make(map[string]Chapter)
	story, err := ioutil.ReadFile("resources/story.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(story, &storyMap)
	if err != nil {
		panic(err)
	}
	return storyMap
}
