package libManifest

import (
	"net/http"
	"net/url"

	cdn "github.comcast.com/dwill0849e/cdnClient"
	dash "github.comcast.com/viper-cog/libdash"
)

var manifestClient *http.Client

// Manifest represents the received mpd file
type Manifest struct {
	dash.Mpd
}

// Get fetches the MPD file using the provided URL
func (m Manifest) Get(u url.URL) (man *Manifest) {

	if manifestClient == nil {
		manifestClient = cdn.NewClient(3000)
	}
	return
}
