package main

import (
	"./sensu"
	"fmt"
	"github.com/spf13/cobra"
)

var version string

func main() {
	var (
		limit      int
		offset     int
		delete     bool
		consumers  int
		messages   int
		expiration string
		reason     string
	)

	rootCmd := &cobra.Command{
		Use:   "ohgi",
		Short: "Sensu command-line tool by golang",
		Long:  "Sensu command-line tool by golang\nhttps://github.com/hico-horiuchi/ohgi",
	}

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
		Use:   "request [check] [subscriber]",
		Short: "Issues a check execution request",
		Long:  "request [check]               Issues a check execution request\nrequest [check] [subscriber]  Issues a check execution request",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 1:
				fmt.Printf("%s", sensu.PostRequest(args[0], ""))
			case 2:
				fmt.Printf("%s", sensu.PostRequest(args[0], args[1]))
			}
		},
	})

	clientsCmd := &cobra.Command{
		Use:   "clients [client]",
		Short: "Returns the list of clients",
		Long:  "clients           Returns the list of clients\nclients [client]  Returns a client",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Printf("%s", sensu.GetClients(limit, offset))
			case 1:
				if delete {
					fmt.Printf("%s", sensu.DeleteClientsClient(args[0]))
				} else {
					fmt.Printf("%s", sensu.GetClientsClient(args[0]))
				}
			}
		},
	}
	clientsCmd.Flags().IntVarP(&limit, "limit", "l", -1, "The number of clients to return")
	clientsCmd.Flags().IntVarP(&offset, "offset", "o", -1, "The number of clients to offset before returning items")
	clientsCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Removes a client, resolving its current events (delayed action)")
	rootCmd.AddCommand(clientsCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "history [client]",
		Short: "Returns the client history",
		Long:  "history [client]  Returns the client history",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 1:
				fmt.Printf("%s", sensu.GetHistory(args[0]))
			}
		},
	})

	eventsCmd := &cobra.Command{
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
				if delete {
					fmt.Printf("%s", sensu.DeleteEventsClientCheck(args[0], args[1]))
				} else {
					fmt.Printf("%s", sensu.GetEventsClientCheck(args[0], args[1]))
				}
			}
		},
	}
	eventsCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Resolves an event (delayed action)")
	rootCmd.AddCommand(eventsCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "resolve [client] [check]",
		Short: "Resolves an event (delayed action)",
		Long:  "resolve [client] [check]  Resolves an event (delayed action)",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 2:
				fmt.Printf("%s", sensu.PostResolve(args[0], args[1]))
			}
		},
	})

	healthCmd := &cobra.Command{
		Use:   "health",
		Short: "Returns the API info",
		Long:  "Returns the API info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s", sensu.GetHealth(consumers, messages))
		},
	}
	healthCmd.Flags().IntVarP(&consumers, "consumers", "c", 1, "The minimum number of transport consumers to be considered healthy")
	healthCmd.Flags().IntVarP(&messages, "messages", "m", 1, "The maximum number of transport queued messages to be considered healthy")
	rootCmd.AddCommand(healthCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "info",
		Short: "Returns the API info",
		Long:  "Returns the API info",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s", sensu.GetInfo())
		},
	})

	silenceCmd := &cobra.Command{
		Use:   "silence [client] [check]",
		Short: "Returns a list of silences",
		Long:  "silence                   Returns a list of silences\nsilence [client]          Create a silence\nsilence [client] [check]  Create a silence",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Printf("%s", sensu.GetSilence())
			case 1:
				if delete {
					fmt.Printf("%s", sensu.DeleteSilence(args[0], ""))
				} else {
					fmt.Printf("%s", sensu.PostSilence(args[0], "", expiration, reason))
				}
			case 2:
				if delete {
					fmt.Printf("%s", sensu.DeleteSilence(args[0], args[1]))
				} else {
					fmt.Printf("%s", sensu.PostSilence(args[0], args[1], expiration, reason))
				}
			}
		},
	}
	silenceCmd.Flags().StringVarP(&expiration, "expiration", "e", "", "15m, 1h, 1d")
	silenceCmd.Flags().StringVarP(&reason, "reason", "r", "", "Enter a reason")
	silenceCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Delete a silence")
	rootCmd.AddCommand(silenceCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print ohgi version",
		Long:  "Print ohgi version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("ohgi version %s\n", version)
		},
	})

	rootCmd.Execute()
}
