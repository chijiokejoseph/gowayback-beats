package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zmb3/spotify"

	"github.com/chijiokejoseph/gowayback-beats/consts"

)

func Serve() (*spotify.Client, error) {
	// register the authenticate function to the /callback route
	http.HandleFunc("/callback", authenticate)

	// launch a go routine to listen to and serve requests across a port
	go func() {
		http.ListenAndServe(consts.Port, nil)
	}()

	// launch authentication url on a browser
	launchOauth()
	var client *spotify.Client

	// wait for 30 seconds for authentication to complete, if not return nil
	wait := 30
	select {
	case client = <-clientstream:
	case <-time.After(time.Duration(wait) * time.Second):
		msg := fmt.Sprintf("Authentication Timed Out after %d seconds\n", wait)
		fmt.Println(msg)
		return nil, fmt.Errorf("%s", msg)
	}
	return client, nil
}
