// +build webface

package samcatweb

import (
	"net/http"
	"strings"
)

func dedouble(s, t, u string) string {
	v := s
	for {
		if !strings.Contains(v, t) {
			return v
		} else {
			v = strings.Replace(v, t, u, -1)
		}
	}
}

func (p *pagestring) PopulateChild(s, value string) {
	slashed := dedouble(strings.Replace(strings.Replace(s, ",", "/", -1), "_", "/", -1), "//", "/")
	commaed := dedouble(strings.Replace(strings.Replace(s, "/", ",", -1), "_", ",", -1), ",,", ",")
	underscored := dedouble(strings.Replace(strings.Replace(s, ",", "_", -1), "/", "_", -1), "__", "_")
	p.children = append(p.children, &pagestring{dir: p.dir,
		url: p.url + "/" + slashed, apiurl: p.apiurl + "/" + slashed, desc: p.desc + " : " + s + ":" + value,
		id: p.id + "_" + underscored, class: p.class + "," + commaed, manager: p.manager,
	})
}

func (p *pagestring) URL() string {
	return p.dir + p.url
}
func (p *pagestring) APIURL() string {
	return p.dir + p.apiurl
}

func condemit(pr, s string) string {
	if s != "" {
		return pr + s
	}
	return ""
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
	message := p.render_div(query)
	w.Write([]byte(message))
}

func (p *pagestring) SayAPI(w http.ResponseWriter, r *http.Request) {
	query := strings.Replace(strings.TrimPrefix(r.URL.Path, p.URL()), "/", ",", -1)
	message := p.render_apiurl(query)
	w.Write([]byte(message))
}
