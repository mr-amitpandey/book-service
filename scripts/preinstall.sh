#!/bin/bash

install_node() {
    echo "Installing Node.js..."
    
    # Update the package index
    sudo apt update -y
    
    curl -fsSL https://deb.nodesource.com/setup_16.x | sudo -E bash -
    sudo apt install -y nodejs
    
    # Verify the installation
    echo "Node.js version:"
    node -v

    echo "NPM version:"
    npm -v

  # installing swagger cli 
    echo "installing swagger cli"

    npm install -g swagger-cli
}

install_go() {
    echo "Installing Go..."
    
    GO_VERSION="1.21.1"
    GO_URL="https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz"
    
    wget -q ${GO_URL} -O go${GO_VERSION}.linux-amd64.tar.gz
    sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
    
    echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile
    source ~/.profile
    
    echo "Go version:"
    go version
}

if command -v node &>/dev/null; then
    echo "Node.js is already installed, version: $(node -v)"
else
    install_node
fi

if command -v go &>/dev/null; then
    echo "Go is already installed, version: $(go version)"
else
    install_go
fi

echo "Installation complete!"
