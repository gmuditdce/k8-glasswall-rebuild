FROM python:3-alpine
WORKDIR /usr/src/app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
ENV ELASTICSEARCH_HOST="elasticsearch" \
    ELASTICSEARCH_PORT="9200" \
    ELASTICSEARCH_INDEX="reports"
CMD [ "python", "uploader" ]