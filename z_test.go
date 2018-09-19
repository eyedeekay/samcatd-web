package samcatweb

import "testing"

import "github.com/eyedeekay/sam-forwarder/manager"

func TestAll(t *testing.T) {
	if manager, err := sammanager.NewSAMManagerFromOptions(
		sammanager.SetManagerHost("127.0.0.1"),
		sammanager.SetManagerSAMHost("127.0.0.1"),
		sammanager.SetManagerPort("8080"),
		sammanager.SetManagerSAMPort("7656"),
		sammanager.SetManagerWebHost("127.0.0.1"),
		sammanager.SetManagerWebPort("7958"),
		sammanager.SetManagerFilePath("../sam-forwarder/etc/samcatd/tunnels.ini"),
	); err == nil {
		go Serve(manager, "127.0.0.1", "7958")
	} else {
		t.Fatal(err)
	}
}
