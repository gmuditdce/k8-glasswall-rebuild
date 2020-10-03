from elasticsearch import Elasticsearch
import os
import json
import pathlib

if __name__ == "__main__":
    elasticsearch_host = os.getenv("ELASTICSEARCH_HOST")
    elasticsearch_port = os.getenv("ELASTICSEARCH_PORT")
    elasticsearch_index = os.getenv("ELASTICSEARCH_INDEX")

    es = Elasticsearch(hosts=[{
        "host": elasticsearch_host,
        "port": int(elasticsearch_port)
    }])

    if es.ping():
        print(
            f"Connected to Elasticsearch on {elasticsearch_host}:{elasticsearch_port}")
        for path in pathlib.Path("/input").iterdir():
            if path.is_file():
                with open(path, "r") as file:
                    print(f"Indexing {file.name}")
                    es.index(index=elasticsearch_index, body=json.load(file))
                    print(f"Indexed {file.name}")
    else:
        print(
            f"Couldn't connect to Elasticsearch on {elasticsearch_host}:{elasticsearch_port}")
