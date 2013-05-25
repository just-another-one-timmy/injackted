injackted
=========
Ever dreamed to perform fast search queries on your local files? Injackted is what you're looking for! It's kind of like Google, but not as cool and doesn't have adwords.


How to run
==========
To clone the repo and compile crawler and client:
```
git clone https://github.com/just-another-one-timmy/injackted
cd injackted
export GOPATH=$(pwd)
go build src/crawler.go
go build src/sampleclient.go
```

To build an index:
```
# Assuming you are in the 'injackted' dir:
find `pwd` -type f -name "*.go" | ./crawler < -o example-index
```

To run samle client:
```
echo "load example-index list-all bye-bye" | ./sampleclient
```
