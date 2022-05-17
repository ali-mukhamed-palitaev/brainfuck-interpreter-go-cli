# brainfuck-interpreter-go-cli
Brainfuck interpreter cli

It is a brainfuck interpreter cli.

Here is the steps you need to start using it:

1) First of all, build the brainfuck-interpreter-go-cli:

`go build`

2) Write you commands in brainfuck language in the file (i.e. example.bf)

3) Run the command using brainfuck-interpreter-go-cli, don't forget to provide the name of the file with brainfuck code:

`./brainfuck-interpreter-go-cli example.bf`

If you want to add a custom command, create a new one in `custom_commands.go`. There is an example of a custom command
in the file. After that run the command AddCommand in `main.go` before running the Execute command and provide
a commandName and a Command implementation itself.

`interpreter.AddCommand("%", Remainder)`

If you want to remove a custom command, run the command RemoveCommand in `main.go` and provide a commandName.

`interpreter.RemoveCommand("%")`
