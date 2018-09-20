package samcatweb

import (
	"golang.org/x/time/rate"
	"net/http"
)

import "github.com/eyedeekay/sam-forwarder/manager"

type SAMWebConfig struct {
	host         string
	port         string
	lang         string
	title        string
	csspath      string
	cssstring    string
	jspath       string
	jsstring     string
	limiter      *rate.Limiter
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
