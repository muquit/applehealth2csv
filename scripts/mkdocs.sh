#!/bin/sh
# muquit@muquit.com Jan-10-2021 
MH="markdown_helper"
RM="/bin/rm -f"
DOC_DIR="./docs"
PROG=./applehealth2csv

if [ -f $PROG ]; then
    echo " - Geneate $DOC_DIR/usage.txt"
    $PROG -h > $DOC_DIR/usage.txt 2>&1
    sed -i.bak 's:default \".*:default "/path/working/directory"):' docs/usage.txt
    rm -f docs/usage.txt.bak
fi
pushd $DOC_DIR >/dev/null 
echo " - Assembling README.md"
${MH} include --pristine main.md ../README.md
${MH} include --pristine chl.md ../ChangeLog.md
popd >/dev/null

