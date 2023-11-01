package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	beurse "github.com/erfahtech/be_erfahtech/module"
)

func init() {
	functions.HTTP("urse-signup", urseSignup)
}

func urseSignup(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://erfahtech.github.io")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://erfahtech.github.io")
	//insertUser
	fmt.Fprintf(w, beurse.GCFHandlerSignup("MONGOSTRING", "db_urse", "user", r))
}