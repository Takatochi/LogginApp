FROM ubuntu:latest

RUN mkdir -p /certs

RUN apt update && apt install -y openssl

RUN openssl req -x509 -newkey rsa:4096 -sha256 -days 365 -nodes -keyout /certs/server.key -out /certs/server.crt

CMD ["sh", "-c", "echo 'Сертифікати створено в /certs. Скопіюйте їх на хост за допомогою docker cp або volume mount.' && ls -l /certs"]


