package home

import (
	"fmt"
	"html/template"
	"net/http"

	"app/conf"
)

var home_tpl = template.Must(template.ParseFiles(fmt.Sprintf("%shome.html", conf.AppConfInfo.WWWDir)))

var HomePage = func(w http.ResponseWriter, r *http.Request) {
	/*方便调试，TODO delete*/
	home_tpl = template.Must(template.ParseFiles(fmt.Sprintf("%shome.html", conf.AppConfInfo.WWWDir)))
	var _ = home_tpl.Execute(w, nil)
	return
}
