/*
Copyright © 2022 Andi Powers-Holmes <aholmes@omnom.net>
*/

// Package cmd handle the cli commands
package cmd

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "requisition [url|command]",
	Short: "requisition me that file",
	Long:  `a multithreaded http downloader`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"download", "version"}, cobra.ShellCompDirectiveNoFileComp
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/requisition/requisition.yaml)")
	rootCmd.Args = cobra.MinimumNArgs(1)

	viper.BindPFlag("download.threads", rootCmd.PersistentFlags().Lookup("threads"))
	viper.BindPFlag("download.out", rootCmd.PersistentFlags().Lookup("out"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		conf_dir := path.Join(home, ".config", "requisition")
		if _, err := os.Stat(conf_dir); os.IsNotExist(err) {
			os.Mkdir(conf_dir, 0755)
		}
		cobra.CheckErr(err)

		// Search config in home directory with name ".requisition" (without extension) or current directory
		viper.SetConfigType("yaml")
		viper.SetConfigName("requisition")
		viper.AddConfigPath(conf_dir)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
