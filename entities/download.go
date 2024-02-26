package entities

import "github.com/zmb3/spotify/v2"

type DownloadEntity struct {
	files []spotify.PlaylistTrack
}

type Links struct {
	links []DownloadEntity
}

func GetDownloadLinks(playlist string) (*Links, error) {

}
