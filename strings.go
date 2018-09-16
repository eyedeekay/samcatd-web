// +build webface

package samcatweb

type pagestring struct {
    dir string
	url  string
	desc string
    fields func
}

var (
	samcatd_headline       = &pagestring{url: "index", desc: "SAMcatd Control Panel"}
	samcatd_ntcptun        = &pagestring{url: "ntcpserver", desc: "ntcp server tunnels"}
	samcatd_httptun        = &pagestring{url: "httpserver", desc: "http/ntcp server tunnels"}
	samcatd_ssutun         = &pagestring{url: "ssuserver", desc: "ssu server tunnels"}
	samcatd_ntcptun_client = &pagestring{url: "ntcpclient", desc: "ntcp client tunnels"}
	samcatd_ssutun_client  = &pagestring{url: "ssuclient", desc: "ssu client tunnels"}
)
