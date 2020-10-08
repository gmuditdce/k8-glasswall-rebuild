from minio import Minio
from minio.error import (ResponseError, BucketAlreadyOwnedByYou,
                         BucketAlreadyExists)
from pathlib import Path
import os
import concurrent.futures


def upload_file_to_bucket(path, bucket, client):
    try:
        object_name = str(path)[7:]
        client.fput_object(bucket, object_name, path)
        print(f"Uploaded {object_name}")
        Path.unlink(path)
        print(f"Deleted {object_name}")
        return True
    except ResponseError as err:
        print(err)
        return False


if __name__ == "__main__":
    endpoint = os.getenv("ENDPOINT")
    access_key = os.getenv("ACCESS_KEY")
    secret_key = os.getenv("SECRET_KEY")
    bucket = os.getenv("BUCKET")
    workers = os.getenv("WORKERS")

    client = Minio(endpoint=endpoint, access_key=access_key,
                   secret_key=secret_key)

    try:
        client.make_bucket(bucket)
    except BucketAlreadyOwnedByYou as err:
        pass
    except BucketAlreadyExists as err:
        pass
    except ResponseError as err:
        raise

    paths = [p for p in Path("/input").rglob("*") if os.path.isfile(str(p))]

    with concurrent.futures.ThreadPoolExecutor(max_workers=int(workers)) as executor:
        for path in paths:
            c = Minio(endpoint=endpoint, access_key=access_key,
                      secret_key=secret_key)
            executor.submit(upload_file_to_bucket, path, bucket, c)
