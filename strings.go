// +build webface

package samcatweb

import (
	"fmt"
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
		return strings.Trim(pr+s, " ")
	}
	return ""
}

func makeclass(s, p string) string {
	replacedslashes := strings.Replace(p+","+s, "/", " ", -1)
	replacedcommas := strings.Replace(replacedslashes, ",", " ", -1)
	replacedunderscores := strings.Replace(replacedcommas, "_", " ", -1)
	return strings.TrimPrefix(dedouble(replacedunderscores, "  ", " "), " ")
}

func makeid(s, p string) string {
	replacedslashes := strings.Replace(p+"_"+s, "/", "_", -1)
	replacedcommas := strings.Replace(replacedslashes, ",", "_", -1)
	replacedperiods := strings.Replace(replacedcommas, ".", "_", -1)
	return strings.Replace(dedouble(replacedperiods, "__", "_"), " ", "", -1)
}

func makeurl(s, p string) string {
	replacedcommas := strings.Replace(p+"/"+s, ",", "/", -1)
	replacedunderscores := strings.Replace(replacedcommas, "_", "/", -1)
	return dedouble(replacedunderscores, "//", "/")
}

func render_header(title, lang, desc string) string {
	r := "<!doctype html>\n"
	r += "<html lang=\"" + lang + "\">\n"
	r += "  <head>\n"
	r += "    <meta charset=\"utf-8\">\n"
	r += "    <title>" + title + "</title>\n"
	r += "    <meta name=\"description\" content=\"" + desc + "\">\n"
	r += "    <meta name=\"author\" content=\"eyedeekay\">\n"
	r += "    <style>"
	r += defaultCSS()
	r += "    </style>"
	r += "    <link rel=\"stylesheet\" type=\"text/css\" href=\"/css/styles.css\" media=\"screen\" />\n"
	r += "  </head>\n"
	r += "  <body>\n"
	r += "\n"
	return r
}

func render_bar() string {
	r := "  <div id=\"toolbar\" class=\"toolbar\">\n"
	r += "    <a href=\"/index\" id=\"btn_index\" class=\"btn\"> Home </a>\n"
	r += "    <a href=\"/server/ntcp\" id=\"btn_ntcpserver\" class=\"btn\"> NTCP Server </a>\n"
	r += "    <a href=\"/server/http\" id=\"btn_httpserver\" class=\"btn\"> HTTP Server </a>\n"
	r += "    <a href=\"/server/ssu\" id=\"btn_ssuserver\" class=\"btn\"> SSU Server </a>\n"
	r += "    <a href=\"/client/ntcp\" id=\"btn_ntcpclient\" class=\"btn\"> NTCP Client </a>\n"
	r += "    <a href=\"/client/ssu\" id=\"btn_ssuclient\" class=\"btn\"> SSU Clients </a>\n"
	r += "  </div>\n"
	r += "<br>\n"
	r += "  <div id=\"showhidebar\" class=\"toolbar\">\n"
	/*r += "<a href=\"/server/ntcp\" id=\"btn_ntcpserver\" class=\"btn\"> NTCP Server </a>"*/
	r += "  </div>\n"
	r += "<br>\n"
	r += "\n"
	return r
}

func render_footer() string {
	r := "    <script src=\"/js/scripts.js\"></script>\n"
	r += "  </body>\n"
	r += "</html>\n"
	r += "\n"
	return r
}

func (p *pagestring) PopulateChild(s, value string) {
	p.children = append(p.children, &pagestring{dir: p.dir, title: p.title, lang: p.lang,
		url: makeurl(s, p.url), apiurl: makeurl(s, p.apiurl), desc: p.desc + " : " + s + ":" + value,
		id: makeid(s, p.id), class: makeclass(s, p.class), manager: p.manager,
	})
}

func (p *pagestring) URL() string {
	return "/" + strings.Replace(dedouble(p.dir+"/"+p.url, "//", "/"), "./", "", -1)
}

func (p *pagestring) APIURL() string {
	return "/" + strings.Replace(dedouble(p.dir+"/"+p.apiurl, "//", "/"), "./", "", -1)
}

func (p *pagestring) sub_div(val string) string {
	split := strings.Split(val, "\n")
	var r string
	for _, v := range split {
		log.Println("processing tunnel", v)
		if len(strings.SplitN(v, "=", 2)) == 2 {
			prefix := strings.SplitN(v, "=", 2)[0]
			log.Println("full category", prefix)
			name := strings.SplitN(v, "=", 2)[1]
			log.Println("item name", name)
			splitprefix := strings.Split(prefix, " ")
			log.Println("trimmed category", splitprefix)
			option := splitprefix[len(splitprefix)-1]
			log.Println("classifier", option)
			r += "    <div "
			r += "class=\"" + makeclass(option, p.class+",label") + "\" "
			r += "id=\"" + makeid(condemit("_", name), p.id+"_label") + "\" >"
			r += option + " : "
			r += "</div> "
			r += "    <div "
			r += "class=\"" + makeclass(option, p.class+",content") + "\" "
			r += "id=\"" + makeid(condemit("_", name), p.id) + "\" >"
			r += name
			r += "</div> \n"
		}
	}
	log.Println(r)
	return r
}

func (p *pagestring) render_div(s string) string {
	query := p.class + "," + s
	var r string
	for _, val := range *p.manager.List(query) {
		r += "  <div "
		r += "class=\"" + makeclass(s, p.class+",parent") + "\" "
		r += "id=\"" + makeid(condemit("_", s), p.id) + "\">\n"
		r += p.sub_div(val)
		r += "  </div>\n"
	}
	return r
}

func (p *pagestring) render_apiurl(s string) string {
	query := p.class + "," + s
	r := stringify(p.manager.List(query)) + ""
	return r
}

func (p *pagestring) Say(w http.ResponseWriter, r *http.Request) {
	query := dedouble(strings.Replace(strings.TrimPrefix(r.URL.Path, p.URL()), "/", ",", -1), ",,", ",")
	log.Println("Responding to the page request", r.URL.Path)
	fmt.Fprintln(w, render_header(p.title, p.lang, p.desc))
	log.Println("header sent")
	fmt.Fprintln(w, render_bar())
	log.Println("toolbar sent")
	fmt.Fprintln(w, p.render_div(query))
	log.Println("content sent")
	fmt.Fprintln(w, render_footer())
	log.Println("footer sent")
	return
}

func (p *pagestring) SayAPI(w http.ResponseWriter, r *http.Request) {
	query := dedouble(strings.Replace(strings.TrimPrefix(r.URL.Path, p.APIURL()), "/", ",", -1), ",,", ",")
	log.Println("Responding to the API request", r.URL.Path, p.render_apiurl(query))
	fmt.Fprintln(w, p.render_apiurl(query))
}
