#!/bin/sh
for f in part*; do
    d="$(head -1 "$f" | awk '{print $2}').txt"
    if [ ! -f "$d" ]; then
        mv "$f" "$d"
    else
        echo "File '$d' already exists! Skiped '$f'"
    fi
done

