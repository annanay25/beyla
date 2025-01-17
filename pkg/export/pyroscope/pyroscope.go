package pyroscope

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/mariomac/pipes/pipe"

	"github.com/grafana/beyla/pkg/internal/request"
)

type PyroscopeExport string

const (
	PyroscopeExportDisabled = PyroscopeExport("disabled")
	PyroscopeExportText     = PyroscopeExport("text")
)

func mlog() *slog.Logger {
	return slog.With("component", "debug.TracePrinter")
}

func (t PyroscopeExport) Valid() bool {
	switch t {
	case PyroscopeExportText:
		return true
	}

	return false
}

func (t PyroscopeExport) Enabled() bool {
	return t.Valid() && t != PyroscopeExportDisabled
}

type PyroscopeExporter struct {
	Endpoint string
}

func resolvePyroscopeExportFunc(p PyroscopeExport, pyroscopeEndpoint string) pipe.FinalFunc[[]request.Span] {
	e := &PyroscopeExporter{Endpoint: pyroscopeEndpoint}

	switch p {
	case PyroscopeExportText:
		return e.textExporter
	}

	return pipe.IgnoreFinal[[]request.Span]()
}

func ProfileNode(p PyroscopeExport, pyroscopeEndpoint string) pipe.FinalProvider[[]request.Span] {
	exportFunc := resolvePyroscopeExportFunc(p, pyroscopeEndpoint)

	return func() (pipe.FinalFunc[[]request.Span], error) {
		return exportFunc, nil
	}
}

func (e *PyroscopeExporter) textExporter(input <-chan []request.Span) {
	finalTextDoc := ""
	for spans := range input {
		for _, span := range spans {
			// Dummy value 100 for now
			finalTextDoc += fmt.Sprintf("%v 100\n", span)
		}
	}

	// HTTP Post request to send this doc to the Pyroscope endpoint in Grafana Cloud.
	timeNow := time.Now().Unix()
	endpoint := e.Endpoint + "/ingest?name=beyla-test-app&from=" + string(timeNow-1) + "&until=" + string(timeNow)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(finalTextDoc)))
	if err != nil {
		mlog().Error("failed to create request", "error", err)
		return
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		mlog().Error("failed to send request", "error", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		mlog().Error("received non-OK response", "status", resp.StatusCode)
		return
	}

	mlog().Info("successfully sent document to Pyroscope endpoint")
}
