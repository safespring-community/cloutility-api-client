/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/safespring/cloutility-api-client/cloutapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var osTypeListCmd = &cobra.Command{
	Use:   "list",
	Short: "ostype list will list the available OS types when enrolling a new backup node",
	Long: `
The command will list all the available Operating System types supported by the 
backup server.	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		osTypeList()
	},
}

func osTypeList() {
	client, err := cloutapi.Init(
		context.Background(),
		viper.GetString("client_id"),
		viper.GetString("client_origin"),
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetString("url"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	twriter := new(tabwriter.Writer)
	twriter.Init(os.Stdout, 8, 8, 1, '\t', 0)
	defer twriter.Flush()

	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "ID", "Name", "Short name", "Url")
	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "--", "----", "----------", "---")

	ostypes, err := client.GetNodeOperatingSystem()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, ostype := range ostypes {
		fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", ostype.ID, ostype.Name, ostype.ShortName, ostype.Href)
	}
}

func init() {
	osTypeCmd.AddCommand(osTypeListCmd)
}