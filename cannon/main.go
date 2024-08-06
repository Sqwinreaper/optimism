package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum-optimism/optimism/op-service/ctxinterrupt"
	"github.com/urfave/cli/v2"
	"os"

	"github.com/ethereum-optimism/optimism/cannon/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "cannon"
	app.Usage = "MIPS Fault Proof tool"
	app.Description = "MIPS Fault Proof tool"
	app.Commands = []*cli.Command{
		cmd.LoadELFCommand,
		cmd.WitnessCommand,
		cmd.RunCommand,
	}
	// This used to punt stdout on an interrupt, should that be a default behaviour?
	ctx := ctxinterrupt.WithSignalWaiterMain(context.Background())

	err := app.RunContext(ctx, os.Args)
	if err != nil {
		if errors.Is(err, ctx.Err()) {
			_, _ = fmt.Fprintf(os.Stderr, "command interrupted")
			os.Exit(130)
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v", err)
			os.Exit(1)
		}
	}
}
