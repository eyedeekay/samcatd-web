// +build webface

package samcatweb

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func (s *SAMWebConfig) Serve() {
	s.populate()
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
	s.localService.HandleFunc("/js/scripts.js", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, s.jsstring)
		return
	})
	s.localService.HandleFunc("/css/styles.css", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, s.cssstring)
		return
	})
	s.localService.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, render_bar())
		fmt.Fprintln(w, "Dave's not here man.")
		return
	})
	log.Println("Starting web service")
	if err := http.ListenAndServe(s.host+":"+s.port, s.localService); err != nil {
		log.Fatal(err)
	}
}

func NewSAMWebConfigFromOptions(opts ...func(*SAMWebConfig) error) (*SAMWebConfig, error) {
	var s SAMWebConfig
	s.host = "127.0.0.1"
	s.port = "7957"
	s.lang = "en"
	s.jspath = ""
	s.csspath = ""
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
	if s.csspath != "" {
		if b, err := ioutil.ReadFile(s.csspath); err == nil {
			s.cssstring = string(b)
		} else {
			s.cssstring = defaultCSS()
		}
	} else {
		s.cssstring = defaultCSS()
	}
	if s.jspath != "" {
		if b, err := ioutil.ReadFile(s.jspath); err == nil {
			s.jsstring = string(b)
		} else {
			s.jsstring = defaultJS()
		}
	} else {
		s.jsstring = defaultJS()
	}
	s.localService = http.NewServeMux()
	return &s, nil
}

func Serve(s *sammanager.SAMManager, cssfile, jsfile string, hp ...string) {
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
		SetCSSPath(cssfile),
		SetJSPath(jsfile),
	); webinterfaceerr == nil {
		log.Println("Starting web interface")
		webinterface.Serve()
	}
}
