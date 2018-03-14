## Prereqs
docker

## How to Run
```
$ docker build -t app .
$ docker run -d -p "8000:8000" --restart always --name app app 
```
