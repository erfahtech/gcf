package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	beurse "github.com/erfahtech/be_erfahtech/module"
)

func init() {
	functions.HTTP("urse-getprofile", UrseGetProfile)
}

func UrseGetProfile(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://ursmartecosystem.my.id")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://ursmartecosystem.my.id")
	fmt.Fprintf(w, beurse.GCFGetUserByEmail("PASETOPUBLICKEY", "MONGOSTRING", "db_urse", "user", r))

}