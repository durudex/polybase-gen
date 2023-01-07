/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	configPath string

	root = &cobra.Command{
		Use:   "polybase-gen",
		Short: "CLI for generating Polybase Collections API code.",
		Long:  "polybase-gen is a CLI that generates the API code of the collections you specify.",
		Run: func(cmd *cobra.Command, args []string) {
			if configPath == "" {
				log.Fatal().Msg("Config file not specified!")
			}
		},
	}
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	root.Flags().StringVarP(&configPath, "config", "c", "", "path to config file")
}

func main() {
	if err := root.Execute(); err != nil {
		log.Fatal().Err(err).Msg("error executing cli")
	}
}
