# Statgen

Statgen is a simple and efficient static site generator written in Go. With Statgen, you can easily create static websites from markdown files, allowing you to focus on content creation without worrying about the complexities of web development.
This is a Go rewrite after having the first version written in Python following steps from [Boot.dev](https://boot.dev/)

## Features

- **Markdown Support:** Write your content in markdown format, a simple and intuitive markup language.
- **Fast Rendering:** Quickly generate static websites without the need for server-side processing.
- **Command-Line Interface:** Use the command-line interface to generate your website with a single command.

## Installation

To install Statgen, simply use `go get`:

```bash
go get -u github.com/Shobhit-Nagpal/statgen
```

Make sure your Go environment is set up correctly.

## Usage

Here's a basic example of how to use Statgen:

```bash
./main.sh
```

This command runs your static website using the markdown files in the `content` directory and the templates in the `templates` directory. The generated website will be placed in the `public` directory.

## Contributing

Contributions are welcome! If you have ideas for new features, improvements, or bug fixes, feel free to open an issue or submit a pull request.

## Acknowledgements

A huge shoutout to [Boot.dev](https://boot.dev/) for breaking down how to make a static site generator on your own.
