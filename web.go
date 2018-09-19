// +build webface

package samcatweb

import (
	"log"
	"net/http"
	"strings"
)

import "github.com/eyedeekay/sam-forwarder/manager"

func stringify(s *[]string) string {
	var p string
	for _, x := range *s {
		p += x + ","
	}
	r := strings.Replace(p, ",,", ",", -1)
	return r
}

func name(s string) string {
	for _, r := range strings.Split(s, "\n") {
		if strings.Contains(r, "name") {
			return strings.TrimPrefix("name=", r)
		}
	}
	return "NULL"
}

func (s *SAMWebConfig) populate() {
	for _, i := range *s.manager.List("") {
		s.pages[0].PopulateChild(name(i), i)
	}
	for _, i := range *s.manager.List("ntcpserver") {
		s.pages[1].PopulateChild(name(i), i)
	}
	for _, i := range *s.manager.List("httpserver") {
		s.pages[2].PopulateChild(name(i), i)
	}
	for _, i := range *s.manager.List("ssuserver") {
		s.pages[3].PopulateChild(name(i), i)
	}
	for _, i := range *s.manager.List("nctpclient") {
		s.pages[4].PopulateChild(name(i), i)
	}
	for _, i := range *s.manager.List("ssuserver") {
		s.pages[5].PopulateChild(name(i), i)
	}
}

func (s *SAMWebConfig) Serve() {
	s.populate()
	for _, i := range s.pages {
		s.localService.HandleFunc(i.URL(), i.Say)
		s.localService.HandleFunc(i.APIURL(), i.SayAPI)
	}
	if err := http.ListenAndServe(s.host+""+s.port, s.localService); err != nil {
		log.Fatal(err)
	}
}

func (s *SAMWebConfig) render_header() string {
	r := "<!doctype html>\n"
	r += "<html lang=\"" + s.lang + "\">\n"
	r += "<head>\n"
	r += "  <meta charset=\"utf-8\">\n"
	r += "  <title>" + s.title + "</title>\n"
	r += "  <meta name=\"description\" content=\"" + s.title + "\">\n"
	r += "  <meta name=\"author\" content=\"eyedeekay\">\n"
	r += "  <link rel=\"stylesheet\" href=\"css/styles.css\">\n"
	r += "</head>\n"
	r += "<body>\n"
	return r
}

func (s *SAMWebConfig) render_footer() string {
	r := "  <script src=\"js/scripts.js\"></script>\n"
	r += "</body>\n"
	r += "</html>\n"
	return r
}

func NewSAMWebConfigFromOptions(opts ...func(*SAMWebConfig) error) (*SAMWebConfig, error) {
	var s SAMWebConfig
	s.host = "127.0.0.1"
	s.port = "7957"
	s.lang = "en"
	s.title = "SAMcatd Web Console"
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}

	s.pages = append(s.pages, &pagestring{dir: "./",
		url: "index", apiurl: "api/index", desc: "SAMcatd Control Panel",
		id: "control_panel", class: "", manager: s.manager,
	})
	s.pages = append(s.pages, &pagestring{dir: "./server/",
		url: "ntcp", apiurl: "api/ntcp", desc: "ntcp server tunnels",
		id: "ntcp_server", class: "server,ntcp", manager: s.manager,
	})
	s.pages = append(s.pages, &pagestring{dir: "./server/",
		url: "http", apiurl: "api/http", desc: "http/ntcp server tunnels",
		id: "http_server", class: "server,http", manager: s.manager,
	})
	s.pages = append(s.pages, &pagestring{dir: "./server/",
		url: "ssu", apiurl: "api/ssu", desc: "ssu server tunnels",
		id: "ssu_server", class: "server,ssu", manager: s.manager,
	})
	s.pages = append(s.pages, &pagestring{dir: "./client/",
		url: "ntcp", apiurl: "api/ntcp", desc: "ntcp client tunnels",
		id: "ntcp_client", class: "client,ntcp", manager: s.manager,
	})
	s.pages = append(s.pages, &pagestring{dir: "./client/",
		url: "ssu", apiurl: "api/ssu", desc: "ssu client tunnels",
		id: "ssu_client", class: "client,ssu", manager: s.manager,
	})

	s.localService = http.NewServeMux()
	return &s, nil
}

func Serve(s *sammanager.SAMManager) {
	if webinterface, webinterfaceerr := NewSAMWebConfigFromOptions(
        SetHost("127.0.0.1"),
        SetPort("7957"),
        SetManager(s),
    ); webinterfaceerr == nil {
		log.Println("Starting web interface")
		go webinterface.Serve()
	}
}
