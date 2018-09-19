package samcatweb

import (
	"net/http"
)

import "github.com/eyedeekay/sam-forwarder/manager"

type SAMWebConfig struct {
	host         string
	port         string
	lang         string
	title        string
	pages        []*pagestring
	manager      *sammanager.SAMManager
	localService *http.ServeMux
}

type pagestring struct {
	dir      string
	url      string
	apiurl   string
	desc     string
	id       string
	class    string
    title    string
    lang     string
	manager  *sammanager.SAMManager
	children []*pagestring
}
