package main

import (
	"./sensu"
	"fmt"
	"github.com/spf13/cobra"
)

const (
	OHGI_REVISION = "__OHGI_REVISION__"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "ohgi",
		Short: "Sensu command-line tool by golang",
		Long:  "Sensu command-line tool by golang\nhttps://github.com/hico-horiuchi/ohgi",
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "events",
		Short: "List and resolve current events",
		Long:  "List and resolve current events",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s", sensu.GetEvents())
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print git revision of ohgi",
		Long:  "Print git revision of ohgi",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("ohgi revision %s\n", OHGI_REVISION)
		},
	})

	rootCmd.Execute()
}
