package get

import (
	"github.com/ThunderGod77/domain-manager/database"
	"github.com/ThunderGod77/domain-manager/goDaddy"
	"github.com/jedib0t/go-pretty/v6/table"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var domain string
var recordType string
var recordName string

// domainsCmd represents the domains command
var recordsCmd = &cobra.Command{
	Use:   "records",
	Short: "Get dns records of a domain",
	Long:  `Gets all the dns records of a domain,the domain should already be registered along with it's provider.Required flags - 'domain'.Options flags - 'type'(record type of a dns record),'name'(record name of a dns record).Returns the record name,record type,record value and ttl of the record`,
	Run: func(cmd *cobra.Command, args []string) {
		providerData, err := database.GetCredentials(domain)
		if err != nil {
			log.Println("failed to get credentials", err)
			return
		}
		records, err := goDaddy.GetRecords(providerData.AccessKey, providerData.Secret, recordType, recordName, domain)
		if err != nil || len(records) == 0 {
			log.Println("function get records failed", err)
			return
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Name", "Type", "Data", "TTL"})
		for i, val := range records {
			t.AppendRow([]interface{}{i, val.Name, val.Type, val.Data, val.Ttl})
		}
		t.Render()
	},
}

func init() {
	getCmd.AddCommand(recordsCmd)
	recordsCmd.Flags().StringVarP(&domain, "domain", "d", "", "domain name to look records for")
	recordsCmd.MarkFlagRequired("domain")
	recordsCmd.Flags().StringVar(&recordType, "type", "", "record types to get")
	recordsCmd.Flags().StringVar(&recordName, "name", "", "record name to get")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// domainsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// domainsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
