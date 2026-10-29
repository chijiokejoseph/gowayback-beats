package consts

import "time"

const (
	URL = "https://billboard.com/charts/hot-100/"
	DATE_FORMAT = time.DateOnly
)

const AUTH_URL = "https://accounts.spotify.com"

type Link string

const (
	Authorize Link   = "https://accounts.spotify.com/authorize"
	Token     Link   = "https://accounts.spotify.com/api/token"
	Local     Link   = "http://localhost:8080"
	Port      string = ":8080"
	Header    string = "application/x-www-form-urlencoded"
)