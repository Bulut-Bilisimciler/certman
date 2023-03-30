package prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HandlePrometheusRequest() http.Handler {

	return promhttp.Handler()
}
