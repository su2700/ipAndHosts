#!/bin/zsh

# Check if the script is running on macOS
if [[ "$(uname)" != "Darwin" ]]; then
    echo "This script is designed to run on macOS."
    exit 1
fi

# Prompt the user for the app path
echo -n "Enter the path to the app you want to add to PATH: "
read app_path

# Check if the app path is valid
if [[ ! -e "$app_path" ]]; then
    echo "Invalid app path: $app_path"
    exit 1
fi

# Get the Zsh configuration file path
zshrc_file="$HOME/.zshrc"

# Check if the app path is already in the PATH
if grep -q "$app_path" "$zshrc_file"; then
    echo "App already in PATH."
    exit 0
fi

# Append the app path to the PATH in the Zsh configuration file
echo "export PATH=\"$app_path:\$PATH\"" >> "$zshrc_file"
echo "App added to PATH successfully."

# Source the updated Zsh configuration file
source "$zshrc_file"