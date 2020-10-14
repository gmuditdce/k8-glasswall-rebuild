FROM python:3-alpine
WORKDIR /usr/src/app
COPY requirements.txt .
RUN pip install -r requirements.txt
VOLUME [ "/input" ]
ENV ENDPOINT=play.min.io \
    ACCESS_KEY=Q3AM3UQ867SPQQA43P2F \
    SECRET_KEY=zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG \
    BUCKET=test-bucket \
    WORKERS=50
COPY . .
CMD [ "python", "upload-directory-to-minio" ]