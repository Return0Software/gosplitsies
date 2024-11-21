/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * SPDX-FileCopyrightText: 2024 Tristan Partin <tristan@partin.io>
 */

package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Return0Software/gosplitsies/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var serverPort uint16

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()

		mux.HandleFunc("/logs/level", zapConfig.Level.ServeHTTP)

		port := viper.GetUint16("server.port")
		zap.S().Infof("Starting server on port %d", port)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), middleware.NewLogger(mux)); err != nil {
			zap.S().Fatalf("failed to start the server on port %d: %s", serverPort, err)
			os.Exit(1)
		}
	},
}
