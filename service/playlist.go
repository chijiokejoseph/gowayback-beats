package service

import (
	"fmt"
	"os"
	"time"

	"github.com/zmb3/spotify"
)

func getUser(client *spotify.Client) (*spotify.PrivateUser, error) {
	user, err := client.CurrentUser()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't get the user's details. \nError: %s\n", err)
		return nil, err
	}
	fmt.Printf("Retrieved User: %s\n", user.DisplayName)
	return user, nil
}

func NewPlaylist(client *spotify.Client, timeline *time.Time) (*spotify.FullPlaylist, error) {
	user, err := getUser(client)
	if err != nil {
		return nil, err
	}
	month := timeline.Month().String()
	day := timeline.Day()
	year := timeline.Year()
	name := fmt.Sprintf("Timeline %s %d %d", month, day, year)
	desc := fmt.Sprintf("The memories from the year %d, the month of %s, right straight to your library.", year, month)
	playlist, err := client.CreatePlaylistForUser(user.ID, name, desc, false)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Created Playlist %s for user %s", playlist.Name, user.DisplayName)
	return playlist, nil
}

func SearchSong(client *spotify.Client, song *Song) (*spotify.FullTrack, error) {
	query := fmt.Sprintf("%s by %s", song.Title, song.Artist)
	result, err := client.Search(query, spotify.SearchTypeTrack)
	if err != nil {
		return nil, err
	}
	tracks := result.Tracks.Tracks
	firstTrack := tracks[0]
	return &firstTrack, nil
}

func SearchAllSongs(client *spotify.Client, songs []*Song) ([]*spotify.FullTrack, error) {
	tracks := make([]*spotify.FullTrack, len(songs))
	for i, song := range songs {
		track, err := SearchSong(client, song)
		if err != nil {
			return tracks, err
		}
		tracks[i] = track
	}
	return tracks, nil
}

func AddToPlaylist(client *spotify.Client, playlist *spotify.FullPlaylist, tracks []*spotify.FullTrack) error {
	trackIDs := make([]spotify.ID, len(tracks))
	for i, track := range tracks {
		trackIDs[i] = track.ID
	}
	_, err := client.AddTracksToPlaylist(playlist.ID, trackIDs...)
	if err != nil {
		return err
	}
	fmt.Printf("\nTracks have been added to your playlist %s\n", playlist.Name)
	return nil
}
