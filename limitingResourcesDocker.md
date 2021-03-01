#Limiting resources using Docker

    $ docker run -d --name kuard \
    --publish 8080:8080 \
    --memory 200m \
    --memory-swap 1G \
    --cpu-shares 1024 \
    gcr.io/kuar-demo/kuard-amd64:blue