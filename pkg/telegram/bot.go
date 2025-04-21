package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Bot wraps the Telegram bot functionality
type Bot struct {
	api       *tgbotapi.BotAPI
	channelID string
}

// New creates a new Telegram bot instance
func New(token, channelID string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api:       api,
		channelID: channelID,
	}, nil
}

// SendMessage sends a message to the configured channel
func (b *Bot) SendMessage(message string) error {
	msg := tgbotapi.NewMessageToChannel(b.channelID, message)
	msg.ParseMode = "Markdown"
	_, err := b.api.Send(msg)
	if err != nil {
		log.Printf("Failed to send Telegram message: %v", err)
		return err
	}
	return nil
}

// SendStartupMessage sends the initial startup message
func (b *Bot) SendStartupMessage(chains []string) error {
	startupMsg := "🤖 TokenRegistry Monitor Bot is online!\n\n📊 Monitoring on chains:\n"
	for _, chain := range chains {
		startupMsg += fmt.Sprintf("- %s\n", chain)
	}
	return b.SendMessage(startupMsg)
} 