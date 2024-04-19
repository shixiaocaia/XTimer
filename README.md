# XTimer

## Hello Xtimer
```shell
curl 'http://127.0.0.1:8000/ping/kratos'
```
## Create Timer
```shell
curl --location 'http://127.0.0.1:8000/createTimer' \
--header 'Content-Type: application/json' \
--data '{
    "app": "test-xtimer-go",
    "name":"测试Xtimer-每分钟运行一次",
    "cron":"* * * * *",
    "notifyHTTPParam":{
        "url":"http://127.0.0.1:8000/xtimer/callback",
        "method":"POST",
        "body":" its time on. this is a callback msg"
    }
}'
```

## wrk
```shell
git clone https://github.com/wg/wrk
make
sudo cp wrk /usr/local/bin

wrk -v
wrk -t8 -c200 -d10s --latency  "http://www.baidu.com"
```