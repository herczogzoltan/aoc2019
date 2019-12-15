package cmd

import (
	"fmt"
	"os"

	"github.com/herczogzoltan/aoc2019/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Day will be the base struct for the ongoing days.
type Day struct {
	filePath string
}

var (
	// InputSource is the filepath to read input values from.
	InputSource string
	// Part of the daily assignment. Either 1 or 2.
	Part int

	rootCmd = &cobra.Command{
		Use:   "aoc2019",
		Short: "Advent of Code 2019 ",
		Long: `Advent of Code 2019 solutions written in
	Go-lang as a practice and the joy of my own.`,
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of AoC 2019",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Build Date:", version.BuildDate)
			fmt.Println("Git Commit:", version.GitCommit)
			fmt.Println("Version:", version.Version)
			fmt.Println("Go Version:", version.GoVersion)
			fmt.Println("OS / Arch:", version.OsArch)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&InputSource, "source", "s", "", "File path which contains the input for the assignment")
	rootCmd.PersistentFlags().IntVarP(&Part, "part", "p", 1, "Which part do you want to run for the specific day?")
	rootCmd.MarkFlagRequired("source")
	rootCmd.MarkFlagRequired("part")
	rootCmd.AddCommand(versionCmd)

}

// Execute the root command.
func Execute() error {
	return rootCmd.Execute()
}

// LoadSourceFile loads a file to process by another command.
func LoadSourceFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("An error occured when tried to open the file")
	}
	return file
}

// Check checks if error is not nil, and then logs it.
func Check(e error) {
	if e != nil {
		log.Error(e)
	}
}
