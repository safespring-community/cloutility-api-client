/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listDomainsCmd represents the listDomains command
var domainListCmd = &cobra.Command{
	Use:   "list",
	Short: "domain list will list the available backup domains",
	Long: `
The domain list command will list all available backup domains supported by
the server.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		domainList()
	},
}

func domainList() {
	twriter := new(tabwriter.Writer)
	twriter.Init(os.Stdout, 8, 8, 1, '\t', 0)
	defer twriter.Flush()

	user, err := client.GetUser()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "ID", "Name", "Description", "Url")
	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "--", "----", "-----------", "---")

	domains, err := client.GetDomains(user.UserBUnit.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, domain := range domains {
		fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", domain.ID, domain.Name, domain.Description, domain.Href)
	}
}

func init() {
	domainCmd.AddCommand(domainListCmd)
}
