package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/avila-r/moon/config"

	_ "github.com/manifoldco/promptui" // Blank import to avoid go mod tidy
)

var root = &cobra.Command{
	Use:   config.Get().Info.Command,
	Short: config.Get().Info.ShortDescription,
	Long:  config.Get().Info.LongDescription,
	Run: func(cmd *cobra.Command, args []string) {
		// ...
	},
}

func Execute() {
	if err := root.Execute(); err != nil {
		log.Fatalf("error occurred: %v", err)
	}
}
