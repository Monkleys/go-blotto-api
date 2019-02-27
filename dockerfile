FROM golang:1.11.5
RUN mkdir  /blotto
WORKDIR /blotto
COPY . .
RUN go get -u github.com/gorilla/mux


#CMD ["app"]
