#!bin/bash

yay -Sy

bash scripts/yay/chrome-install.sh
bash scripts/yay/heroku-cli-install.sh
bash scripts/yay/teamviewer-install.sh
bash scripts/yay/zoom-install.sh

# optional (default installed by flatpak)
# bash scripts/yay/slack-desktop-install.sh
# bash scripts/yay/spotify-install.sh
# bash scripts/yay/visual-studio-code-install.sh