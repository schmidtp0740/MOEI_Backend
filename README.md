## Prereqs
docker

## How to Run
```
$ docker build -t goapp .
$ docker run -d -p "8000:8000" --rm --name app goapp 
```