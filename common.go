package samcatweb

import "github.com/eyedeekay/sam-forwarder/manager"

type SAMWebConfig struct {
	host    string
	port    string
	manager *sammanager.SAMManager
}
