<!--
SPDX-License-Identifier: AGPL-3.0-or-later

SPDX-FileCopyrightText: 2024 Tristan Partin <tristan@partin.io>
-->

# Server Configuration

Refer to the [example] for a list of all available settings.

[example]: ./gosplitsies.example.yaml

## Logging

If you want to log to a file, set `log.filepath` to a non-empty string, and set
`log.location` to `file`.

## Change Configuration via Files

Configuration files can be passed on the command line via the root
`-c`/`--config` option.

## Change Configuration via Environment Variables

All environment variables are prefixed with `GSPLIT_`. For example, if you want
to set the setting, `server.port`, via an environment variable, you can set the
`GSPLIT_SERVER_PORT` environment variable. `xxx-yyy` would translate to
`GSPLIT_XXX_YYY`.

## Change Configuration via Command Line Options

Please refer to the `-h`/`--help` documentation of the relevant commands.
Command line options take precedence over environment variables and config
files.

## Change Configuration via HTTP Requests

### Logging

The logging level defaults to `DEBUG` in debug mode and `INFO` in production
mode. It can be changed at [runtime].

[runtime]: https://pkg.go.dev/go.uber.org/zap#AtomicLevel.ServeHTTP
