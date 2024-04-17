# XTimer

## Hello Xtimer
```shell
curl 'http://127.0.0.1:8000/ping/kratos'
```

## wrk
```shell
git clone https://github.com/wg/wrk
make
sudo cp wrk /usr/local/bin

wrk -v
wrk -t8 -c200 -d10s --latency  "http://www.baidu.com"
```