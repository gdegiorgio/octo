#!/bin/sh

OCTO_HOME=$HOME/.octo/



install(){


    if [ -f $OCTO_HOME/bin/octo ]; then
        echo "Octo binary already exists"
        return
    fi

    # Check if curl is installed
    if ! command -v curl &> /dev/null; then
        echo "curl could not be found"
        return
    fi

    if [ $(uname -m) == "x86_64" ]; then
        ARCH="amd64"
    elif [ $(uname -m) == "aarch64" ]; then
        ARCH="arm64"
    else
        echo "Unsupported architecture"
        return
    fi
  
    if [ $(uname -s) == "Darwin" ]; then
        OS="darwwin"
    elif [ $(uname -s) == "Linux" ]; then
        OS="linux"
    else
        echo "Unsupported OS"
        return
    fi

    mkdir -p $OCTO_HOME/bin
    curl -L -o $OCTO_HOME/bin/octo_$OS_$ARCH.tar.gz https://github.com/gdegiorgio/octo/releases/latest/download/octo_$OS_$ARCH.tar.gz

    tar -xzf $OCTO_HOME/bin/octo_$OS_$ARCH.tar.gz -C $OCTO_HOME/bin
    rm $OCTO_HOME/bin/octo_$OS_$ARCH.tar.gz 

    chmod +x $OCTO_HOME/bin/octo


    if [ $SHELL == "/bin/bash" ]; then
        echo 'export PATH=$PATH:$OCTO_HOME/bin' >> $HOME/.bashrc
        source $HOME/.bashrc

    elif [ $SHELL == "/bin/zsh" ]; then
        echo 'export PATH=$PATH:$OCTO_HOME/bin' >> $HOME/.zshrc
        source $HOME/.zshrc
    elif [ $SHELL == "/bin/sh" ]; then
        echo 'export PATH=$PATH:$OCTO_HOME/bin' >> $HOME/.profile
        source $HOME/.profile
    else
        echo "Unsupported shell"
        return
    fi


    echo "Octo installed successfully"

}



install
