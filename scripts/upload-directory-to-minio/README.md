# Upload Directory to MinIO

This container uploads and deletes all files from its `/input` directory to an S3 bucket.

## Configuration

Several environment variables may be set for configuration.

| Environment Variable | Description                          | Default                                    |
| -------------------- | ------------------------------------ | ------------------------------------------ |
| `ENDPOINT`           | S3 endpoint                          | `play.min.io`                              |
| `ACCESS_KEY`         | S3 access key                        | `Q3AM3UQ867SPQQA43P2F`                     |
| `SECRET_KEY`         | S3 secret key                        | `zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG` |
| `BUCKET`             | S3 bucket name                       | `test-bucket`                              |
| `WORKERS`            | Amount of threads in the thread pool | `50`                                       |
