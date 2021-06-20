package main

import (
	"encoding/json"
	"net/http"
)

// Post struct - data stucture
// this struct is going to represents request and response
// and return this Post using this 'json' format as part of payload
// as able to encode and decode this struct , to json format and from json format
//  we need to add json tag i.e `json:""` from encoding/json package

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{ID: 1, Title: "Go setup", Text: "This is first go article"}}
}

// getPosts return all posts

func getPosts(res http.ResponseWriter, req *http.Request) {
	// this func going to return json so set header
	res.Header().Set("Content-Type", "application/json")
	// encode/marshal array into json format
	result, err := json.Marshal(posts)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

// addPost create a new post
func addPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	// parse the request body and unmarshal the json from the receive from req
	// and going create a post object and add that into a posts array

	var post Post
	err := json.NewDecoder(req.Body).Decode(&post) // decode the json into post variable
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	post.ID = len(posts) + 1
	posts = append(posts, post)
	res.WriteHeader(http.StatusOK)
	result, _ := json.Marshal(post)
	res.Write(result)
}
