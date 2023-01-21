package commands

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"com.son.server.attendaceservice/configs"
	"com.son.server.attendaceservice/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:          "serve",
	Short:        "Run the Goginbase server",
	RunE:         serverCmdF,
	SilenceUsage: true,
}

func init() {
	RootCmd.AddCommand(serverCmd)
	RootCmd.RunE = serverCmdF
}

func serverCmdF(command *cobra.Command, args []string) error {
	interruptChan := make(chan os.Signal, 1)
	return runServer(interruptChan)
}

func runServer(interruptChan chan os.Signal) error {
	// Setting the highest traceback level from the code.
	// This is done to print goroutines from all threads (see golang.org/issue/13161)
	// and also preserve a crash dump for later investigation.
	debug.SetTraceback("cash")

	environment := flag.String("e", "development", "")
	// environment := flag.String("e", "production", "")
	// environment := flag.String("e", "test", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	configs.Init(*environment)
	server.Init()

	// wait for kill signal before attempting to gracefully shutdown
	// the running service
	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan
	return nil
}
