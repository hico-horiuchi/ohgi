package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hico-horiuchi/ohgi/ohgi"
	"github.com/hico-horiuchi/ohgi/sensu"
	isatty "github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

var version string

func main() {
	var (
		age        int
		datacenter string
		limit      int
		offset     int
		delete     bool
		expiration string
		reason     string
		results    bool
		consumers  int
		messages   int
	)

	if !isatty.IsTerminal(os.Stdout.Fd()) {
		ohgi.EscapeSequence = false
	}

	rootCmd := &cobra.Command{
		Use:   "ohgi",
		Short: "Sensu command-line tool by Golang",
		Long:  "Sensu command-line tool by Golang\nhttps://github.com/hico-horiuchi/ohgi",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			sensu.DefaultAPI = ohgi.LoadConfig(datacenter)
		},
	}
	rootCmd.PersistentFlags().StringVarP(&datacenter, "datacenter", "x", "", "Specify a datacenter")

	clientsCmd := &cobra.Command{
		Use:   "clients [client]",
		Short: "List and delete client(s) information",
		Long:  "clients           Returns the list of clients\nclients [client]  Returns a client",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Print(ohgi.GetClients(sensu.DefaultAPI, limit, offset))
			case 1:
				if delete {
					fmt.Print(ohgi.DeleteClientsClient(sensu.DefaultAPI, args[0]))
				} else {
					if strings.Contains(args[0], "*") {
						fmt.Print(ohgi.GetClientsWildcard(sensu.DefaultAPI, args[0]))
					} else {
						fmt.Print(ohgi.GetClientsClient(sensu.DefaultAPI, args[0]))
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
		Use:   "jit [client] [address]",
		Short: "Dynamically created clients, added to the client registry",
		Long:  "jit [client] [address]                  Create or update client data\njit [client] [address] [subscriptions]  Create or update client data",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 2:
				fmt.Print(ohgi.PostClients(sensu.DefaultAPI, args[0], args[1], []string{}))
			case 3:
				fmt.Print(ohgi.PostClients(sensu.DefaultAPI, args[0], args[1], strings.Split(args[2], ",")))
			default:
				cmd.Help()
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "history [client]",
		Short: "Returns the history for a client",
		Long:  "history [client]  Returns the history for a client",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 1:
				fmt.Print(ohgi.GetClientsHistory(sensu.DefaultAPI, args[0]))
			default:
				cmd.Help()
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "checks [check]",
		Short: "List locally defined checks and request executions",
		Long:  "checks          Returns the list of checks\nchecks [check]  Returns a check",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Print(ohgi.GetChecks(sensu.DefaultAPI))
			case 1:
				if strings.Contains(args[0], "*") {
					fmt.Print(ohgi.GetChecksWildcard(sensu.DefaultAPI, args[0]))
				} else {
					fmt.Print(ohgi.GetChecksCheck(sensu.DefaultAPI, args[0]))
				}
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "request [check] [subscribers]",
		Short: "Issues a check execution request",
		Long:  "request [check]               Issues a check execution request\nrequest [check] [subscribers]  Issues a check execution request",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 1:
				fmt.Print(ohgi.PostRequest(sensu.DefaultAPI, args[0], []string{}))
			case 2:
				fmt.Print(ohgi.PostRequest(sensu.DefaultAPI, args[0], strings.Split(args[1], ",")))
			default:
				cmd.Help()
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
				fmt.Print(ohgi.GetEvents(sensu.DefaultAPI))
			case 1:
				if strings.Contains(args[0], "*") {
					fmt.Print(ohgi.GetEventsWildcard(sensu.DefaultAPI, args[0]))
				} else {
					fmt.Print(ohgi.GetEventsClient(sensu.DefaultAPI, args[0]))
				}
			case 2:
				if delete {
					fmt.Print(ohgi.DeleteEventsClientCheck(sensu.DefaultAPI, args[0], args[1]))
				} else {
					fmt.Print(ohgi.GetEventsClientCheck(sensu.DefaultAPI, args[0], args[1]))
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
				fmt.Print(ohgi.PostResolve(sensu.DefaultAPI, args[0], args[1]))
			default:
				cmd.Help()
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "results [client] [check]",
		Short: "List current check results",
		Long:  "results                   Returns a list of current check results for all clients\nresults [client]          Returns a list of current check results for a given client\nresults [client] [check]  Returns a check result for a given client & check name",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Print(ohgi.GetResults(sensu.DefaultAPI))
			case 1:
				if strings.Contains(args[0], "*") {
					fmt.Print(ohgi.GetResultsWildcard(sensu.DefaultAPI, args[0]))
				} else {
					fmt.Print(ohgi.GetResultsClient(sensu.DefaultAPI, args[0]))
				}
			case 2:
				fmt.Print(ohgi.GetResultsClientCheck(sensu.DefaultAPI, args[0], args[1]))
			}
		},
	})

	aggregatesCmd := &cobra.Command{
		Use:   "aggregates [check] [issued]",
		Short: "List and delete check aggregates",
		Long:  "aggregates                   Returns the list of aggregates\naggregates [check]           Returns the list of aggregates for a given check\naggregates [check] [issued]  Returns an aggregate",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Print(ohgi.GetAggregates(sensu.DefaultAPI, limit, offset))
			case 1:
				if delete {
					fmt.Print(ohgi.DeleteAggregatesCheck(sensu.DefaultAPI, args[0]))
				} else {
					fmt.Print(ohgi.GetAggregatesCheck(sensu.DefaultAPI, args[0], age))
				}
			case 2:
				fmt.Print(ohgi.GetAggregatesCheckIssued(sensu.DefaultAPI, args[0], args[1], results))
			}
		},
	}
	aggregatesCmd.Flags().IntVarP(&limit, "limit", "l", -1, "The number of aggregates to return")
	aggregatesCmd.Flags().IntVarP(&offset, "offset", "o", -1, "The number of aggregates to offset before returning items")
	aggregatesCmd.Flags().IntVarP(&age, "age", "a", -1, "The number of seconds old an aggregate must be to be listed")
	aggregatesCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Deletes all aggregates for a check")
	aggregatesCmd.Flags().BoolVarP(&results, "results", "r", false, "Return the raw result data")
	rootCmd.AddCommand(aggregatesCmd)

	silenceCmd := &cobra.Command{
		Use:   "silence [client] [check]",
		Short: "Create, list, and delete silence stashes",
		Long:  "silence                   Returns a list of silence stashes\nsilence [client]          Create a silence stash\nsilence [client] [check]  Create a silence stash",
		Run: func(cmd *cobra.Command, args []string) {
			switch len(args) {
			case 0:
				fmt.Print(ohgi.GetSilence(sensu.DefaultAPI))
			case 1:
				if delete {
					fmt.Print(ohgi.DeleteSilence(sensu.DefaultAPI, args[0], ""))
				} else {
					fmt.Print(ohgi.PostSilence(sensu.DefaultAPI, args[0], "", expiration, reason))
				}
			case 2:
				if delete {
					fmt.Print(ohgi.DeleteSilence(sensu.DefaultAPI, args[0], args[1]))
				} else {
					fmt.Print(ohgi.PostSilence(sensu.DefaultAPI, args[0], args[1], expiration, reason))
				}
			}
		},
	}
	silenceCmd.Flags().StringVarP(&expiration, "expiration", "e", "", "e.g. 15m, 1h, 1d")
	silenceCmd.Flags().StringVarP(&reason, "reason", "r", "", "Enter a reason")
	silenceCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Remove silence stash")
	rootCmd.AddCommand(silenceCmd)

	healthCmd := &cobra.Command{
		Use:   "health",
		Short: "Check the status of the API's transport & Redis connections, and query the transport's status",
		Long:  "health  Returns health information on transport & Redis connections",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(ohgi.GetHealth(sensu.DefaultAPI, consumers, messages))
		},
	}
	healthCmd.Flags().IntVarP(&consumers, "consumers", "c", 1, "The minimum number of transport consumers to be considered healthy")
	healthCmd.Flags().IntVarP(&messages, "messages", "m", 0, "The maximum ammount of transport queued messages to be considered healthy")
	rootCmd.AddCommand(healthCmd)

	rootCmd.AddCommand(&cobra.Command{
		Use:   "info",
		Short: "List the Sensu version and the transport and Redis connection information",
		Long:  "info  Returns information on the API",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(ohgi.GetInfo(sensu.DefaultAPI))
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print and check the version of ohgi",
		Long:  "version  Print and check the version of ohgi",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(ohgi.Version(version))
		},
	})

	rootCmd.Execute()
}
