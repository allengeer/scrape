# Web Content Extractor Microservice

## Summary
This package is intended to stand up a microservice on port 5000 that will accept a url and return the extracted
text content from that URL.

## Installation
To install run the command

    go get github.com/allengeer/scrape

## Usage
To start the microservice simply ensure that $GOPATH/bin is on your $PATH and execute the program

    scrape

Visit [http://localhost:5000/status](http://localhost:5000/status) to see the Microservice is running.

## Example
To get the content of a website simply include the url in the url parameter for example

[http://localhost:5000/scrape?url=http://www.allengeer.com](http://localhost:5000/scrape?url=http://www.allengeer.com)