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

	if len(t.Tel.GetAll()) == 0 {
		w.WriteHeader(204)
		return
	}

	output, err := json.MarshalIndent(t.Tel.GetAll(), "", " ")
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\": \"" + err.Error() + "\"}"))
	}

	w.Write(output)
}
