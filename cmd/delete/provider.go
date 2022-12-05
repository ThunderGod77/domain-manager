/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"
	"github.com/ThunderGod77/domain-manager/database"
	"github.com/spf13/cobra"
)

// providerCmd represents the provider command
var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if provider == "" {
			fmt.Println("please provide a valid registrar")
			return
		}
		err := database.DeleteProvider(provider)
		if err != nil {
			fmt.Println("could not deleted provider", err)
			return
		}
		fmt.Println("successfully deleted the provider")
	},
}

var provider string

func init() {
	deleteCmd.AddCommand(providerCmd)
	recordCmd.Flags().StringVarP(&provider, "provider", "p", "", "provider/domain registrar to be deleted")
	recordCmd.MarkFlagRequired("provider")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// providerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// providerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
