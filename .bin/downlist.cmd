@echo off

rem Download a file listed in a list file
rem @author Dumitru Uzun (DUzun.Me)

set url=%1
set dest=%2

if [ -z "$url" ]; then
    echo "$(basename "$0") <url> [<filename>]"
    exit 1;
fi;

if [ -z "$dest" ]; then
    dest="$pwd/$(basename "$url")";
fi;

base=$(dirname "$url")

echo -n > "$dest"
echo Last segment: "$(curl -s "$url" | grep -v "^#" | tail -1)"
curl -s "$url" | grep -v "^#" | while read -r i; do echo "$i"; curl -# "$base/$i" >> "$dest"; done;
