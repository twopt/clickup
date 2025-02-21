package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/twopt/clickup/utils"
)

// add new config options here with the flag command and help
// message for the setting. only string flags allowed
var configOptions = map[string]string{
	"team":  "set the Team ID",
	"token": "set the Auth Token manually",
}

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "sets config options",
	Long: `set is used to configure extended options and save
	them to the config file`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		for flag := range configOptions {
			value, _ := cmd.Flags().GetString(flag)
			if value != "" {
				viper.BindPFlag(flag, cmd.Flags().Lookup(flag))
				err := viper.WriteConfigAs(utils.GetConfigFile())
				if err != nil {
					log.Fatalf("cannot write config to %s: %v", utils.GetConfigFile(), err)
				}
				fmt.Printf("Saved %s to %s\n", flag, utils.GetConfigFile())
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	for flag, description := range configOptions {
		setCmd.Flags().StringP(flag, "", "", description)
	}
}
