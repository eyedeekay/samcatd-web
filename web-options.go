// +build webface

package samcatweb

import (
	"fmt"
	"strconv"
)

import "github.com/eyedeekay/sam-forwarder/manager"

//Option is a SAMWebConfig Option
type Option func(*SAMWebConfig) error

//SetHost sets the host of the SAMWebConfig's SAM bridge
func SetHost(s string) func(*SAMWebConfig) error {
	return func(c *SAMWebConfig) error {
		c.host = "127.0.0.1" //s
		return nil
	}
}

//SetPort sets the port of the SAMWebConfig's SAM bridge using a string
func SetPort(s string) func(*SAMWebConfig) error {
	return func(c *SAMWebConfig) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid TCP Server Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.port = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

func SetManager(s *sammanager.SAMManager) func(*SAMWebConfig) error {
	return func(c *SAMWebConfig) error {
		c.manager = s
		return nil
	}
}
