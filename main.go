package main

import (
	"./sensu"
	"fmt"
	"github.com/spf13/cobra"
)

var revision string

func main() {
	var client string
	var check string

	rootCmd := &cobra.Command{
		Use:   "ohgi",
		Short: "Sensu command-line tool by golang",
		Long:  "Sensu command-line tool by golang\nhttps://github.com/hico-horiuchi/ohgi",
	}

	eventsCmd := &cobra.Command{
		Use:   "events",
		Short: "List and resolve current events",
		Long:  "List and resolve current events",
		Run: func(cmd *cobra.Command, args []string) {
			if check != "" && client != "" {
				fmt.Printf("%s", sensu.GetEventsClientCheck(client, check))
			} else if client != "" {
				fmt.Printf("%s", sensu.GetEventsClient(client))
			} else {
				fmt.Printf("%s", sensu.GetEvents())
			}
		},
	}
	eventsCmd.Flags().StringVarP(&client, "client", "l", "", "Returns the list of current events for a client")
	eventsCmd.Flags().StringVarP(&check, "check", "c", "", "Returns an event")
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
