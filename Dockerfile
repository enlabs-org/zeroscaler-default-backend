FROM golang:1.19 AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o zeroscaler-default-backend


FROM alpine:latest

COPY errors /errors
COPY --from=build /app/zeroscaler-default-backend /zeroscaler-default-backend

EXPOSE 80

# Run
CMD ["/zeroscaler-default-backend"]
