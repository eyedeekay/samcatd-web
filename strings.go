// +build webface

package samcatweb

type pagestring struct {
	dir    string
	url    string
	desc   string
	Render func() (ret string, err error)
}

var (
	samcatd_headline       = &pagestring{dir: "./", url: "index", desc: "SAMcatd Control Panel", Render: render_index}
	samcatd_ntcptun        = &pagestring{dir: "./", url: "ntcpserver", desc: "ntcp server tunnels", Render: render_ntcpserver}
	samcatd_httptun        = &pagestring{dir: "./", url: "httpserver", desc: "http/ntcp server tunnels", Render: render_ntcpserver_http}
	samcatd_ssutun         = &pagestring{dir: "./", url: "ssuserver", desc: "ssu server tunnels", Render: render_ssuserver}
	samcatd_ntcptun_client = &pagestring{dir: "./", url: "ntcpclient", desc: "ntcp client tunnels", Render: render_ntcpclient}
	samcatd_ssutun_client  = &pagestring{dir: "./", url: "ssuclient", desc: "ssu client tunnels", Render: render_ssuclient}
)

func render_header() string {
	return ""
}

func render_footer() string {
	return ""
}

func render_div(input string, class string, id ...string) string {
	return ""
}

func render_index() (ret string, err error) {

}

func render_ntcpserver() (ret string, err error) {

}
func render_ntcpserver_http() (ret string, err error) {

}

func render_ssuserver() (ret string, err error) {

}

func render_ntcpclient() (ret string, err error) {

}

func render_ssuclient() (ret string, err error) {

}
