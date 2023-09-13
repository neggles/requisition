/*
Copyright © 2023 Andi Powers-Holmes <aholmes@omnom.net>
*/
package cmd

import (
	"fmt"

	"github.com/neggles/requisition/internal/cmd/download"
	"github.com/spf13/cobra"
)

var downloadTask = &download.DownloadRequest{
	Url:     "",
	Out:     "",
	Threads: 1,

}

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download one or more files",
	Long: `Download one or more files to the current directory.

	Wow
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download called")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
