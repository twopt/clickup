package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/twopt/clickup/client"
	"github.com/twopt/clickup/internal"
)

var listsCmd = &cobra.Command{
	Use:   "lists FOLDERID",
	Short: "get data for lists in a folder",
	Long:  `Request JSON data for all lists by folder ID`,
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if authed := internal.CheckTokenExists(); !authed {
			internal.SaveToken(internal.GetToken())
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("archived", cmd.Flags().Lookup("archived"))
		l := client.ListRequest{
			FolderID: strings.Trim(args[0], " "),
			Archived: viper.GetBool("archived"),
		}
		client.Request(l)
	},
}

func init() {
	getCmd.AddCommand(listsCmd)
	listsCmd.Flags().BoolP("archived", "a", false, "include archived lists in output")
}
