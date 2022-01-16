package commander

import "github.com/bwmarrin/discordgo"

// A struct that all components must implement
type BaseComponent struct {
	CustomID string                       // A unique string to identify the component internally and with Discord
	Disabled bool                         // If the component is disabled or not
	Run      func(context *Context) error // The handler function for the component
}

// Button component
type Button struct {
	BaseComponent

	Style discordgo.ButtonStyle     // The button style
	Label string                    // The button text
	Emoji *discordgo.ComponentEmoji // The emoji for the button
	URL   string                    // A URL for link buttons
}

// Select Menu component
type SelectMenu struct {
	BaseComponent

	Options     *[]discordgo.SelectMenuOption // Options for the select menu
	Placeholder string                        // Placeholder text
	MinValues   uint                          // Min values allowed to be selected
	MaxValues   uint                          // Max values allowed to be selected
}
