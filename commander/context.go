package commander

import (
	"github.com/bwmarrin/discordgo"
)

// A struct to provide useful information to command run functions
type Context struct {
	Session         *discordgo.Session                                   // A pointer to the discordgo session
	Event           *discordgo.InteractionCreate                         // A pointer to the event that triggered the command
	Manager         *Manager                                             // A pointer to the command manager
	Name            string                                               // The command name used to invoke the command
	Options         []*discordgo.ApplicationCommandInteractionDataOption // The options the user passed
	ResolvedOptions *discordgo.ApplicationCommandInteractionDataResolved // Resolved options from Discord such as user and channel options
	Member          *discordgo.Member                                    // The member object for the command caller
}

// Respond to an interaction
func (c *Context) Respond(response *discordgo.InteractionResponse) error {
	return c.Session.InteractionRespond(c.Event.Interaction, response)
}

// Respond to an interaction with plain text
func (c *Context) RespondText(text string) error {
	return c.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: text,
		},
	})
}

// A struct to provide useful information to command run functions
type ComponentContext struct {
	Session         *discordgo.Session                                   // A pointer to the discordgo session
	Event           *discordgo.InteractionCreate                         // A pointer to the event that triggered the command
	Manager         *Manager                                             // A pointer to the command manager
	Name            string                                               // The command name used to invoke the command
	Member          *discordgo.Member                                    // The member object for the command caller
}

// Respond to an interaction
func (c *ComponentContext) Respond(response *discordgo.InteractionResponse) error {
	return c.Session.InteractionRespond(c.Event.Interaction, response)
}

// Respond to an interaction with plain text
func (c *ComponentContext) RespondText(text string) error {
	return c.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: text,
		},
	})
}

