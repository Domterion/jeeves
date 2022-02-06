package commander

import (
	"github.com/bwmarrin/discordgo"
)

// TODO: In the Context there should be a Data field that allows for any data, useful for stuff like databases

// A struct to provide useful information to command run functions
type CommandContext struct {
	Session         *discordgo.Session                                   // A pointer to the discordgo session
	Event           *discordgo.InteractionCreate                         // A pointer to the event that triggered the command
	Manager         *Manager                                             // A pointer to the command manager
	Name            string                                               // The command name used to invoke the command
	Options         []*discordgo.ApplicationCommandInteractionDataOption // The options the user passed
	ResolvedOptions *discordgo.ApplicationCommandInteractionDataResolved // Resolved options from Discord such as user and channel options
	Member          *discordgo.Member                                    // The member object for the command caller
}

// A struct to provide useful information to command run functions
type ComponentContext struct {
	Session *discordgo.Session           // A pointer to the discordgo session
	Event   *discordgo.InteractionCreate // A pointer to the event that triggered the command
	Manager *Manager                     // A pointer to the command manager
	Name    string                       // The command name used to invoke the command
	Member  *discordgo.Member            // The member object for the command caller
}

// Get a key from the dependency provider
func (c *CommandContext) Get(key string) interface{} {
	if c.Manager.options.DependencyProvider != nil {
		return c.Manager.options.DependencyProvider.Get(key)
	}

	return nil
}

// Get a key from the dependency provider
func (c *ComponentContext) Get(key string) interface{} {
	if c.Manager.options.DependencyProvider != nil {
		return c.Manager.options.DependencyProvider.Get(key)
	}

	return nil
}

// Respond to an interaction
func (c *CommandContext) Respond(response *discordgo.InteractionResponse) error {
	return c.Session.InteractionRespond(c.Event.Interaction, response)
}

// Respond to an interaction with plain text
func (c *CommandContext) RespondText(text string) error {
	return c.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: text,
		},
	})
}

// Respond to an interaction with plain text but an ephemeral message
func (c *CommandContext) RespondTextEphemeral(text string) error {
	return c.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: 1 << 6,
			Content: text,
		},
	})
}

// Defer a response
func (c *CommandContext) DeferResponse() error {
	return c.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
}

// Edit a response
func (c *CommandContext) ResponseEdit(response *discordgo.WebhookEdit) error {
	_, err := c.Session.InteractionResponseEdit(c.Session.State.User.ID, c.Event.Interaction, response)
	return err
}

// Respond to an interaction by editing text
func (c *CommandContext) ResponseEditText(text string) error {
	return c.ResponseEdit(&discordgo.WebhookEdit{
		Content: text,
	})
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
