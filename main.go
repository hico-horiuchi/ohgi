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
		Use:   "clients [client]",
		Short: "Returns the list of clients",
		Long:  "clients           Returns the list of clients\nclients [client]  Returns a client",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Printf("%s", sensu.GetClients())
			case 1:
				fmt.Printf("%s", sensu.GetClientsClient(args[0]))
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "checks [check]",
		Short: "Returns the list of checks",
		Long:  "checks          Returns the list of checks\nchecks [check]  Returns a check",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Printf("%s", sensu.GetChecks())
			case 1:
				fmt.Printf("%s", sensu.GetChecksCheck(args[0]))
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "info",
		Short: "Returns the API info",
		Long:  "Returns the API info",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Printf("%s", sensu.GetInfo())
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
