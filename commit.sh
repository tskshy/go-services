#!/usr/bin/env sh
# quick commit

INFO=$1

if [ ! -n "$INFO" ]; then
	echo "you can commit with your update info, e.g.: ./commit.sh 'UPDATE INFO'"
	INFO="update: `date '+%Y-%m-%d %H:%M:%S'`"
fi

git pull
git add -A
git commit -m "$INFO"
git push
