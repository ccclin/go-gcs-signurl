# Demo for File Object Upload/Download via GCS signed URL

## Before Start
1. create a GCS bucket for this demo
2. create a GAE service
3. setup CORS to GCS bucket, you HAVE TO edit <YOUR_GAE_URL> on `cors.txt`
4. create a service account with roles `Storage Object Creator` and `Storage Object Viewer`
5. create a service account key which type is `P12`
6. change key from `pkcs12` to `pem`
```
$ openssl pkcs12 -in <YOUR_SERVICE_ACCOUNT_P12> -passin pass:notasecret -out <YOUR_OUT_KEY_NAME> -nodes
```

## Run on Local
### Backend
```
$ PORT=8082 ORIGIN_ALLOWED="http://localhost:8080" \
BUCKET_NAME="<YOUR_BUCKET_NAME>" GOOGLE_ACCESS_ID="<YOUR_SERVICE_ACCOUNT_ID>" \
PRIVATE_KEY_PATH="<YOUR_OUT_KEY_NAME>" go run ./main.go
```
- `PORT` is backend port which is different with frontend service (:8080).
- `ORIGIN_ALLOWED` is local host frontend service URL.
- `BUCKET_NAME` is GCE bucket name.
- `GOOGLE_ACCESS_ID` is GCP service account id which is email format.
- `PRIVATE_KEY_PATH` is your service account key which is pem format.

### Frontend
```
$ cd ./signurl-frontend
```

```
$ npm install
```

```
$ npm run serve
```

## Deploy on GAE
1. edit the app.yaml

2. build frontend
```
$ cd ./signurl-frontend
```

```
$ npm run build
```

3. deploy on GAE
```
$ cd ..
```

```
$ gcloud app deploy app.yaml
```
