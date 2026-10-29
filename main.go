package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chijiokejoseph/gowayback-beats/consts"
	"github.com/chijiokejoseph/gowayback-beats/inputs"
	"github.com/chijiokejoseph/gowayback-beats/service"
	"github.com/chijiokejoseph/gowayback-beats/utils"
	"github.com/chijiokejoseph/gowayback-beats/web"

)

func main() {
	// get date from user input
	date := utils.Must(inputs.InputDate())

	// initialize string to contain app printouts
	printout := ""

	// get billboard songs for the matched dates
	songs := utils.Must(service.Scrape(date))

	// append songs metadata to printout
	for _, song := range songs {
		printout += fmt.Sprintf("%s by %s\n", song.Title, song.Artist)
	}

	// get result datapath
	datapath := utils.Must(GetPath())
	// construct filename to store metadata of songs gotten from billboard
	filename := fmt.Sprintf("chart-%s.txt", date.Format(consts.DATE_FORMAT))
	newpath := filepath.Join(datapath, filename)

	// open filepath
	file := utils.Must(os.OpenFile(newpath, os.O_CREATE+os.O_WRONLY, os.ModePerm))
	defer func() {
		utils.Handle(file.Close())
	}()
	// write printout to file opened from filepath
	_ = utils.Must(file.WriteString(printout))

	// sign in to user's spotify account and extract its spotify client
	client, err := web.Serve()

	if err != nil {
		return
	}

	// create new playlist to hold songs
	playlist := utils.Must(service.NewPlaylist(client, &date))
	// get song tracks from spotify
	tracks := utils.Must(service.SearchAllSongs(client, songs))
	// add songs to the user's spotify account through its client
	utils.Handle(service.AddToPlaylist(client, playlist, tracks))
}

// GetPath creates a `data` folder in the working
// directory and return it
//
// Returns:
//   - string: path to the `data` directory as a string
//   - error: error if operation fails or nil on success
func GetPath() (string, error) {
	// get working directory
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// construct data path
	datapath := filepath.Join(dir, "data")
	// create directory
	if err := os.MkdirAll(datapath, os.ModePerm); err != nil {
		return "", err
	}
	
	return datapath, nil
}
