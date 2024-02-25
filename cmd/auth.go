package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/zmb3/spotify/v2"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:     "auth",
	Aliases: []string{"a", "authenticate"},
	Short:   "authenticates spotify client",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			startAuth()
		} else {
			fmt.Println("you suck noob")
		}
	},
}

const redirectURI = "http://localhost:8080/callback"

var (
	auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithClientID("b18a0445ca6f4b57b67c285670079765"), spotifyauth.WithClientSecret("d26322f9ef6047dcac5bbdca63ea5120"),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeImageUpload))
	ch         = make(chan *spotify.Client)
	state      = "abc123"
	showDialog = spotifyauth.ShowDialog
)

func startAuth() {
	// first start an HTTP server
	http.HandleFunc("/callback", finishAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.User)
}

func finishAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// use the token to get an authenticated client
	client := spotify.New(auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}

func init() {
	rootCmd.AddCommand(authCmd)
}
