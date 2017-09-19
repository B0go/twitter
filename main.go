package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/B0go/twitter/model"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {

		config := oauth1.NewConfig(r.Header.Get("X-ConsumerKey"), r.Header.Get("X-ConsumerSecret"))
		token := oauth1.NewToken(r.Header.Get("X-AccessToken"), r.Header.Get("X-AccessTokenSecret"))

		client := twitter.NewClient(config.Client(oauth1.NoContext, token))

		var post model.PostDTO

		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			fmt.Fprintln(w, err)
		}

		client.Statuses.Update(post.Message, nil)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
