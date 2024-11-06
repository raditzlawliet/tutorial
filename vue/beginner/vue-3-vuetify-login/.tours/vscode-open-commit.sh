#!/bin/bash

# Ensure a commit hash is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <commit-hash>"
  exit 1
fi

COMMIT_HASH=$1
LOCAL_REPO_PATH="../../../."
TEMP_DIR=$(mktemp -d -t project-$COMMIT_HASH-XXXXXX)

echo $TEMP_DIR

# Clone the local repository to the temporary directory
git clone $LOCAL_REPO_PATH $TEMP_DIR

echo $TEMP_DIR
# Navigate to the temporary directory
cd $TEMP_DIR

# Checkout the specified commit
git checkout $COMMIT_HASH

# additional project path, because old commit are using different folder
PROJECT_PATH="/vue/beginner/vue-3-vuetify-login"

# Open the folder in VS Code
code $TEMP_DIR$PROJECT_PATH
