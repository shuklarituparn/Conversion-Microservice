#!/bin/bash

# Define the output file
output_file="folder_names.txt"

# List all directories and write their names to the output file
ls -d */ > "$output_file"

echo "Folder names have been written to $output_file"

