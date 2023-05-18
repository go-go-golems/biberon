package cmds

import (
	"fmt"
	"github.com/caltechlibrary/bibtex"
	"github.com/go-go-golems/glazed/pkg/cli"
	"github.com/go-go-golems/glazed/pkg/middlewares/table"
	"github.com/spf13/cobra"
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

		gp, err := cli.CreateGlazedProcessorFromCobra(cmd)
		cobra.CheckErr(err)

		gp.OutputFormatter().AddTableMiddleware(table.NewReorderColumnOrderMiddleware([]string{"id", "type", "keys", "title", "author", "year"}))

		ctx := cmd.Context()

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
				err = gp.ProcessInputObject(ctx, row)
				cobra.CheckErr(err)
			}

		}

		err = gp.OutputFormatter().Output(ctx, os.Stdout)
		cobra.CheckErr(err)
	},
}

func init() {
	BibtexCmd.Flags().SortFlags = false
	err := cli.AddGlazedProcessorFlagsToCobraCommand(BibtexCmd)
	if err != nil {
		panic(err)
	}
}
