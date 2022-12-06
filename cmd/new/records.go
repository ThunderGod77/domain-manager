/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package new

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
	Short: "add a new dns record to your domain",
	Long:  `Adds a new dns record to your domain name.Required flags - 'domain'.Optional flags = 'type'(record type of your dns record),'name'(record name of your dns record),'data'(value of the dns record - defaults to 3600),'ttl'(ttl of the dns record),'weight'(weight of the dns record - defaults to 1)`,
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
		if data == "" {
			dPrompt := prompt.CreatePrompt("Record Data ", prompt.Validate)
			data, err = dPrompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		}
		err = goDaddy.AddRecord(providerData.AccessKey, providerData.Secret, domain, data, recordName, recordType, weight, ttl, 0)
		if err != nil {
			fmt.Println("some error occurred", err)
			return
		}
		fmt.Println("successfully added a new dns record")
	},
}

var domain string
var recordType string
var recordName string
var ttl int
var weight float64
var data string

func init() {
	newCmd.AddCommand(recordCmd)

	recordCmd.Flags().StringVarP(&domain, "domain", "d", "", "domain name to look records for")
	recordCmd.MarkFlagRequired("domain")
	recordCmd.Flags().StringVarP(&recordType, "type", "t", "", "record type")
	recordCmd.Flags().StringVarP(&recordName, "name", "n", "", "record name")

	recordCmd.Flags().StringVar(&data, "data", "", "value of the dns record")

	recordCmd.Flags().IntVar(&ttl, "ttl", 3600, "ttl of dns record")
	recordCmd.Flags().Float64Var(&weight, "weight", 1, "weight of dns record")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
