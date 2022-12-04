package new

import (
	"github.com/spf13/cobra"
)

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

// domainCmd represents the domain command
var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "Adds a provider",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("running provider command")

	},
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
