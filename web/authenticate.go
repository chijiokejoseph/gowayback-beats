package web

import (
	"fmt"
	"net/http"

	"github.com/pkg/browser"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

// authenticate is a `http.HandleFunc` that handles authenticating a user to a spotify account
// It generates a token and confirms that the token matches for the request session before
// creating a spotify client and sending across a channel
func authenticate(w http.ResponseWriter, r *http.Request) {
	// generate token from session byte strings and codeVerifier byte strings
	token, err := auth.Token(r.Context(), session, r, oauth2.SetAuthURLParam("code_verifier", codeVerifier))
	if err != nil {
		msg := fmt.Sprintf("Could not get access token. \nThe following error occurred \n%s", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// verify that the `state` value from the request session string
	if st := r.FormValue("state"); st != session {
		msg := "Invalid Session. Could not locate page"
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	// create a http.Client from the request context and token
	newClient := auth.Client(r.Context(), token)

	// generate a new spotify.Client from the http.Client
	sp := spotify.NewClient(newClient)

	// send spotify.Client through the channel
	clientstream <- &sp
}

// launchOauth launches a generated spotify authentication url to the browser
func launchOauth() {
	// generate authentication url to be launched on the browser
	// codeChallenge is a secret or password string generated from the S256 algorithm
	link := auth.AuthURL(session, oauth2.SetAuthURLParam("code_challenge_method", "S256"), oauth2.SetAuthURLParam("code_challenge", codeChallenge))

	// open authentication link on the default browser of the client
	browser.OpenURL(link)
}
