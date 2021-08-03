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
	"net/http"
	"strings"
)

// DefaultHTTPHeaders is middleware to handle default HTTP headers
func DefaultHTTPHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isGRPCRequest := r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc")

		if !isGRPCRequest {
			addCORSHeaders(w, r)
		}

		addSecurityHeaders(w, r)

		next.ServeHTTP(w, r)
	})
}

func addCORSHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Vary", "Origin")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
			w.Header().Set("Access-Control-Allow-Headers", headers)
			w.Header().Set("Access-Control-Allow-Methods", methods)
			return
		}
	}
}

func addSecurityHeaders(w http.ResponseWriter, r *http.Request) {
	if r.TLS != nil {
		// Only on TLS per https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
	}

	addHeaderIfNotExists(w, "X-Frame-Options", "deny")
	addHeaderIfNotExists(w, "X-Content-Type-Options", "nosniff")
	addHeaderIfNotExists(w, "Content-Security-Policy", "default-src 'self'")
	addHeaderIfNotExists(w, "X-XSS-Protection", "1; mode=block")
}

func addHeaderIfNotExists(w http.ResponseWriter, headerKey string, value string) {
	if len(w.Header().Get(headerKey)) == 0 {
		w.Header().Add(http.CanonicalHeaderKey(headerKey), value)
	}
}
