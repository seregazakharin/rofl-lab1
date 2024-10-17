package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/BaldiSlayer/rofl-lab1/internal/app/mclient"
	"github.com/BaldiSlayer/rofl-lab1/internal/app/tgbot"
)

func cli() {
	r := bufio.NewReader(os.Stdin)
	data, err := io.ReadAll(r)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	model, err := mclient.NewMistralClient()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("Executing model request")

	answer, err := model.Ask(string(data))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	fmt.Println(answer)
}

func main() {
	useCli := flag.Bool("cli", false, "run with cli interface")

	flag.Parse()

	if *useCli {
		cli()
		return
	}

	app := tgbot.New(
		tgbot.WithConfig(),
	)

	app.Run()
}
