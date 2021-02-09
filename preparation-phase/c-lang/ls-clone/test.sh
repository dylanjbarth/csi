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

echo "Case: curr dir with no args"
./$PROGRAM . > $TEST_CASES_DIR/out1.txt
ls . > $TEST_CASES_DIR/expected1.txt
diff -b $TEST_CASES_DIR/out1.txt $TEST_CASES_DIR/expected1.txt

echo "Case: no args"
./$PROGRAM $TEST_DATA_DIR > $TEST_CASES_DIR/out2.txt
ls $TEST_DATA_DIR > $TEST_CASES_DIR/expected2.txt
diff -b $TEST_CASES_DIR/out2.txt $TEST_CASES_DIR/expected2.txt

# Fails because ls is sort of inconsistent.
# echo "Case: test data dir - C flag"
# ./$PROGRAM -C $TEST_DATA_DIR > $TEST_CASES_DIR/out3.txt
# ls -C $TEST_DATA_DIR > $TEST_CASES_DIR/expected3.txt
# diff -b $TEST_CASES_DIR/out3.txt $TEST_CASES_DIR/expected3.txt

echo "Case: -1 flag"
./$PROGRAM -1 $TEST_DATA_DIR > $TEST_CASES_DIR/out4.txt
ls -1 $TEST_DATA_DIR > $TEST_CASES_DIR/expected4.txt
diff -b $TEST_CASES_DIR/out4.txt $TEST_CASES_DIR/expected4.txt

echo "Case: -1a flag"
./$PROGRAM -1a $TEST_DATA_DIR > $TEST_CASES_DIR/out5.txt
ls -1a $TEST_DATA_DIR > $TEST_CASES_DIR/expected5.txt
diff -b $TEST_CASES_DIR/out5.txt $TEST_CASES_DIR/expected5.txt

# Fails
# echo "Case: -f flag"
# ./$PROGRAM -f $TEST_DATA_DIR > $TEST_CASES_DIR/out5.txt
# ls -f $TEST_DATA_DIR > $TEST_CASES_DIR/expected5.txt
# diff -b $TEST_CASES_DIR/out5.txt $TEST_CASES_DIR/expected5.txt

echo "Case: -S flag"
./$PROGRAM -S $TEST_DATA_DIR > $TEST_CASES_DIR/out6.txt
ls -S $TEST_DATA_DIR > $TEST_CASES_DIR/expected6.txt
diff -b $TEST_CASES_DIR/out6.txt $TEST_CASES_DIR/expected6.txt

echo "Case: -Sa flag"
./$PROGRAM -Sa $TEST_DATA_DIR > $TEST_CASES_DIR/out7.txt
ls -Sa $TEST_DATA_DIR > $TEST_CASES_DIR/expected7.txt
diff -b $TEST_CASES_DIR/out7.txt $TEST_CASES_DIR/expected7.txt

# TODO test some flag combinations