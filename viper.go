package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"log/slog"
)

func initViper() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Dynamically watch the config file for changes
	viper.WatchConfig()
	viper.OnConfigChange(onConfigChange)
}

func onConfigChange(e fsnotify.Event) {
	fmt.Println("Config file changed:", e.Name)

	adminMethod := viper.GetString("admin")
	slog.Info("Method to run", slog.String("@admin", adminMethod))
}
