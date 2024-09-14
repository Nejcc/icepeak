#!/bin/bash

# Create directories
mkdir -p resources/views/welcome
mkdir -p resources/views/errors
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

# Create the 404.html template
cat <<EOL > resources/views/errors/404.html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>404 Not Found</title>
</head>
<body>
    <h1>404 - Page Not Found</h1>
    <p>The page you are looking for does not exist.</p>
</body>
</html>
EOL

# Create the 500.html template
cat <<EOL > resources/views/errors/500.html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>500 Internal Server Error</title>
</head>
<body>
    <h1>500 - Internal Server Error</h1>
    <p>Something went wrong on our end. Please try again later.</p>
</body>
</html>
EOL

# Create the 403.html template
cat <<EOL > resources/views/errors/403.html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>403 Forbidden</title>
</head>
<body>
    <h1>403 - Forbidden</h1>
    <p>You do not have permission to access this page.</p>
</body>
</html>
EOL

# Create the under_construction.html template
cat <<EOL > resources/views/errors/under_construction.html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Under Construction</title>
</head>
<body>
    <h1>Page Under Construction</h1>
    <p>This page is currently under construction. Please check back later.</p>
</body>
</html>
EOL

# Create the debug.html template
cat <<EOL > resources/views/errors/debug.html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Error Debug Information</title>
</head>
<body>
    <h1>Error Debug Information</h1>
    <pre>{{ . }}</pre>
</body>
</html>
EOL

# Create the view.yaml configuration file
cat <<EOL > config/view.yaml
VIEW_ROOT: "./resources/views/"
EOL

echo "Setup completed: Directories and files created."
