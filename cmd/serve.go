package cmd

import (
	"go-blog/pkg/config"
	"go-blog/pkg/routing"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	r := routing.GetRouter()

	// r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app name": viper.Get("App.name"),
		})
	})

	routing.Serve()
}
