<div align="center">
    <h1>💻 Commander</h1>
    Commander is an interaction and command framework for <a href="https://github.com/bwmarrin/discordgo">Discordgo</a>.
    <br>
    <br>
</div>

# Usage

> Creating a new command
>
> This creates a basic `ping` command that responds with `pong!`
```go
...imports

var PingCommand commander.Command = commander.Command{
	BaseCommand: commander.BaseCommand{
		Name:        "ping",
		Description: "pong",
		Type:        discordgo.ChatApplicationCommand,
		Run: func(context *commander.Context) error {
			return context.RespondText("pong!")
		},
	},
}
```

more examples can be found in `examples/`

# TODO

- [ ] Interaction handling such as buttons, currently only slash commands are supported and handled.
- [ ] More concise error handling and reporting
- [ ] Move commander to a separate repository and make it a proper package

# License
MIT