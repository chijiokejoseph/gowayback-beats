package web

import (
	"github.com/zmb3/spotify"
	spotifyauth "github.com/zmb3/spotify/v2/auth"

	e "github.com/chijiokejoseph/gowayback-beats/env"
)

var (
	// url encoded random byte string
	codeVerifier = NewCode()
	// url encoded random byte string
	session = NewCode()
	// sha256 algorithm url encoded byte string
	codeChallenge = NewSecret(codeVerifier)
	// `Env` struct with environment values
	env = e.NewEnv()
	// spotify client permissions/scopes
	scopes = []string{
		spotify.ScopePlaylistModifyPrivate,
		spotify.ScopePlaylistModifyPublic,
	}
	// spotify authenticator
	auth = spotifyauth.New(spotifyauth.WithClientID(env.ClientID), spotifyauth.WithRedirectURL(env.RedirectUri), spotifyauth.WithScopes(scopes...))
	// channel for spotify.Client
	clientstream = make(chan *spotify.Client)
)
