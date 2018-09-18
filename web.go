// +build webface

package samcatweb

import "log"

import "github.com/eyedeekay/sam-forwarder/manager"

func (s *SAMWebConfig) Serve() {

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
	return &s, nil
}

func Serve(s *sammanager.SAMManager) {
	if webinterface, webinterfaceerr = samcatweb.NewSAMWebConfigFromOptions(); webinterfaceerr == nil {
		s.manager = s
		log.Println("Starting web interface")
		go webinterface.Serve()
	}
}
