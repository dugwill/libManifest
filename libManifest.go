package libManifest

import (
	"net/url"

	dash "github.comcast.com/viper-cog/libdash"
)

// Manifest represents the received mpd file
type Manifest struct {
	dash.Mpd
}

// Get fetches the MPD file using the provided URL
func (m Manifest) Get(u url.URL) (man *Manifest) {

	return
}
