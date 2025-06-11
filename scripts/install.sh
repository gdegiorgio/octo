
#!/bin/bash
OCTO_HOME=$HOME/.octo




install(){


    if [ -f $OCTO_HOME/bin/octo ]; then
        echo "Octo binary already exists in $OCTO_HOME/bin/octo"
        return
    fi

    
    ARCH=""
    MACHINE="$(uname -m)"

    if [ "$MACHINE" = "x86_64" ]; then
        ARCH="amd64"
    elif [ "$MACHINE" = "aarch64" ] || [ "$MACHINE" = "arm64" ]; then
        ARCH="arm64"
    else
        echo "Unsupported architecture: $MACHINE"
        return 1
    fi

    OS_NAME="$(uname -s)"
    if [ "$OS_NAME" = "Darwin" ]; then
        OS="darwin"
    elif [ "$OS_NAME" = "Linux" ]; then
        OS="linux"
    else
        echo "Unsupported OS: $OS_NAME"
        return 1
    fi

    echo "Installing Octo for $OS $ARCH"

    mkdir -p $OCTO_HOME/bin

    cd $OCTO_HOME

    curl -L -o "octo_${OS}_${ARCH}.tar.gz" "https://github.com/gdegiorgio/octo/releases/latest/download/octo_${OS}_${ARCH}.tar.gz"

    tar -xzf "octo_${OS}_${ARCH}.tar.gz"
    rm "octo_${OS}_${ARCH}.tar.gz"

    mv octo $OCTO_HOME/bin/octo
    chmod +x $OCTO_HOME/bin/octo
    


    if [ $SHELL == "/bin/bash" ]; then
        echo "export OCTO_HOME=$OCTO_HOME" >> $HOME/.bashrc
        echo 'export PATH=$PATH:$OCTO_HOME/bin' >> $HOME/.bashrc    
        source_line="source $HOME/.bashrc"
    elif [ $SHELL == "/bin/zsh" ]; then
        echo "export OCTO_HOME=$OCTO_HOME" >> $HOME/.zshrc
        echo 'export PATH=$PATH:$OCTO_HOME/bin' >> $HOME/.zshrc
        source_line="source $HOME/.zshrc"
    elif [ $SHELL == "/bin/sh" ]; then
        echo "export OCTO_HOME=$OCTO_HOME" >> $HOME/.profile
        echo 'export PATH=$PATH:$OCTO_HOME/bin' >> $HOME/.profile
        source_line="source $HOME/.profile"
    else
        echo "Unsupported shell"
        return
    fi


    echo "Octo installed successfully, restart your shell or run $source_line to use it"

}



install
