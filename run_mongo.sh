docker pull mongo:4.4
mkdir data
docker run --name mongoDB -v $(pwd)/data:/data/db -p 27017:27017 -d --rm mongo:4.4