#!bin/bash

if [ ! "`whoami`" = "root" ]
then
    echo "Use sudo to run this script"
    exit 1
fi

pamac install pyenv pyenv-virtualenv --no-confirm
