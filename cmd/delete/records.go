/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"fmt"
	"github.com/ThunderGod77/domain-manager/cmd/prompt"
	"github.com/ThunderGod77/domain-manager/database"
	"github.com/ThunderGod77/domain-manager/goDaddy"
	"log"

	"github.com/spf13/cobra"
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "deletes the dns record",
	Long:  `Deletes the dns record of a particular domain.Domain should already be registered along with it's domain registrar(provider) api keys.Required flags 'domain'`,
	Run: func(cmd *cobra.Command, args []string) {
		providerData, err := database.GetCredentials(domain)
		if err != nil {
			log.Println("failed to get credentials", err)
			return
		}
		if recordType == "" {
			rtPrompt := prompt.CreatePrompt("Record Type ", prompt.Validate)
			recordType, err = rtPrompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		}
		if recordName == "" {
			rnPrompt := prompt.CreatePrompt("Record Name ", prompt.Validate)
			recordName, err = rnPrompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		}

		err = goDaddy.DeleteRecord(providerData.AccessKey, providerData.Secret, recordType, recordName, domain)
		if err != nil {
			fmt.Println("could not delete record", err)
			return
		}
		fmt.Println("successfully deleted the dns record")
	},
}
var domain string
var recordType string
var recordName string

func init() {
	deleteCmd.AddCommand(recordCmd)

	recordCmd.Flags().StringVarP(&domain, "domain", "d", "", "domain name to look records for")
	recordCmd.MarkFlagRequired("domain")
	recordCmd.Flags().StringVarP(&recordType, "record type", "r", "", "record type")
	recordCmd.Flags().StringVarP(&recordName, "record name", "n", "", "record name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
