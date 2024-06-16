#!/bin/bash

# Enable exit on error
set -e

# Function to convert a string to snake_case
to_snake_case() {
  echo "$1" | awk '{print tolower($0)}' | sed 's/ /_/g'
}

# Prompt for the lesson name
read -p "Enter the lesson name: " lesson_name

# Convert the lesson name to snake_case
lesson=$(to_snake_case "$lesson_name")
lesson_dir=src/$lesson

# Create the lesson directory
mkdir -p $lesson_dir

# Create the implementation file
impl_file="$lesson_dir/${lesson}.go"
cat <<EOL > "$impl_file"
package $lesson

// TODO: Add implementation here
EOL

# Create the test file
test_file="$lesson_dir/${lesson}_test.go"
cat <<EOL > "$test_file"
package $lesson

import "testing"

// TODO: Add tests here
func TestExample(t *testing.T) {
    // Example test
}
EOL

echo "Lesson '$lesson_name' created successfully in directory '$lesson_dir'."
echo "Files created:"
echo "  - $impl_file"
echo "  - $test_file"