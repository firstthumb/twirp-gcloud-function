# twirp-cloud-function

Google Cloud Function using Twirp

## Just try

On your local machine

```
protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./proto/greeter.proto

go run cmd/main.go

... split terminal ...

curl --request "POST" \
     --location "http://localhost:3000/twirp/com.ekocaman.greeter.Greeter/Greet" \
     --header "Content-Type:application/json" \
     --data '{"name":"John"}' \
     --verbose

> {"greeting":"Hello, John!"}
```

Deploy to Google Cloud

```
gcloud functions deploy Greeter --runtime go111 --trigger-http --region=us-central1

```
