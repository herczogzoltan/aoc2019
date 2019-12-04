package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "aoc2019",
		Short: "Advent of Code 2019 ",
		Long: `Advent of Code 2019 solutions written in
	Go-lang as a practice and the joy of my own.`,
	}
)

// Execute the root command.
func Execute() error {
	return rootCmd.Execute()
}
