# JDInventoryMonitor

起因是朋友的一个需求，让我帮忙看看有没有办法监控京东某件商品有没有库存。网上也没搜到现成的代码，就自己粗略写了一个简陋版，原理很简单，如下：
1. 先抓包分析哪个接口返回的信息是库存信息，抓到一个接口，响应形如：`{"1566641027":{"freshEdi":null,"sidDely":"-1","channel":1,"rid":"0","sid":"-1","dcId":"-1","IsPurchase":false,"eb":"99","ec":"-1","StockState":34,"ab":"-1","canAddCart":"0","ac":"-1","ad":"-1","ae":"-1","skuState":1,"PopType":0,"af":"-1","ag":"-1","StockStateName":"无货","m":"0","rfg":0,"ArrivalDate":"","v":"0","rn":-1,"dc":"-1"}}`；
2. 再通过前端找到`data-sku`对应的`data-value`(如：data-sku="1566641027" data-value="40张白边相纸")，进行一一对应；
3. 不清楚这种方式获取数据的话，Cookie的有效期能撑多久；所以我把请求包当成了一个json数据放到了外部，每分钟先读取json数据，然后重放请求包，如果有库存则通过企业微信机器人告警，没库存则静默，如果Cookie失效则通过企业微信机器人告警，重新配置参数。

### 效果
<img src="https://github.com/Ed1s0nZ/JDInventoryMonitor/blob/main/%E6%95%88%E6%9E%9C.png" width="500px">   

### 不知道有没有其他现成的工具，或者其他办法。时间原因，我也没详细分析各个参数的用处，如果大家有更好的建议的话，请辛苦告诉我一下。
