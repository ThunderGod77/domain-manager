package new

import (
	"fmt"
	"github.com/ThunderGod77/domain-manager/cmd/prompt"
	"github.com/ThunderGod77/domain-manager/database"
	"github.com/spf13/cobra"
)

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "adds a domain",
	Long:  `Adds a new domain,the domain is stored locally with the cli application.Requires use to enter the domain name,provider(currently only supports go daddy) and description.All the data is stored locally in plaintext format.`,
	Run: func(cmd *cobra.Command, args []string) {

		domainPrompt := prompt.CreatePrompt("Domain Name", prompt.Validate)
		domainResult, err := domainPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		providerPrompt := prompt.CreateSelectPrompt("select a domain name registrar", database.Providers)
		_, ProviderResult, err := providerPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		descriptionPrompt := prompt.CreatePrompt("Description ", prompt.Validate)
		descriptionResult, err := descriptionPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		err = database.AddDomain(domainResult, ProviderResult, descriptionResult)
		if err != nil {
			fmt.Printf("could not add domain %v\n", err)
			return
		}
		fmt.Printf("added the domain %s successfully\n", domainResult)

	},
}

func init() {
	newCmd.AddCommand(domainCmd)
	//domainCmd.Flags().StringVarP(&domainName, "domain name to be added", "d", "", "domain name you want to add to the programme")
	//domainCmd.Flags().StringVarP(&provider)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
