# spider
本代码主要用于抓取京东评论数据，并找出含有指定关键字的内容

### 依赖环境
golang

### 依赖包
```
go get github.com/tea-yy/spider/model
go get github.com/tea-yy/spider/util
```

### 启动方式
```
go run main.go
```

### 实现原理
例如：抓取京东某一个商品的评价 https://item.jd.com/100012720924.html
1.分析网页得到，评论内容通过ajax加载
2.第一步获取评论总条数：https://club.jd.com/comment/productCommentSummaries.action?referenceIds=100012720924&callback=jQuery6089975&_=1589160432285
3.通过评论总条数，计算页数，通过翻页请求获取所有的评论数据：https://club.jd.com/comment/productPageComments.action?callback=fetchJSON_comment98&productId=100012720924&score=0&sortType=5&page=0&pageSize=10&isShadowSku=0&fold=1
4.分析评论数据找出含有【好评】关键字数据，输出到spider.txt文件中