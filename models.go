package main

import "time"

type UrbanDictionaryDefinition struct {
	Definition  string    `json:"definition,omitempty"`
	Permalink   string    `json:"permalink,omitempty"`
	ThumbsUp    int       `json:"thumbs_up,omitempty"`
	Author      string    `json:"author,omitempty"`
	Word        string    `json:"word,omitempty"`
	Defid       int       `json:"defid,omitempty"`
	CurrentVote string    `json:"current_vote,omitempty"`
	WrittenOn   time.Time `json:"written_on,omitempty"`
	Example     string    `json:"example,omitempty"`
	ThumbsDown  int       `json:"thumbs_down,omitempty"`
}

type UrbanDictionaryResponse struct {
	List []UrbanDictionaryDefinition `json:"list,omitempty"`
}

