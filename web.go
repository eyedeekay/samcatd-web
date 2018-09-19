// +build webface

package samcatweb

import (
	"log"
	"net/http"
	"strings"
)

import "github.com/eyedeekay/sam-forwarder/manager"

func (s *SAMWebConfig) populate() {
	for _, i := range *s.manager.List("") {
		log.Println("Registering control page", name(i))
		s.pages[0].PopulateChild(name(i), i)
	}
	for _, i := range *s.manager.List("ntcpserver") {
		log.Println("Registering control page", name(i))
		s.pages[1].PopulateChild(name(i), i)
	}
	for _, i := range *s.manager.List("httpserver") {
		log.Println("Registering control page", name(i))
		s.pages[2].PopulateChild(name(i), i)
	}
	for _, i := range *s.manager.List("ssuserver") {
		log.Println("Registering control page", name(i))
		s.pages[3].PopulateChild(name(i), i)
	}
	for _, i := range *s.manager.List("nctpclient") {
		log.Println("Registering control page", name(i))
		s.pages[4].PopulateChild(name(i), i)
	}
	for _, i := range *s.manager.List("ssuserver") {
		log.Println("Registering control page", name(i))
		s.pages[5].PopulateChild(name(i), i)
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
	r += ""
	return r
}

func (s *SAMWebConfig) render_footer() string {
	r := "  <script src=\"js/scripts.js\"></script>\n"
	r += "</body>\n"
	r += "</html>\n"
	r += ""
	return r
}

func (p *SAMWebConfig) render_div(s string) string {
	query := s
	var r string
	for _, val := range *p.manager.List(query) {
		r += "<div "
		r += "class=\"" + s + "\" "
		r += "id=\"" + condemit("_", s) + "\" >"
		r += val
		r += "</div>"
	}
	return r
}

func (p *SAMWebConfig) render_apiurl(s string) string {
	query := s
	r := stringify(p.manager.List(query)) + "\n"
	return r
}

func (s SAMWebConfig) Say(w http.ResponseWriter, r *http.Request) {
	message := s.render_header()
	message += r.URL.Path
	message += s.render_footer()
	log.Println("Responnding to the page request", r.URL.Path)
	w.Write([]byte(message))
}

func (s SAMWebConfig) SayAPI(w http.ResponseWriter, r *http.Request) {
	query := strings.Replace(strings.TrimPrefix(r.URL.Path, "api/index.config"), "/", ",", -1)
	message := s.render_apiurl(query)
	log.Println("Responnding to the API request", r.URL.Path)
	w.Write([]byte(message))
}

func (s *SAMWebConfig) Serve() {
	s.populate()
    s.localService.HandleFunc("index.html", s.Say)
	log.Println("Registering control function for index.html")
	s.localService.HandleFunc("api/index.config", s.SayAPI)
	log.Println("Registering control function for index API")
	for _, i := range s.pages {
		log.Println("Registering control function", i.URL())
		s.localService.HandleFunc(i.URL(), i.Say)
		log.Println("Registering control API function", i.APIURL())
		s.localService.HandleFunc(i.APIURL(), i.SayAPI)
		for _, j := range i.children {
			log.Println("Registering control function", j.URL())
			s.localService.HandleFunc(j.URL(), j.Say)
			log.Println("Registering control API function", j.APIURL())
			s.localService.HandleFunc(j.APIURL(), j.SayAPI)
		}
	}
	if err := http.ListenAndServe(s.host+":"+s.port, s.localService); err != nil {
		log.Fatal(err)
	}
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
		lang: s.lang, title: s.title,
		url: "index", apiurl: "api/index", desc: "SAMcatd Control Panel",
		id: "control_panel", class: "", manager: s.manager,
	})
	s.pages = append(s.pages, &pagestring{dir: "./server/",
		lang: s.lang, title: s.title,
		url: "ntcp", apiurl: "api/ntcp", desc: "ntcp server tunnels",
		id: "ntcp_server", class: "server,ntcp", manager: s.manager,
	})
	s.pages = append(s.pages, &pagestring{dir: "./server/",
		lang: s.lang, title: s.title,
		url: "http", apiurl: "api/http", desc: "http/ntcp server tunnels",
		id: "http_server", class: "server,http", manager: s.manager,
	})
	s.pages = append(s.pages, &pagestring{dir: "./server/",
		lang: s.lang, title: s.title,
		url: "ssu", apiurl: "api/ssu", desc: "ssu server tunnels",
		id: "ssu_server", class: "server,ssu", manager: s.manager,
	})
	s.pages = append(s.pages, &pagestring{dir: "./client/",
		lang: s.lang, title: s.title,
		url: "ntcp", apiurl: "api/ntcp", desc: "ntcp client tunnels",
		id: "ntcp_client", class: "client,ntcp", manager: s.manager,
	})
	s.pages = append(s.pages, &pagestring{dir: "./client/",
		lang: s.lang, title: s.title,
		url: "ssu", apiurl: "api/ssu", desc: "ssu client tunnels",
		id: "ssu_client", class: "client,ssu", manager: s.manager,
	})

	s.localService = http.NewServeMux()
	return &s, nil
}

func Serve(s *sammanager.SAMManager, hp ...string) {
	var host, port string
	switch len(hp) {
	case 0:
		host = "127.0.0.1"
		port = "7957"
	case 1:
		host = "127.0.0.1"
		port = hp[0]
	case 2:
		host = hp[0]
		port = hp[1]
	default:
		host = "127.0.0.1"
		port = "7957"
	}
	if webinterface, webinterfaceerr := NewSAMWebConfigFromOptions(
		SetHost(host),
		SetPort(port),
		SetManager(s),
	); webinterfaceerr == nil {
		log.Println("Starting web interface")
		go webinterface.Serve()
	}
}
