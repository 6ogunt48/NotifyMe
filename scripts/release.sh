#!/bin/bash
# This is a helper script to tag and push a new release.

latest_tag=$(git describe --tags "$(git rev-list --tags --max-count=1)")

if [[ -z "$latest_tag" ]]; then
    echo -e "No tags found (yet) - Continue to create and push your first tag"
    latest_tag="[unknown]"
fi

echo -e "The latest release tag is: ${latest_tag}"

read -r -p 'Enter a new release tag (vX.X.X format): ' new_tag

tag_regex='v[0-9]+\.[0-9]+\.[0-9]+$'
if echo "$new_tag" | grep -q -E "$tag_regex"; then
    echo -e "Tag: ${new_tag} is valid"
else
    echo -e "Tag: $new_tag is not valid (must be in vX.X.X format)"
    exit 1
fi

read -r -p 'Have you updated the CHANGELOG.md file? (y/n): ' changelog_updated

if [[ "$changelog_updated" != "y" ]]; then
    echo -e "Please update the CHANGELOG.md file before proceeding."
    exit 1
fi

git tag -a "$new_tag" -m "$new_tag Release"
echo -e "Tagged: $new_tag"

git push --tags
echo -e "Release tag pushed to remote"
echo -e "Done!"