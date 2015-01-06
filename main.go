package main

import (
	"./sensu"
	"fmt"
	"github.com/spf13/cobra"
)

var revision string

func main() {
	var client string

	rootCmd := &cobra.Command{
		Use:   "ohgi",
		Short: "Sensu command-line tool by golang",
		Long:  "Sensu command-line tool by golang\nhttps://github.com/hico-horiuchi/ohgi",
	}

	eventsCmd := &cobra.Command{
		Use:   "events [-c client]",
		Short: "List and resolve current events",
		Long:  "List and resolve current events",
		Run: func(cmd *cobra.Command, args []string) {
			if client != "" {
				fmt.Printf("%s", sensu.GetEventsClient(client))
			} else {
				fmt.Printf("%s", sensu.GetEvents())
			}
		},
	}
	eventsCmd.Flags().StringVarP(&client, "client", "c", "", "Returns the list of current events for a client")
	rootCmd.AddCommand(eventsCmd)

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
