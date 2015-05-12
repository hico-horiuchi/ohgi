package main

import (
	"fmt"
	"os"
	"strings"

	"./ohgi"
	"github.com/mattn/go-isatty"
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

	ohgi.LoadConfig()
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		ohgi.EscapeSequence = false
	}

	rootCmd := &cobra.Command{
		Use:   "ohgi",
		Short: "Sensu command-line tool by golang",
		Long:  "Sensu command-line tool by golang\nhttps://github.com/hico-horiuchi/ohgi",
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "checks [check]",
		Short: "List locally defined checks and request executions",
		Long:  "checks          Returns the list of checks\nchecks [check]  Returns a check",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Print(ohgi.GetChecks())
			case 1:
				if strings.Contains(args[0], "*") {
					fmt.Print(ohgi.GetChecksWildcard(args[0]))
				} else {
					fmt.Print(ohgi.GetChecksCheck(args[0]))
				}
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
				fmt.Print(ohgi.PostRequest(args[0], ""))
			case 2:
				fmt.Print(ohgi.PostRequest(args[0], args[1]))
			}
		},
	})

	clientsCmd := &cobra.Command{
		Use:   "clients [client]",
		Short: "List and delete client(s) information",
		Long:  "clients           Returns the list of clients\nclients [client]  Returns a client",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Print(ohgi.GetClients(limit, offset))
			case 1:
				if delete {
					fmt.Print(ohgi.DeleteClientsClient(args[0]))
				} else {
					if strings.Contains(args[0], "*") {
						fmt.Print(ohgi.GetClientsWildcard(args[0]))
					} else {
						fmt.Print(ohgi.GetClientsClient(args[0]))
					}
				}
			}
		},
	}
	clientsCmd.Flags().IntVarP(&limit, "limit", "l", -1, "The number of clients to return")
	clientsCmd.Flags().IntVarP(&offset, "offset", "o", -1, "The number of clients to offset before returning items")
	clientsCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Removes a client, resolving its current events")
	rootCmd.AddCommand(clientsCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "history [client]",
		Short: "Returns the history for a client",
		Long:  "history [client]  Returns the history for a client",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 1:
				fmt.Print(ohgi.GetHistory(args[0]))
			}
		},
	})

	eventsCmd := &cobra.Command{
		Use:   "events [client] [check]",
		Short: "List and resolve current events",
		Long:  "events                   Returns the list of current events\nevents [client]          Returns the list of current events for a given client\nevents [client] [check]  Returns an event for a given client & check name",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Print(ohgi.GetEvents())
			case 1:
				fmt.Print(ohgi.GetEventsClient(args[0]))
			case 2:
				if delete {
					fmt.Print(ohgi.DeleteEventsClientCheck(args[0], args[1]))
				} else {
					fmt.Print(ohgi.GetEventsClientCheck(args[0], args[1]))
				}
			}
		},
	}
	eventsCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Resolves an event for a given check on a given client")
	rootCmd.AddCommand(eventsCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "resolve [client] [check]",
		Short: "Resolves an event",
		Long:  "resolve [client] [check]  Resolves an event",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 2:
				fmt.Print(ohgi.PostResolve(args[0], args[1]))
			}
		},
	})

	healthCmd := &cobra.Command{
		Use:   "health",
		Short: "Check the status of the API's transport & Redis connections, and query the transport's status",
		Long:  "health  Returns health information on transport & Redis connections",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(ohgi.GetHealth(consumers, messages))
		},
	}
	healthCmd.Flags().IntVarP(&consumers, "consumers", "c", 1, "The minimum number of transport consumers to be considered healthy")
	healthCmd.Flags().IntVarP(&messages, "messages", "m", 1, "The maximum ammount of transport queued messages to be considered healthy")
	rootCmd.AddCommand(healthCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "info",
		Short: "List the Sensu version and the transport and Redis connection information",
		Long:  "info  Returns information on the API",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(ohgi.GetInfo())
		},
	})

	silenceCmd := &cobra.Command{
		Use:   "silence [client] [check]",
		Short: "Create, list, and delete silences",
		Long:  "silence                   Returns a list of silences\nsilence [client]          Create a silence\nsilence [client] [check]  Create a silence",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Print(ohgi.GetSilence())
			case 1:
				if delete {
					fmt.Print(ohgi.DeleteSilence(args[0], ""))
				} else {
					fmt.Print(ohgi.PostSilence(args[0], "", expiration, reason))
				}
			case 2:
				if delete {
					fmt.Print(ohgi.DeleteSilence(args[0], args[1]))
				} else {
					fmt.Print(ohgi.PostSilence(args[0], args[1], expiration, reason))
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
		Short: "Print and check version of ohgi",
		Long:  "Print and check version of ohgi",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(ohgi.Version(version))
		},
	})

	rootCmd.Execute()
}
