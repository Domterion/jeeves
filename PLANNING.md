<div align="center">
    <h1>Planning</h1>
    This is the file where I keep track of ideas aswell as my thoughts when writing this bot
    <br>
    <br>
</div>

# Ideas

- Space Exploration/RPG
    - Jeeves is your captain, maybe he can be an AI helping to guide the player
    - Interplanet travel, you can travel between planets but this consumes rocket fuel
        - Each Rocket has a different tank size allowing for further and more exploration
            - You must start a journey with enough fuel to make it to the next planet
            - Fuel isnt cheap
        - Maybe theres obstacles between planets that you have to destroy?
        - Varying distances taking more time and fuel to get to some as opposed to others
    - Rocket ships are upgradeable and you start with a low-end version
        - You can only get a bigger fuel tank by upgrading the whole rocket
        - Upgradable armor and shields to protect when travelling between planets
        - Upgradable weapons to fight off attackers
        - All rockets will have a safe for players to securely keep money in 
            - This safe size is static and the same for all ships
    - Fighting off enemies trying to break into your ship
        - When travelling between planets you may encounter enemies attacking you and your ship
        - They will steal try and attack your fuel supply and can leave you stranded
            - If stranded you...
    - Players
        - Each player can get better armor/weapons or upgrade their current ones 
        - Specific amount of health
            - Can be replenished with food/water/rest
        - Maybe limited inventory space or space limited on a backpack or ship size
        - Currency name is specks
            - Money can be stolen by invaders
            - Used to purchase food, water, items, upgrades and various other things
            - Can be deposited into rocket safe for extra security

Enemies are sometimes referred to as invaders and will attack your ship to steal your money and fuel. You can fight them off and upgrade your rockets defenses to prevent entry

Rocket ships can also be called rockets or ships

The currency is specks or SPC but is also referred to as money, cash or currency

# Thoughts

The `handler`, aka the command handler, should define three different command interface: `Command`, `SubCommmand` and `SubCommandGroup`. 

All commands must have a `Name() string`, `Description() string` and `Run(context *handler.Context) error` functions

The type of the command will be determine by the functions defined:

`Command` has `SubCommands() *[]SubCommand` and `SubCommandGroups() *[]SubCommandGroup`

`SubCommandGroup` has `SubCommands() *[]SubCommand`

`SubCommand` has none

...so basically the above was scrapped because using interfaces doesnt work the best for this. It now uses structs and is fully working, I hope atleast...

Next we are on to storage of commands which should be simple enough...

...well, it wasnt as simple as I thought but `interface{}` is beautiful when used properly but now you can call `RegisterCommand` and itll store all commands and subcommands in a `map[string]interface{}`. 

