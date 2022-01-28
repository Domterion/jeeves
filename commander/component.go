package commander

import "github.com/bwmarrin/discordgo"

// A struct that all components must implement
type BaseComponent struct {
	CustomID string                       // A unique string to identify the component internally and with Discord
	Disabled bool                         // If the component is disabled or not
	Run      func(context *ComponentContext) error // The handler function for the component
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
	MinValues   int                           // Min values allowed to be selected
	MaxValues   int                           // Max values allowed to be selected
}

// An Action Row Component
type ActionRow struct {
	components []interface{} // All components the action row holds
}

// A utility struct
type Components struct {
	components []interface{} // All components the components struct knows of
}

// TODO: This is just a placeholder, hopefully I figure it out later
func NewComponents() *Components {
	return &Components{}
}

func (c *Components) AddButton(button Button) {
	c.components = append(c.components, button)
}

func (c *Components) AddSelectMenu(selectMenu SelectMenu) {
	c.components = append(c.components, selectMenu)
}

func (c *Components) AddActionRow(actionRow ActionRow) {
	c.components = append(c.components, actionRow)
}

func (b Button) ToMessageComponent() discordgo.Button {
	return discordgo.Button{
		Emoji:    *b.Emoji,
		Label:    b.Label,
		CustomID: b.BaseComponent.CustomID,
		Style:    b.Style,
		Disabled: b.BaseComponent.Disabled,
	}
}

func (s SelectMenu) ToMessageComponent() discordgo.SelectMenu {
	return discordgo.SelectMenu{
		CustomID:    s.BaseComponent.CustomID,
		Placeholder: s.Placeholder,
		MinValues:   s.MinValues,
		MaxValues:   s.MaxValues,
		Options:     *s.Options,
	}
}

func (a ActionRow) ToMessageComponent() discordgo.ActionsRow {
	components := []discordgo.MessageComponent{}

	for _, component := range a.components {
		switch c := component.(type) {
		case Button:
			components = append(components, c.ToMessageComponent())
		case SelectMenu:
			components = append(components, c.ToMessageComponent())
		}
	}

	return discordgo.ActionsRow{Components: components}
}

func (a *ActionRow) AddButton(button Button) {
	a.components = append(a.components, button)
}

func (a *ActionRow) AddSelectMenu(selectMenu SelectMenu) {
	a.components = append(a.components, selectMenu)
}

func (c *Components) ToMessageComponent() []discordgo.MessageComponent {
	components := []discordgo.MessageComponent{}

	for _, component := range c.components {
		switch c := component.(type) {
		case Button:
			components = append(components, c.ToMessageComponent())
		case SelectMenu:
			components = append(components, c.ToMessageComponent())
		case ActionRow:
			components = append(components, c.ToMessageComponent())
		}
	}

	return components
}
