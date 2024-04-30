# Demo for File Object Upload/Download via GCS signed URL

## Before Start
1. create a GCS bucket for this demo
2. create a GAE service
3. setup CORS to GCS bucket, you HAVE TO edit <YOUR_GAE_URL> on `cors.txt`
4. create a service account with role `Service Account Token Creator` on GCP project and role `Storage Object Admin` with the GCS bucket for this demo

## Run on Local
### Backend
```sh
PORT=8082 \
ORIGIN_ALLOWED="http://localhost:5173" \
BUCKET_NAME="<YOUR_BUCKET_NAME>" \
go run ./main.go
```
- `PORT` is backend port which is different with frontend service (:5173).
- `ORIGIN_ALLOWED` is local host frontend service URL.
- `BUCKET_NAME` is GCE bucket name.

### Frontend
#### Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

#### Customize configuration

See [Vite Configuration Reference](https://vitejs.dev/config/).

#### Start
```sh
cd ./signurl-frontend
```

1. install js package
```sh
yarn
```

2. refine js format
```sh
yarn format
```

3. run server with dev mode
```sh
yarn dev
```

4. lint with [ESLint](https://eslint.org/)
```sh
yarn lint
```

## Deploy on GAE
1. edit the app.yaml

2. build frontend
```sh
cd ./signurl-frontend
```

```sh
yarn build
```

3. deploy on GAE
```sh
cd ..
```

```sh
gcloud app deploy app.yaml
```
