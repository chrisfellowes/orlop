// Copyright (c) 2020 Ketch Kloud, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package orlop

import (
	"github.com/felixge/httpsnoop"
	"go.ketch.com/lib/orlop/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	"net/http"
)

// Metrics is middleware for handling metrics
func Metrics(next http.Handler) http.Handler {
	inflightRequests, err := metrics.NewInt64UpDownCounter("requests.in.flight", metric.WithUnit("requests"))
	if err != nil {
		log.Fatal(err)
	}

	requestDuration, err := metrics.NewFloat64ValueRecorder("request.duration.seconds", metric.WithUnit("s"))
	if err != nil {
		log.Fatal(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := attribute.String("method", r.Method)
		route := attribute.String("route", r.URL.Path)

		inflightRequests.Add(r.Context(), 1, method, route)
		defer inflightRequests.Add(r.Context(), -1, method, route)

		m := httpsnoop.CaptureMetrics(next, w, r)
		requestDuration.Record(r.Context(), m.Duration.Seconds(), method, route, attribute.Int("status_code", m.Code))
	})
}

var exporter *prometheus.Exporter

// MetricsHandler is the Prometheus metrics exporter
type MetricsHandler struct{}

// NewMetricsHandler creates a new MetricsHandler
func NewMetricsHandler() http.Handler {
	return &MetricsHandler{}
}

func (s *MetricsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if exporter != nil {
		exporter.ServeHTTP(w, r)
	} else {
		log.Error("metrics exporter not defined")
	}
}
