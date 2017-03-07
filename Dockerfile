FROM golang:latest
RUN go get github.com/PuerkitoBio/goquery
RUN go get github.com/allengeer/scrape
ENTRYPOINT /go/bin/scrape
EXPOSE 5000