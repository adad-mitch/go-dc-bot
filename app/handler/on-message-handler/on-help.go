package onmessagehandlers

import (
	e "bott-the-pigeon/app/error"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Sends a bot usage help message from the provided bot.
func OnHelp(bot *discordgo.Session, msg *discordgo.MessageCreate) error {
	_, err := bot.ChannelMessageSend(msg.ChannelID,
		"Hello! My command prefix is a `>`.\n\n"+
			"**Commands**:\n"+
			"`support`: Sends bot usage help - like you're seeing right now!\n"+
			"`pigeon`: Sends a random picture of a pigeon.\n"+
			"`todo {Some feature}`: Submit a suggestion to the project todo list.\n\n"+
			"_More features coming soon_")
	if err != nil {
		err = fmt.Errorf("failed to send help message: %v", err)
		e.ThrowBotError(bot, msg.ChannelID, err)
		return err
	}
	return nil
}

// TODO: This is very hard-coded, and should be called from an API in some way.
