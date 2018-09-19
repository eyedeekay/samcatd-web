// +build webface

package samcatweb

import (
	"log"
	"net/http"
	"strings"
)

func dedouble(s, t, u string) string {
	v := s
	for {
		if !strings.Contains(v, t) {
			if len(v) > 1 {
				return strings.TrimSuffix(v, u)
			} else {
				return v
			}
		} else {
			v = strings.Replace(v, t, u, -1)
		}
	}
}

func stringify(s *[]string) string {
	var p string
	for _, x := range *s {
		p += x + ","
	}
	r := dedouble(p, ",,", ",")
	return r
}

func name(s string) string {
	for _, r := range strings.Split(s, "\n") {
		if strings.Contains(r, "name=") {
			name := strings.SplitN(r, "name=", 2)
			if len(name) == 2 {
				trimmedname := strings.Trim(name[1], " ")
				returnedname := strings.Trim(trimmedname, "\n")
				return returnedname
			}
		}
	}
	return "NULL"
}

func condemit(pr, s string) string {
	if s != "" {
		return pr + s
	}
	return ""
}

func makeclass(s, p string) string {
	replacedslashes := strings.Replace(p+s, "/", ",", -1)
	replacedunderscores := strings.Replace(replacedslashes, "_", ",", -1)
	return dedouble(replacedunderscores, ",,", ",")
}

func makeid(s, p string) string {
	replacedslashes := strings.Replace(p+s, "/", "_", -1)
	replacedcommas := strings.Replace(replacedslashes, ",", "_", -1)
	return dedouble(replacedcommas, "__", "_")
}

func makeurl(s, p string) string {
	replacedcommas := strings.Replace(p+s, ",", "/", -1)
	replacedunderscores := strings.Replace(replacedcommas, "_", "/", -1)
	return dedouble(replacedunderscores, "//", "/")
}

func (s *pagestring) render_header() string {
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

func (s *pagestring) render_footer() string {
	r := "  <script src=\"js/scripts.js\"></script>\n"
	r += "</body>\n"
	r += "</html>\n"
	r += ""
	return r
}

func (p *pagestring) PopulateChild(s, value string) {
	p.children = append(p.children, &pagestring{dir: p.dir, title: p.title, lang: p.lang,
		url: makeurl(s, p.url), apiurl: makeurl(s, p.apiurl), desc: p.desc + " : " + s + ":" + value,
		id: makeid(s, p.id), class: makeclass(s, p.id), manager: p.manager,
	})
}

func (p *pagestring) URL() string {
	return p.dir + "/" + p.url
}

func (p *pagestring) APIURL() string {
	return p.dir + "/" + p.apiurl
}

func (p *pagestring) render_div(s string) string {
	query := p.class + "," + s
	var r string
	for _, val := range *p.manager.List(query) {
		r += "<div "
		r += "class=\"" + p.class + "\" "
		r += "id=\"" + p.id + condemit("_", s) + "\" >"
		r += val
		r += "</div>"
	}
	return r
}

func (p *pagestring) render_apiurl(s string) string {
	query := p.class + "," + s
	r := stringify(p.manager.List(query)) + "\n"
	return r
}

func (p *pagestring) Say(w http.ResponseWriter, r *http.Request) {
	query := strings.Replace(strings.TrimPrefix(r.URL.Path, p.URL()), "/", ",", -1)
	message := p.render_header()
	message += p.render_div(query)
	message += p.render_footer()
	log.Println("Responnding to the page request", r.URL.Path)
	w.Write([]byte(message))
}

func (p *pagestring) SayAPI(w http.ResponseWriter, r *http.Request) {
	query := strings.Replace(strings.TrimPrefix(r.URL.Path, p.APIURL()), "/", ",", -1)
	message := p.render_apiurl(query)
	log.Println("Responnding to the API request", r.URL.Path)
	w.Write([]byte(message))
}
