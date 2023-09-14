

***Routes are***
```
    /ping
    /statistic
    /processing
    /config
```

***/statistic***
Request
``` 
curl localhost:8080/statistic\
>  -H "user: user1"\
>  -H "exchange: BINANCE"\
>  -H "pair: BTCUSDT"

```
Response
```
{"ATOMUSDT":{"orders":[{"entry":6.287,"volume":1000,"side":"BUY"},{"entry":6.451,"volume":1000,"side":"SELL"},{"entry":6.359,"volume":0.98,"side":"BUY"}]},"BTC":{"orders":[{"entry":24900,"volume":0.98,"side":"BUY"},{"entry":25810,"volume":0.98,"side":"SELL"},{"entry":25630,"volume":0.98,"side":"BUY"},{"entry":25915,"volume":0.98,"side":"SELL"},{"entry":26000,"volume":0.98,"side":"BUY"}]},"DOTUSDT\u003eATOMUSDT":{"orders":[]},"ETH":{"orders":[{"entry":1434,"volume":2.4,"side":"BUY"},{"entry":1531,"volume":2.4,"side":"SELL"},{"entry":1518,"volume":2.4,"side":"BUY"},{"entry":1581,"volume":2.4,"side":"SELL"}]},"MATICUSDT":{"orders":[{"entry":0.494,"volume":10000,"side":"BUY"},{"entry":0.499,"volume":0.98,"side":"SELL"},{"entry":0.501,"volume":0.98,"side":"BUY"}]}}
```


***/processing***
Request
```
curl localhost:8080/processing \
> -H "user: user1" \
> -H "exchange: BINANCE"
```

Response
```
{"ATOMUSDT":{"status":"IN_STAGE_OF_PLACED_ORDER_TRACKING","configurations":{"algorithms":["algorithm2","algorithm3"]}},"BTCUSDT":{"status":"IN_STAGE_OF_ESTIMATION","configurations":{"algorithms":["algorithm1","algorithm2","algorithm3"]}},"DOTUSDT\u003eATOMUSDT":{"status":"IN_STAGE_OF_PLACED_ORDER_TRACKING","configurations":{"algorithms":["algorithm3","algorithm8"]}},"ETHUSDT":{"status":"IN_STAGE_OF_PLACED_ORDER_TRACKING","configurations":{"algorithms":["algorithm2","algorithm3"]}},"MATICUSDT":{"status":"IN_STAGE_OF_TRADING","configurations":{"algorithms":["algorithm3","algorithm5"]}}}
```
