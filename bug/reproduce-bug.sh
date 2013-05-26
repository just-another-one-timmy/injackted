# Create files
declare -r fnum=1021
for i in $(seq 1 $fnum)
do
    echo "keyword$i" > "file$i"
done

# Compile crawler and sampleclient.
cd ..
export GOPATH=`pwd`
cd bug
go build ../src/crawler.go
go build ../src/sampleclient.go
output=$(find `pwd` -name "file*" | \
    ./crawler -o testingindex > /dev/null \
    && echo "load testingindex list-all-docs bye-bye" | \
    ./sampleclient | tail -n 3 | head -n 1 | tr " " "\n" | \
    head -n 1)
if [ "$fnum" != "$output" ];
then
    echo "Bug is here. :'("
    echo "expected $expected, but got $output"
else
    echo "Yay, bug is not here!"
fi
rm file*
rm crawler
rm sampleclient
rm testingindex
