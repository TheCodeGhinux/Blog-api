package cmd

import (
	"fmt"
	"go-blog/pkg/config"
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
	// configs := configSet()
	config.Set()
	configs := config.Get()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app name": viper.Get("App.name"),
		})
	})
	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
