package main

import (
	"./sensu"
	"fmt"
	"github.com/spf13/cobra"
)

var revision string

func main() {
	rootCmd := &cobra.Command{
		Use:   "ohgi",
		Short: "Sensu command-line tool by golang",
		Long:  "Sensu command-line tool by golang\nhttps://github.com/hico-horiuchi/ohgi",
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "events [client] [check]",
		Short: "List and resolve current events",
		Long:  "events                   List and resolve current events\nevents [client]          Returns the list of current events for a client\nevents [client] [check]  Returns an event",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Printf("%s", sensu.GetEvents())
			case 1:
				fmt.Printf("%s", sensu.GetEventsClient(args[0]))
			case 2:
				fmt.Printf("%s", sensu.GetEventsClientCheck(args[0], args[1]))
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print git revision of ohgi",
		Long:  "Print git revision of ohgi",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("ohgi revision %s\n", revision)
		},
	})

	rootCmd.Execute()
}
