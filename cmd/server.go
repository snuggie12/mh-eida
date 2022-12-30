/*
Copyright © 2022 snuggie12

*/
package cmd

import (
	cmdutil "snuggie12/eida/cmd/util"
	srv "snuggie12/eida/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the eida server",
	Long:  "",
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serverCmd)

	viper.SetDefault("logging", map[string]string{
		"level": "info",
	})

	serverCmd.PersistentFlags().Bool("admin-listen-local", false, "Only listen on localhost")
	serverCmd.PersistentFlags().StringP("admin-port", "p", "8712", "port for admin endpoint")
	serverCmd.PersistentFlags().String("log-level", "info", "Output level of logs (debug, info, warn, error, dpanic, panic, fatal)")

	viper.BindPFlag("admin.listenLocal", serverCmd.PersistentFlags().Lookup("admin-listen-local"))
	viper.BindPFlag("admin.port", serverCmd.PersistentFlags().Lookup("admin-port"))
	viper.BindPFlag("logging.level", serverCmd.PersistentFlags().Lookup("log-level"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func serve(cmd *cobra.Command, args []string) {
	logger := cmdutil.Logger(viper.GetString("logging.level"))
	logger.Info("Starting Admin Server")

	var config CmdConfig

	err := viper.Unmarshal(&config)
	if err != nil {
		logger.Fatalw("Unable to decode into config struct", "error", err)
	}
	logger.Debugf("Starting with config: %v", config)

	server := srv.NewServer(&config.Config, logger)
	server.StartAdminServer()
}
