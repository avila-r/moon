package cmd

import (
	"github.com/spf13/cobra"

	"github.com/avila-r/moon/config"
)

func init() {
	config := config.Get()

	root.PersistentFlags().String("license", config.Info.License, "License information")
	root.PersistentFlags().String("author", config.Info.Author, "Author of the application")
	root.PersistentFlags().String("version", config.Info.Version, "Application version")
	root.PersistentFlags().String("repository", config.Info.Repository, "Repository URL")
	root.PersistentFlags().Bool("debug", config.Advanced.Debug, "Enable debug mode")

	root.AddCommand(&cobra.Command{
		Use:   "config",
		Short: "Print configuration details",
		Long:  "Displays the current configuration details loaded from the moon.toml file.",
		Run: func(cmd *cobra.Command, args []string) {
			config.Log()
		},
	})
}
