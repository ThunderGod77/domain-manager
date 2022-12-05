//Copyright © 2022 NAME HERE <EMAIL ADDRESS>

// */
package get

import (
	"fmt"
	"github.com/ThunderGod77/domain-manager/database"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"os"
)

// domainsCmd represents the domains command
var domainsCmd = &cobra.Command{
	Use:   "domains",
	Short: "Get all the domains registered with the cli.",
	Long:  `Get all the domains registered with the cli application - domain name,registrar and description`,
	Run: func(cmd *cobra.Command, args []string) {
		domains, err := database.GetDomains()
		if err != nil {
			fmt.Println("could not get the domains", err)
			return
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Domain", "Provider", "Description"})
		i := 1
		for _, val := range domains {
			t.AppendRow([]interface{}{i, val.DomainName, val.Provider, val.Description})
			i++
		}
		t.Render()
	},
}

func init() {
	getCmd.AddCommand(domainsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
