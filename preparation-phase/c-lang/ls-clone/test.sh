#!/bin/bash

PROGRAM=ls-clone.out
TEST_DATA_DIR=./test/data
TEST_CASES_DIR=./test/cases

echo "Clearing out test dir if it exists..."
rm -rf ./test

echo "Creating test dir and sample files..."

mkdir -p $TEST_DATA_DIR
mkdir -p $TEST_CASES_DIR

# Create a bunch of files.
for i in {1..500}
do
  echo "Test contents $i" > $TEST_DATA_DIR/testfile-$i.txt;
done
# create a few more with interesting permissions or file lengths.
touch $TEST_DATA_DIR/areallllllllllllllllllllllllllllllylongfilename.txt && chmod 0777 $TEST_DATA_DIR/areallllllllllllllllllllllllllllllylongfilename.txt
touch $TEST_DATA_DIR/global_all.txt && chmod 0777 $TEST_DATA_DIR/global_all.txt
touch $TEST_DATA_DIR/user_read.txt && chmod 0400 $TEST_DATA_DIR/user_read.txt
touch $TEST_DATA_DIR/group_read.txt && chmod 0740 $TEST_DATA_DIR/group_read.txt
echo "A bit biggggggggggggggggger" > $TEST_DATA_DIR/bigfile.txt
echo "A bit smaller" > $TEST_DATA_DIR/smallfile.txt
# test another directory
mkdir $TEST_DATA_DIR/another_dir

echo "Running tests..."

function run_test {
  echo "Test Case $1: $2"
  ./$PROGRAM $2 > "$TEST_CASES_DIR/out$1.txt"
  ls $2 > "$TEST_CASES_DIR/expected$1.txt"
  if diff -b "$TEST_CASES_DIR/out$1.txt" "$TEST_CASES_DIR/expected$1.txt"; then 
  echo "PASS"
  else 
    echo "FAIL. Diff: $(diff -b $TEST_CASES_DIR/out$1.txt $TEST_CASES_DIR/expected$1.txt)"
  fi
}

i=1
for tcase in "." ".." "test/data" "test/data/" "" "-1 $TEST_DATA_DIR" "-1a $TEST_DATA_DIR" "-S $TEST_DATA_DIR" "-Sa $TEST_DATA_DIR"
do
  run_test $i "$tcase"
  ((i++))
done

# TODO test some flag combinations
# -f not working
# -l not implemented
# seems like flags passed without a dir passed also fails

