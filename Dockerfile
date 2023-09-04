FROM golang:1.20-alpine

# create directory folder
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

# create executable file with name "immersive"
RUN go build -o immersive

# run executable file
CMD ["./immersive"]