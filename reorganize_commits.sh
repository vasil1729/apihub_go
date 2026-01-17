#!/bin/bash
set -e

cd /home/ultimatum/projects_experiments/free_api/apihub_go

echo "Creating scripts/commits directory..."
mkdir -p scripts/commits

echo "Making new commit scripts executable..."
chmod +x scripts/commits/*.sh

echo "Removing old commit scripts from root..."
rm -f initial_commit.sh commit_randomusers.sh commit_randomjokes.sh

echo "Adding changes to git..."
git add .

echo "Committing reorganization..."
git commit -m "chore: reorganize commit scripts into scripts/commits with serialized naming

- Move commit scripts from root to scripts/commits/
- Rename with serialized format: 001_initial_setup.sh, 002_random_users_api.sh, 003_random_jokes_api.sh
- Each script now shows commit identifier with hash
- Keeps root directory clean"

echo "Reorganization complete!"
git log --oneline -1
