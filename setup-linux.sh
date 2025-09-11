#!/bin/bash

USER_HOME="$HOME"
UTILS_DIR="$USER_HOME/utils"
mkdir -p "$UTILS_DIR"

chmod +x ./gosearch

mv ./got "$UTILS_DIR"

if [[ ":$PATH:" != *":$UTILS_DIR:"* ]]; then
    echo "export PATH=\"\$PATH:$UTILS_DIR\"" >> "$USER_HOME/.bashrc"
    echo "Path added to .bashrc. Reload shell to use"
else
    echo "utils is already in the PATH."
fi
echo "Setup ended succesfully, type got help for help"