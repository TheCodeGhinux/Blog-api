package cmd

import (
	"go-blog/pkg/config"
	"go-blog/pkg/routing"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app",
	Long:  `Serve app on the server on host and port in config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		Serve()
	},
}

func Serve() {
	config.Set()

	routing.Init()

	routing.RegisterRoutes()
	routing.Serve()
}
