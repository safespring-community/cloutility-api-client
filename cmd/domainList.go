/*
 * Copyright (c) Blue Safespring AB - Jan Johansson <jj@safespring.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
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
