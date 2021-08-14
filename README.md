# ZhihuZhuanlanCrawler
知乎专栏爬虫


基于solarhell/ZhihuZhuanlanCrawler 原始框架基础上，根据自己的需要做了一点改造。

将自己关心的专栏的文章标题连接全部收集，输出为md文件，便于根据主题的进行学习。

PS: 知乎已经取消专栏关注功能，使用爬虫进行手动关注，也是一种解决方式。

## 测试爬虫功能

```shell script

#快速验证爬虫,爬取内容过多，爬去结果在缓冲区，可能一时无法打印结果，不知道脚本是否可用。

go test -run TestCrawl2 

#个人DB专栏爬取, 默认单个测试10m，超时报错
go test -run TestCrawl  -timeout 1h > 2021.07.09.zhihuColumns.database.md

#大数据
go test -run TestCrawlBigData -timeout 20m  > 2021.07.09.zhihuColumns.bigData.md

#云原生
go test -run TestCrawlCloud  -timeout 20m > 2021.07.09.zhihuColumns.cloud.md

# 编程语言
go test -run TestCrawProgramLang -timeout 20m > 2021.07.09.zhihuColumns.programLang.md

```

自定义关注的专栏，可以可以参考 [`crawler_test.go`](crawler_test.go)的`crawlDB`方法，设置专栏名，专栏名可以打开知乎专栏的最后一个名字。
例如，`https://www.zhihu.com/column/c_1051807253732175872` 即`c_1051807253732175872`。

注意使用go 1.12 以上, 个人现在安装了go 1.16

## 更新内容

- fix 无置顶文章下载bug
- 减轻阅读负担，不打印完整内容，只保留标题，后续根据标题查找感兴趣的内容


TODO：

了解知乎API，更简洁的获取内容，话题，关键词，增量更新。



