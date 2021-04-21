package messages

var RootCommandLong = `
=====================================================================================
                            all-things-go
=====================================================================================
This go module is a place for me to put all of my go lang stuff in.

The structure is to use cobra to call each leason. So I'll create a cobra command
for the thing I want to learn, Like logurus and and then add what would be the 
'main' function for it in the 'funcs' dir and then rather than the main() I create
RunLogurus() and then add funcs.RunLogurus to the sub command.

Directory structure is, for example:

/gocron                      # Top Level
  /cmd                       # Cobra puts commands here
    /funcs                   # Each example goes in specific name.go here
      cron.go
      logurus.go
    /messages                # Place to store long messages
      cobra.go
    root.go
    cron.go
    logurus.go

Steps:
  1. Create a cobra subcommand: cobra command <name>
  2. Create a <name>.go file in the funcs directory.
  3. Add func Run<Name> () {} to that file.
  4. Update the call to the subcommand to Run<Name>()
=====================================================================================
`

var CronCommandLong = `
=====================================================================================
                            cron
=====================================================================================
This module does commands like cron.
=====================================================================================
`

var LogurusCommandLong = `
=====================================================================================
                            logurus
=====================================================================================
This module is a better logger than the default and does appear to be what most
people use.
=====================================================================================
`

var ConfigCommandLong = `
=====================================================================================
                            config
=====================================================================================
This will show the Viper configuration where ever it is.
Cobra was initialized to use cobra. You can see where the default is in this help.

The default is $HOME/.all-things-go.yaml
=====================================================================================
`

var ConfigClickhouse = `
=====================================================================================
                            clickhouse
=====================================================================================

In order to use clickhouse you need to forward the port from the container, eg.

kubectl port-forward --namespace default svc/core-clickhouse-default-1 9000

=====================================================================================
`
