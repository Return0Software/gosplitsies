/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * SPDX-FileCopyrightText: 2024 Tristan Partin <tristan@partin.io>
 */

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var debug string
var configFile string
var zapConfig zap.Config

var rootCmd = &cobra.Command{
	Use:   "gsplit",
	Short: "GoSplitsies is a pay splitting application",
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serverCmd)

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Path to a config file")
	rootCmd.PersistentFlags().StringP("log-filepath", "l", "", "Path to log to if log-location is file")
	rootCmd.PersistentFlags().String("log-location", "stderr", "Location to send logs to")

	serverCmd.Flags().Uint16P("port", "p", 5431, "Port to start the server on")

	flags := map[string]*pflag.Flag{
		"log.filepath": rootCmd.PersistentFlags().Lookup("log-filepath"),
		"log.location": rootCmd.PersistentFlags().Lookup("log-location"),

		"server.port": serverCmd.Flags().Lookup("port"),
	}
	for key, flag := range flags {
		if err := viper.BindPFlag(key, flag); err != nil {
			fmt.Fprintf(os.Stderr, "failed to bind flags for configuration purposes: %s\n", err)
			os.Exit(1)
		}
	}
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("gosplitsies")
		viper.SetConfigType("yaml")

		// TODO: Windows and macOS?
		viper.AddConfigPath("/etc")
	}

	viper.SetEnvPrefix("gsplit")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok && configFile != "" {
			fmt.Fprintf(os.Stderr, "failed to read in config: %s\n", err)
			os.Exit(1)
		}
	}

	if debug == "true" {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.Level.SetLevel(zapcore.DebugLevel)
	} else {
		zapConfig = zap.NewProductionConfig()
		zapConfig.Level.SetLevel(zapcore.InfoLevel)
	}

	logLocation := viper.GetString("log.location")
	if logLocation == "file" {
		logFilepath := viper.GetString("log.filepath")
		if logFilepath == "" {
			fmt.Fprintln(os.Stderr, "no log filepath provided")
			os.Exit(1)
		}

		zapConfig.OutputPaths = []string{logFilepath}
	} else if logLocation == "stdout" {
		zapConfig.OutputPaths = []string{"stdout"}
	} else if logLocation == "stderr" {
		// Nothing to do since this is the default
	} else if logLocation == "syslog" { //nolint:staticcheck
		// TODO: Fail this if macOS or Windows
		// TODO: Also, actually implement this
	}

	logger, err := zapConfig.Build()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to initialize logger")
	}
	//nolint:errcheck
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		zap.S().Error("%s", err)
		os.Exit(1)
	}
}
