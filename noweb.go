// +build !webface

package samcatweb

import "log"

import "github.com/eyedeekay/sam-forwarder/manager"

func Serve(s *sammanager.SAMManager, hp ...string) {
	log.Println("not creating web interface because it wasn't added")
}
