FROM golang:1.17 AS builder
LABEL bayt.cloud.image.authors="laith@bayt.cloud"
ARG GITHUB_ID
ENV GITHUB_ID=$GITHUB_ID
ARG MY_GITHUB_TOKEN
ENV MY_GITHUB_TOKEN=$MY_GITHUB_TOKEN
WORKDIR /app
USER $APP_USER
ADD src .
RUN git config \
  --global \
  url."https://${github_id}:${github_token}@github.com".insteadOf \
  "https://github.com"
ENV GOPRIVATE="github.com/laithrafid"
RUN go mod download
RUN go mod verify
RUN go build -o /oauthapi


FROM alpine:3.15.0 AS runner
ARG CASS_DB_SOURCE
ARG CASS_DB_KEYSPACE
ARG CAS_DB_NODES
ARG OAUTH_API_ADDRESS
ENV OAUTH_API_ADDRESS=$OAUTH_API_ADDRESS
ENV CASS_DB_SOURCE=$CASS_DB_SOURCE
ENV CASS_DB_KEYSPACE=$CASS_DB_KEYSPACE
ENV CAS_DB_NODES=$CAS_DB_NODES
WORKDIR /
COPY --from=builder /oauthapi /oauthapi
EXPOSE $OAUTH_API_ADDRESS
ENTRYPOINT ["/oauthapi"]
