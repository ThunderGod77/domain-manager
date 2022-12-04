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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		providerData, err := database.GetCredentials(domain)
		if err != nil {
			log.Println("failed to get records", err)
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
