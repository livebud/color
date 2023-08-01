# Color

[![Go Reference](https://pkg.go.dev/badge/github.com/livebud/color.svg)](https://pkg.go.dev/github.com/livebud/color)

Minimal color package

![color](https://github.com/livebud/color/assets/170299/333ea10c-d953-48eb-9071-b9fa069f2595)

## Features

- Simple API
- Supports `NO_COLOR`

## Install

```
go get github.com/livebud/color
```

## Example

```go
fmt.Fprintln(os.Stdout,
  color.Red("h"),
  color.Blue("e"),
  color.Yellow("l"),
  color.Green("l"),
  color.Pink("o"),
)
```

## Contributors

- Matt Mueller ([@mattmueller](https://twitter.com/mattmueller))

## License

MIT
