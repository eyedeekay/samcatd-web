// +build webface

package samcatweb

import "log"

func (s *SAMWebConfig) Serve() {

}

func NewSAMWebConfigFromOptions() (*SAMWebConfig, error) {
	var s SAMWebConfig
	s.host = "127.0.0.1"
	s.port = "7656"
	return &s, nil
}

func Serve() {
	if webinterface, webinterfaceerr = samcatweb.NewSAMWebConfigFromOptions(); webinterfaceerr == nil {
		log.Println("Starting web interface")
		go webinterface.Serve()
	}
}
