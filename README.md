# fshlint

`fshlint` is a simple CLI tool that lints FSH (FHIR Shorthand) files without compilation, leveraging the ANTLR language 
recognizer. Unlike SUSHI, which compiles FSH files into FHIR artifacts, fshlint checks .fsh files for syntax issues, 
making it ideal for quick validations in CI pipelines or project Makefiles.


## Installation
### Prerequisites
Ensure you have Go 1.18 >= installed on your machine.

### Installing with go install
You can install the binary directly using go install. Follow these steps:

Set the environment variable for the Go path:

```sh
export GOPATH=$(go env GOPATH)
````
Install the binary:

```sh
go install github.com/devpies/fshlint@latest
```
This command will download the package, compile it, and place the binary in your $GOPATH/bin.

Ensure `$GOPATH/bin` is in your $PATH:

```sh
export PATH=$PATH:$GOPATH/bin
```
You can add this line to your shell configuration file (e.g., `~/.bashrc`, `~/.zshrc`) to make it persistent:

```sh
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc
```

### Installing from Source
You can also install `fshlint` by cloning the repository and building it from source. Follow these steps:


```sh
git clone https://github.com/devpies/fshlint
cd fshlint
go build -o fshlint fshlint.go
sudo mv fshlint /usr/local/bin/fshlint
```

## Getting Started
### Linting FSH Files
Run `fshlint` in a directory containing FSH files to check for syntax and semantic errors.


```yaml
$ fshlint path/to/files

path/to/files/PatientProfile.fsh:3:7: extraneous input '0..0l' expecting {<EOF>, KW_ALIAS, KW_PROFILE, KW_EXTENSION, KW_INSTANCE, KW_INVARIANT, KW_VALUESET, KW_CODESYSTEM, KW_RULESET, KW_MAPPING, KW_LOGICAL, KW_RESOURCE}
* name 0..0l
       ^
```

## Help

Display help text.

```sh
$ fshlint -h

Usage: fshlint <path-to-fsh-file-or-directory>

Options:
--help, -h        Show this help message
```

## Contributing
To contribute, please create an issue or pull request. For common development tasks, utilize the project's Makefile.

```sh
make install
make generate
make test
make build
```

## License
This project is licensed under the MIT License. See the [LICENSE file](./LICENSE) for details.
