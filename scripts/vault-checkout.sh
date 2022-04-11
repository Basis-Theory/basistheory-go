#!/bin/bash
set -e

current_directory="$PWD"

cd $(dirname $0)

cd ../vault-api

if [ ! -f "$(pwd)/docker-compose.yml" ]
then
  cd ../vault-api
  git init .
  mv .git .gitvault
  git --git-dir=.gitvault config --global init.defaultBranch master
  git --git-dir=.gitvault sparse-checkout init
  git --git-dir=.gitvault sparse-checkout set "db/" "local-certs/" "docker-compose.yml"
  if [ "$IS_PR_WORKFLOW" != true ]
  then
    git --git-dir=.gitvault remote add -f vault git@github.com:Basis-Theory/basistheory-vault-api.git
    git --git-dir=.gitvault pull vault master
    git --git-dir=.gitvault remote rm vault
    yes | rm -r .gitvault
  else
    git --git-dir=.gitvault remote add -f vault git@github.com:Basis-Theory/basistheory-vault-api.git >/dev/null 2>&1
    git --git-dir=.gitvault pull vault master >/dev/null 2>&1
    git --git-dir=.gitvault remote rm vault
    yes >/dev/null 2>&1 | rm -r .gitvault >/dev/null 2>&1
  fi
  awk 'NR > 1 && !(/context: \./ && p ~ /build/) { print p } { p = $0 } END { print }' docker-compose.yml > tmp && mv tmp docker-compose.yml
  if [ "$(uname)" == "Darwin" ]; then
    sed -i '' 's|vault-api:latest|ghcr.io/basis-theory/vault-api:latest|g' docker-compose.yml
    sed -i '' '/context: \./d' docker-compose.yml
    sed -i '' '/args:/d' docker-compose.yml
    sed -i '' '/- GIT_SHA/d' docker-compose.yml
    sed -i '' '/- GITHUB_TOKEN/d' docker-compose.yml
    sed -i '' '/9091:443/ { n; n; s/$/\n      - $PWD\/wiremock:\/app\/__admin/; }' docker-compose.yml
  else
    sed -i 's|vault-api:latest|ghcr.io/basis-theory/vault-api:latest|g' docker-compose.yml
    sed -i '/context: \./d' docker-compose.yml
    sed -i '/args:/d' docker-compose.yml
    sed -i '/- GIT_SHA/d' docker-compose.yml
    sed -i '/- GITHUB_TOKEN/d' docker-compose.yml
    sed -i '/9091:443/ { n; n; s/$/\n      - $PWD\/wiremock:\/app\/__admin/; }' docker-compose.yml
  fi
else
  echo "Vault docker-compose.yml found"
fi

result=$?

cd "$current_directory"

exit $result
