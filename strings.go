// +build webface

package samcatweb

import "github.com/eyedeekay/sam-forwarder/manager"

type pagestring struct {
	dir    string
	url    string
	desc   string
	id     string
	class  string
	Render func() (ret string, err error)
}

var (
	samcatd_headline = &pagestring{dir: "./",
		url: "index", desc: "SAMcatd Control Panel",
		id: "control_panel", class: "home,control",
		Render: render_index}
	samcatd_ntcptun = &pagestring{dir: "./server/",
		url: "ntcp", desc: "ntcp server tunnels",
		id: "ntcp_server", class: "server,ntcp",
		Render: render_ntcpserver}
	samcatd_httptun = &pagestring{dir: "./server/",
		url: "http", desc: "http/ntcp server tunnels",
		id: "http_server", class: "server,http",
		Render: render_ntcpserver_http}
	samcatd_ssutun = &pagestring{dir: "./server/",
		url: "ssu", desc: "ssu server tunnels",
		id: "ssu_server", class: "server,ssu",
		Render: render_ssuserver}
	samcatd_ntcptun_client = &pagestring{dir: "./client/",
		url: "ntcp", desc: "ntcp client tunnels",
		id: "ntcp_client", class: "client,ntcp",
		Render: render_ntcpclient}
	samcatd_ssutun_client = &pagestring{dir: "./client/",
		url: "ssu", desc: "ssu client tunnels",
		id: "ssu_client", class: "client,ssu",
		Render: render_ssuclient}
)

func render_div(input string, class string, id ...string) string {
	return ""
}

func render_index(s *sammanager.SAMManager) (ret string, err error) {
	return ""
}

func render_ntcpserver(s *sammanager.SAMManager) (ret string, err error) {
	return ""
}

func render_ntcpserver_http(s *sammanager.SAMManager) (ret string, err error) {
	return ""
}

func render_ssuserver(s *sammanager.SAMManager) (ret string, err error) {
	return ""
}

func render_ntcpclient(s *sammanager.SAMManager) (ret string, err error) {
	return ""
}

func render_ssuclient(s *sammanager.SAMManager) (ret string, err error) {
	return ""
}
