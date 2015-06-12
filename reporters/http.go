package reporters

import (
	"github.com/arussellsaw/telemetry"

	"encoding/json"
	"net/http"
)

//TelemetryHandler - http handler for telemetry
type TelemetryHandler struct {
	Tel *telemetry.Telemetry
}

func (t TelemetryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	output, err := json.MarshalIndent(t.Tel.GetAll(), "", " ")
	if err != nil {
		w.Write([]byte("failed to encode results"))
	}
	w.Write(output)
}
