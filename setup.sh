#!/bin/bash

# Create directories
mkdir -p resources/views/welcome
mkdir -p config

# Create the index.html template
cat <<EOL > resources/views/welcome/index.html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome to Icepeak</title>
</head>
<body>
    <h1>Hello, Icepeak Framework!</h1>
    <p>This is a dynamic template served from the view.</p>
</body>
</html>
EOL

# Create the view.yaml configuration file
cat <<EOL > config/view.yaml
VIEW_ROOT: "./resources/views/"
EOL

echo "Setup completed: Directories and files created."
