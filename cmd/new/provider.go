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
var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "Adds a provider",
	Long:  `Stores the credentials/api keys of your domain name registrar.Stores the credentials locally in plain text form at along with the cli application`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("running provider command")
		providerPrompt := prompt.CreateSelectPrompt("select a domain name registrar", database.Providers)
		_, providerResult, err := providerPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		providers, err := database.GetProviders()
		if err != nil {
			fmt.Println("some error occurred", err)
			return
		}
		providerData, exists := providers[providerResult]

		akInput := "please enter the access key"
		skInput := "please enter the secret"
		validator := prompt.Validate
		if exists {
			akInput = akInput + fmt.Sprintf(" [%s]", maskLeft(providerData.AccessKey))
			skInput = skInput + fmt.Sprintf(" [%s]", maskLeft(providerData.Secret))
			validator = prompt.IgnoreValidation
		}

		accessKeyPrompt := prompt.CreatePrompt(akInput, validator)
		akResult, err := accessKeyPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if akResult == "" {
			akResult = providerData.AccessKey
		}

		secretPrompt := prompt.CreatePrompt(skInput, prompt.IgnoreValidation)
		skResult, err := secretPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if skResult == "" {
			akResult = providerData.Secret
		}
		err = database.AddProvider(providerResult, akResult, skResult)
		if err != nil {
			fmt.Println("could not save credentials", err)
		}
		fmt.Println("Successfully stored credentials")
	},
}

func maskLeft(s string) string {
	rs := []rune(s)
	for i := 0; i < len(rs)-4; i++ {
		rs[i] = 'X'
	}
	return string(rs)
}
func init() {
	newCmd.AddCommand(providerCmd)
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
