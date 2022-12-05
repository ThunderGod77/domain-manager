/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"
	"github.com/ThunderGod77/domain-manager/database"
	"github.com/spf13/cobra"
)

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := database.DeleteDomain(domainName)
		if err != nil {
			fmt.Println("could not delete domain", err)
			return
		}
		fmt.Println("successfully deleted the domain")
	},
}
var domainName string

func init() {
	deleteCmd.AddCommand(domainCmd)
	domainCmd.Flags().StringVarP(&domainName, "domain", "d", "", "domain name to be deleted")
	recordCmd.MarkFlagRequired("domain")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
