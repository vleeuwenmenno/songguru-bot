#!/bin/bash
BRANCH=$1
COMMIT=$(echo $2 | cut -c1-8)
cd ~/$BRANCH

# Delete the old live folder if it exists
if [ -d "previous_version" ]; then
    echo "Deleting previous_version ..."
    rm -rf previous_version
fi

# Move SQL database if it exists & move the old live to the old folder
if [ -d "live" ]; then
    echo "Copying json files ..."
    cp live/users-blacklist.json build/users-blacklist.json
    cp live/guild-blacklist.json build/guild-blacklist.json
    cp live/options.json build/options.json

    echo "Moving 'live' to 'previous_version' ..."
    mv live previous_version

    if [ -d "previous_version/cache" ]; then
        echo "Moving cache to 'build' ..."
        mv previous_version/cache build/
    fi
fi

# Move new version to live
echo "Moving 'build' to 'live' ..."
mv build/ live
mkdir -p versions/$COMMIT/
cp -r live versions/$COMMIT/

# Create files for version
echo "$BRANCH" > live/BRANCH
echo "$COMMIT" > live/COMMIT

# Restart the service and delete the deployer (this) script
echo "Restarting service ..."
sudo systemctl restart songwhip-bot-$BRANCH.service
rm -rf build/