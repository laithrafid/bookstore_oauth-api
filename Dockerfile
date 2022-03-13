FROM golang:1.17 AS builder
LABEL bayt.cloud.image.authors="laith@bayt.cloud"
ARG github_id=laithrafid
ENV github_id=$github_id
ARG github_token
ENV github_token=$github_token
ARG API_ADDRESS=:8080
ARG CASS_DB_SOURCE=192.168.0.42:9042
ENV CASS_DB_SOURCE=$CASS_DB_SOURCE
ARG CASS_DB_KEYSPACE=oauth
ENV CASS_DB_KEYSPACE=$CASS_DB_KEYSPACE
ARG CAS_DB_NODES=172.18.9.140:4092,172.18.9.140:4093
ENV CAS_DB_NODES=$CAS_DB_NODES
ENV OAUTH_API_ADDRESS=$API_ADDRESS
ENV USERS_API_ADDRESS=$API_ADDRESS
ENV ITEMS_API_ADDRESS=$API_ADDRESS
ARG ELASTIC_HOSTS="http://192.168.0.42:9200"
ENV ELASTIC_HOSTS=$ELASTIC_HOSTS
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
WORKDIR /
COPY --from=builder /oauthapi /oauthapi
EXPOSE $OAUTH_API_ADDRESS
ENTRYPOINT ["/oauthapi"]
