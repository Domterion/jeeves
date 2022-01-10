package commander

import "github.com/bwmarrin/discordgo"

// A struct to provide useful information to command run functions
type Context struct {
	Session         *discordgo.Session                                   // A pointer to the discordgo session
	Event           *discordgo.InteractionCreate                         // A pointer to the event that triggered the command
	Options         []*discordgo.ApplicationCommandInteractionDataOption // The options the user passed
	ResolvedOptions *discordgo.ApplicationCommandInteractionDataResolved // Resolved options from Discord such as user and channel options
	Member          *discordgo.Member                                    // The member object for the command caller
	User            *discordgo.User                                      // The user object for the command caller
}

func (c *Context) Respond(response *discordgo.InteractionResponse) {
	c.Session.InteractionRespond(c.Event.Interaction, response)
}

func (c *Context) RespondText(text string) {
	c.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: text,
		},
	})
}