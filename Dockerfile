FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY slack/ slack/
COPY k8s/ k8s/
RUN go mod download

COPY main.go ./

RUN go build -o /aub

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /aub /aub

USER nonroot:nonroot

CMD ["/aub"]
