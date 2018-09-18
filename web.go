// +build webface

package samcatweb

import "log"

import "github.com/eyedeekay/sam-forwarder/manager"

func (s *SAMWebConfig) Serve() {

}

func NewSAMWebConfigFromOptions(opts ...func(*SAMWebConfig) error) (*SAMWebConfig, error) {
	var s SAMWebConfig
	s.host = "127.0.0.1"
	s.port = "7957"
	return &s, nil
}

func Serve(s *sammanager.SAMManager) {
	if webinterface, webinterfaceerr = samcatweb.NewSAMWebConfigFromOptions(); webinterfaceerr == nil {
		s.manager = s
		log.Println("Starting web interface")
		go webinterface.Serve()
	}
}
