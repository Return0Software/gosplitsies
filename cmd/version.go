/* SPDX-License-Identifier: AGPL-3.0-or-later
 *
 * SPDX-FileCopyrightText: 2024 Tristan Partin <tristan@partin.io>
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
