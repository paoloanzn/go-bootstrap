# go-bootstrap

A tool to bootstrap Go projects with a predefined structure based on a JSON template.

## Overview

go-bootstrap is a command-line utility written in Go that helps developers quickly set up new Go projects by generating a directory structure and files based on a customizable JSON template. It simplifies the process of starting a new project by automating the creation of folders and empty files, with support for placeholders like <main_package> that are replaced with the project name.

## Installation

To install go-bootstrap, follow these steps:

1. Clone the repository:

```sh
git clone https://github.com/paoloanzn/go-bootstrap.git
```

2. Navigate to the project directory:

```sh
cd go-bootstrap
```

3. Build and install the binary:
   Run the following command to compile the tool and install it to /usr/local/bin:

```sh
make install
```

Alternatively, you can build it manually:

```sh
go build -o go-bootstrap cmd/go-bootstrap/main.go
```

Then, move the go-bootstrap binary to a directory in your PATH, such as:

```sh
mv go-bootstrap /usr/local/bin/
```

### Requirements:

- Go 1.24.1 or later (as specified in go.mod)

- Git (for cloning the repository)

- Make (optional, for using the Makefile)

## Usage

The primary command for go-bootstrap is init, which creates a new project based on a specified JSON template.
To initialize a new project, run:

```sh
go-bootstrap init <path-to-template>
```

### Example

Using the provided sample template (templates/base.json):

```sh
go-bootstrap init templates/base.json
```

This command will:

- Read the base.json template.

- Create a directory named default-go-project (as specified in the template's config.name).

- Generate the following structure with empty files:

```
default-go-project/
├── cmd/
│   └── default-go-project/
│       └── main.go
├── config/
│   └── config.go
├── LICENSE
├── Makefile
└── README.md
```

The <main_package> placeholder in the template is replaced with the project name (default-go-project in this case).

## Custom Templates

You can create your own JSON template to define a custom project structure. The template consists of two main sections: project and config.

### Template Structure

Here’s an example of a custom template:

```json
{
  "project": {
    "src": {
      "main.go": "file"
    },
    "docs": {
      "README.md": "file"
    }
  },
  "config": {
    "name": "my-custom-project"
  }
}
```

- project: A nested object representing the directory structure. Use "file" as the value to indicate a file should be created.

- config: A map containing configuration options, including:
  - "name": The name of the project directory (required).

#### Creating new templates

You can use the JSON schema in `templates/template.schema.json` to create new templates using LLM AI models such as gpt-4o, Claude, Deepseek and more.

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "go-bootstrap Template Schema",
  "type": "object",
  "properties": {
    "project": {
      "$ref": "#/definitions/node"
    },
    "config": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      },
      "required": ["name"],
      "additionalProperties": true
    }
  },
  "required": ["project", "config"],
  "additionalProperties": false,
  "definitions": {
    "node": {
      "oneOf": [
        {
          "type": "string",
          "enum": ["file"]
        },
        {
          "type": "object",
          "patternProperties": {
            ".*": { "$ref": "#/definitions/node" }
          },
          "additionalProperties": false
        }
      ]
    }
  }
}
```

### Placeholders

You can use wildcards like <main_package> in directory or file names, which will be replaced with the project name from config.name.

### Running with a Custom Template

Save your template (e.g., as my-template.json), then run:

```sh
go-bootstrap init my-template.json
```

This will create a my-custom-project directory with src/main.go and docs/README.md.

## Development

If you’d like to contribute to go-bootstrap, the included Makefile provides several useful targets:

- make build: Build the binary into the build/ directory.
- make clean: Remove build artifacts.
- make test: Run tests (add test files to enable this).
- make run ARGS="init templates/base.json": Build and run with specified arguments.
- make install: Install the binary to /usr/local/bin.
- make cross-build: Build for Linux, macOS, and Windows (amd64).
- make fmt: Format the code using go fmt.
- make vet: Check for potential issues with go vet.
- make dev: Run fmt, vet, and build in sequence.

## License

This project is licensed under the GNU General Public License v3.0. See the LICENSE file for details.

## Contributing

Contributions are welcome! To contribute:
Fork the repository on GitHub.

Make your changes in a feature branch.

Submit a pull request to the main repository: https://github.com/paoloanzn/go-bootstrap.

For bug reports or feature requests, please open an issue on the GitHub repository.

## Version

Current version: 0.1 (defined in config/config.go).
