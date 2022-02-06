package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"rssmq/pkg"
)

var (
	rootCmd = &cobra.Command{
		Use:   "rssmq",
		Short: "RSSMQ is a simple RSS feed aggregator",
		Long:  `RSSMQ is a microservice that aggregates RSS feeds and sends them to a backend of your choosing.`,
		Run: func(cmd *cobra.Command, args []string) {
			pkg.Run()
		},
	}
	cfgFile string
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /opt/.rssmq)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("json")
		viper.SetConfigName("rssmq")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
		fmt.Println("Unable to read config file")
	}
}
