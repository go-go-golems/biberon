package cmds

import (
	"fmt"
	"github.com/caltechlibrary/bibtex"
	"github.com/spf13/cobra"
	"github.com/wesen/glazed/pkg/cli"
	"github.com/wesen/glazed/pkg/middlewares"
	"os"
)

var BibtexCmd = &cobra.Command{
	Use:   "bibtex",
	Short: "Import/export bibtex data",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No input file specified")
			os.Exit(1)
		}

		gp, of, err := cli.SetupProcessor(cmd)
		cobra.CheckErr(err)

		gp.OutputFormatter().AddTableMiddleware(middlewares.NewReorderColumnOrderMiddleware([]string{"id", "type", "keys", "title", "author", "year"}))

		for _, arg := range args {
			buf, err := os.ReadFile(arg)
			cobra.CheckErr(err)

			elts, err := bibtex.Parse(buf)
			cobra.CheckErr(err)

			for _, e := range elts {
				row := make(map[string]interface{})
				row["id"] = e.ID
				row["type"] = e.Type
				row["keys"] = e.Keys
				for k, v := range e.Tags {
					// strip { and } from v
					cleanV := v
					if len(cleanV) > 1 {
						if cleanV[0] == '{' && cleanV[len(cleanV)-1] == '}' {
							cleanV = cleanV[1 : len(cleanV)-1]
						}
					}
					row[k] = cleanV
				}
				err = gp.ProcessInputObject(row)
				cobra.CheckErr(err)
			}

		}

		s, err := of.Output()
		cobra.CheckErr(err)
		fmt.Print(s)

	},
}

func init() {
	BibtexCmd.Flags().SortFlags = false
	cli.AddFlags(BibtexCmd, cli.NewFlagsDefaults())
}
