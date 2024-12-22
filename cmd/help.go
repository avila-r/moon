package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of Moon",
		Long:  `All software has versions. This is Moon's`,
		Run: func(cmd *cobra.Command, args []string) {
			// ...
		},
	})
}
