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
	rootCmd := &cobra.Command{Use: "ohgi"}

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
		Short: "Print ohgi revision",
		Long:  "Print ohgi revision",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("ohgi revision %s\n", OHGI_REVISION)
		},
	})
	rootCmd.Execute()
}
