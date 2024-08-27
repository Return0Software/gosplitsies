# SPDX-License-Identifier: AGPL-3.0-or-later
#
# SPDX-FileCopyrightText: 2024 Tristan Partin <tristan@partin.io>

alias b := build
alias d := debug

version := `git describe --always --long --dirty || cat VERSION`

build:
	go build -o gsplit -ldflags \
		"-X github.com/Return0Software/gosplitsies/cmd.debug=false \
		 -X github.com/Return0Software/gosplitsies/cmd.version={{version}}"

debug:
	go build -o gsplit.debug -ldflags \
		"-X github.com/Return0Software/gosplitsies/cmd.debug=true \
		 -X github.com/Return0Software/gosplitsies/cmd.version={{version}}"
	@echo 'Attach to the server with `dlv attach $(pidof gsplit.debug)`'
