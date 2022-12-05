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
	Short: "Deletes a domain",
	Long:  `This command is used to delete domain locally stored with your cli application.Required flags 'domain'`,
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
