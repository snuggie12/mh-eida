package server

import (
	"encoding/json"
	"net/http"
)

// RegisterProfiler adds pprof endpoints to mux.
func (server *Server) addHealthToAdmin(mux *http.ServeMux) {
	// livenessProbe
	mux.HandleFunc(LivenessProbePath, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
}
