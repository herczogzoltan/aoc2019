package cmd

import (
	"github.com/spf13/cobra"
)

type day3 Day

var (
	day3Command = &cobra.Command{
		Use: "day3",
		Run: func(cmd *cobra.Command, args []string) {
			// day3 := day3{
			// 	filePath: InputSource,
			// }
			switch Part {
			case 1:
				//fmt.Println("First part result: ", day3.firstPart())
			case 2:
				//fmt.Printf("Second part result: %v", day3.secondPart())
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(day3Command)
}

func (d day3) firstPart() {
}

func (d day3) secondPart() {

}
