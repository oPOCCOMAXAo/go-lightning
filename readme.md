# Lightning image generator

Example output:

![Image](example.png)

## Build:

```bash
# windows
go build -o .bin/app.exe main.go

# linux
go build -o .bin/app main.go
```

## Usage:

```bash
# print all flags
app --help

# generating
app --width 1920 --height 1080 --file fullhdlightning.png
```
