/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * SPDX-FileCopyrightText: 2024 Tristan Partin <tristan@partin.io>
 */

package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Logger struct {
	handler http.Handler
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	zap.S().Infof("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}
