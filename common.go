package samcatweb

import "github.com/eyedeekay/sam-forwarder/manager"

type SAMWebConfig struct {
	host    string
	port    string
	lang    string
	title   string
	manager *sammanager.SAMManager
}
