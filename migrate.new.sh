#!/usr/bin/env bash
set -o nounset
set -o errexit

if ! [ $# -eq 1 ]; then
  echo "should provide title"
  exit 1
fi

path="./migrations"

count=$(ls -1 $path | wc -l)
count=$((count/2+1))

filename=$(printf "%06d_$1" $count)
down=$(printf "$path/$filename.down.sql")
up=$(printf "$path/$filename.up.sql")

cat > $down << EOM
begin;

-- Put your alter inside a transaction.

commit;
EOM

cp $down $up

git add $down
git add $up

echo "new migration $filename"
