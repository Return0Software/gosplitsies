/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * SPDX-FileCopyrightText: 2024 Tristan Partin <tristan@partin.io>
 */

package cmd

import (
	"fmt"
	"net"
	"net/http"

	"github.com/Return0Software/gosplitsies/logger"
	"github.com/Return0Software/gosplitsies/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tristan957/go-sd-notify"
	"go.uber.org/zap"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		adminMux := http.NewServeMux()
		adminMux.HandleFunc("/logs/level", logger.HTTPEndpoint())

		rootMux := http.NewServeMux()
		rootMux.Handle("/admin", adminMux)

		port := viper.GetUint16("server.port")
		zap.S().Infof("Starting server on port %d", port)

		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			zap.S().Fatalf("failed to begin listening on port %d: %s", port, err)
		}

		_ = notify.Ready()

		if err := http.Serve(listener, middleware.NewLogger(rootMux)); err != nil {
			_ = notify.Stopping()
			zap.S().Fatalf("failed to start the server: %s", err)
		}
	},
}
