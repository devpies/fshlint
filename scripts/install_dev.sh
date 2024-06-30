#!/bin/bash

set -a; source .env; set +a

# InstallAntlr downloads ANTLR if necessary.
InstallAntlr() {
    if [[ ! -f "$ANTLR_DIR/$ANTLR_JAR" ]]; then
        echo "Downloading ANTLR..."
        mkdir -p "$ANTLR_DIR"
        cd "$ANTLR_DIR" && curl -O "$ANTLR_URL"
        cd - > /dev/null
    fi
}

# InstallJava installs Java on Linux or macOS using the appropriate package manager.
InstallJava() {
  if ! command -v java &> /dev/null; then
    if [[ "$OSTYPE" == "darwin"* ]]; then
      echo "Checking for Homebrew..."
      if ! command -v brew &> /dev/null; then
          echo "Homebrew not found. Installing Homebrew..."
          /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
      fi
      echo "Installing Java using Homebrew..."
      brew install openjdk
      echo 'export PATH="/usr/local/opt/openjdk/bin:$PATH"' >> ~/.zshrc
      echo 'export JAVA_HOME=$(/usr/libexec/java_home)' >> ~/.zshrc
      source ~/.zshrc
    elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
      echo "Updating package list..."
      sudo apt-get update -y
      echo "Installing Java..."
      sudo apt-get install -y openjdk-11-jdk
    fi
  fi
}

# InstallGo installs Go for Linux and macOS operating systems.
InstallGo() {
  if ! command -v go &> /dev/null; then
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
      sudo apt update
      sudo apt install -y golang
    elif [[ "$OSTYPE" == "darwin"* ]]; then
      if ! command -v brew &> /dev/null; then
        echo "Homebrew not found. Installing Homebrew..."
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
        echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> ~/.zprofile
        eval "$(/opt/homebrew/bin/brew shellenv)"
      fi
      brew install go
    fi
  fi
}

# InstallGolangCILint installs Golang CI Linter for Linux and macOS operating systems.
InstallGolangCILint() {
  if ! command -v golangci-lint &> /dev/null; then
    echo "golangci-lint not found. Installing golangci-lint..."
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
      curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
      export PATH=$PATH:$(go env GOPATH)/bin
      echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.profile
    elif [[ "$OSTYPE" == "darwin"* ]]; then
      if ! command -v brew &> /dev/null; then
        echo "Homebrew not found. Installing Homebrew..."
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
        echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> ~/.zprofile
        eval "$(/opt/homebrew/bin/brew shellenv)"
      fi
      brew install golangci-lint
    fi
  fi
}

InstallGo
InstallGolangCILint
InstallJava
InstallAntlr
