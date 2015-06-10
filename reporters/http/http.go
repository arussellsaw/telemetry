package httpReporter

import (
	"github.com/arussellsaw/telemetry"

	"encoding/json"
	"net/http"
)

type telemetryHandler struct {
	tel *telemetry.Telemetry
}

func (t telemetryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	output, err := json.MarshalIndent(t.tel.GetAll(), "", " ")
	if err != nil {
		w.Write([]byte("failed to encode results"))
	}
	w.Write(output)
}

func newHandler(tel *telemetry.Telemetry) http.Handler {
	handler := telemetryHandler{
		tel: tel,
	}
	return handler
}
