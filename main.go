package main

import (
	"os"

	"github.com/renanqts/telegram-sender/pkg/config"
	"github.com/renanqts/telegram-sender/pkg/telegram"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func main() {

	cfg := config.NewConfig()

	flags := []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "telegram-token",
			Usage:       "Telegram token generate in Bot father",
			Destination: &cfg.TelegramToken}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "telegram-chat-id",
			Usage:       "Telegram chat ID to post the message",
			Destination: &cfg.TelegramChatID}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "telegram-message",
			Usage:       "Message to post in Telegram",
			Destination: &cfg.TelegramMessage}),
		&cli.StringFlag{
			Name:  "load",
			Usage: "Load configuration from `FILE`",
		},
	}

	app := &cli.App{
		Name:  "telegram-sender",
		Usage: "Send a message using a Telegram bot",
	}

	app.Action = func(context *cli.Context) error {

		// Set config with the content read in the config file
		if context.IsSet("load") {
			cfg.TelegramChatID = context.String("telegram-chat-id")
			cfg.TelegramToken = context.String("telegram-token")
			cfg.TelegramMessage = context.String("telegram-message")
		}

		return execute(cfg)
	}

	app.Before = altsrc.InitInputSourceWithContext(
		flags,
		func(context *cli.Context) (altsrc.InputSourceContext, error) {
			if context.IsSet("load") {
				return altsrc.NewYamlSourceFromFile(context.String("load"))
			}

			return &altsrc.MapInputSource{}, nil
		},
	)

	app.Flags = flags

	app.Run(os.Args)
}

func execute(cfg *config.Config) error {

	// send telegram message
	err := telegram.SendMessage(cfg.TelegramToken, cfg.TelegramChatID, cfg.TelegramMessage)
	if err != nil {
		return err
	}

	return nil
}
