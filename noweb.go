// +build !webface

package samcatweb

import "log"

func Serve(s *sammanager.SAMManager, hp ...string) {
	log.Println("not creating web interface because it wasn't added")
}
