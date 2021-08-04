package ZhihuZhuanlanCrawler

import (
	"fmt"
	"testing"
)

const columnName = "newsql" // https://zhuanlan.zhihu.com/Otalk
// const columnName = "OTalk" // https://zhuanlan.zhihu.com/Otalk
const pid = 60968502 // https://zhuanlan.zhihu.com/p/60968502

func TestClient_GetPinnedArticlePidAndAuthor(t *testing.T) {
	info, err := GetPinnedArticlePidAndAuthor(columnName)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("%+v\n", *info)
}

func TestClient_GetArticlesListPids(t *testing.T) {
	pids, err := GetArticlesListPids(columnName)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("%v\n", pids)
}

func TestClient_GetSingleArticle(t *testing.T) {
	article, err := GetSingleArticle(pid)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("%+v\n", *article)
}

func crawlColumns(t *testing.T, columns []string) {

	for _, column := range columns {

		articles, err := GetArticlesListHyperlinks(column)
		if err != nil {
			t.Error(err)
			continue
		}
		for _, article := range articles {
			fmt.Println(article)
		}
	}

}

func TestCrawlDB(t *testing.T) {
	//数据库专题
	dbColumnNames := [...]string{"newsql", // tidb后花园
		"c_1166294802416611328", //分布式数据库
		"db-readings",           // 数据系统论文阅读小组
		"c_212000558",           // 一个书魔程序员的读书简评
		"c_158208519",           //用谁都能看懂的方法解释分布式系统
		"c_85050070",            //分布式NoSQL数据库
		"c_190483659",           //数据库前沿技术
		"little-ds",             //分布式系统思考和相关论文
		"distributed-storage",   //分布式和存储的那些事
		"nosql",                 //NoSQL技术剖析
		"b-tree",                //分布式笔记
		"likai",                 // 分布式与存储技术-郁白
		"360infra",              //MySQL内核揭秘
		"c_131009364",           // Distributed Things-网易存储团队
		"c_1053716921761026048", //学习前沿分布式系统的一致性设计
		"c_1036557865221173248", //分布式存储的七方面问题
		"tsangpo",               //tsangpo's note
		"c_1266892102010945536", //跟工作相关的正经事儿
		"c_126076715",           //分布式数据库-路云飞
		"c_1389568413359099904", //数据库存算分离
		"c_209506977",           //硬核技术
		"paxos",                 //Paxos、Raft分布式一致性最佳实践
		"io-meter",              //分布式数据系统小菜
		"c_1238468913098731520", //System全家桶
	}

	crawlColumns(t, dbColumnNames[:])
}

func TestCrawlBigData(t *testing.T) {
	// 大数据专栏
	bigDataColumnNames := [...]string{"c_1225812232481046528", //BigData魔法盒
		"meituantech",           //美团技术博客,主题杂
		"c_128579185",           //阿里大数据玩家
		"c_1051807253732175872", //高可用架构
		"codehole",              //码洞
		"deepinsight",           //洞见实验室
		"zuoqin",                //分布式存储系统和性能优化
		"c_1093515476512169984", //DevTalk
		"fix-bug",               //挖坑/填坑笔记
		"c_1364269765981237248", //Clickhouse源码阅读笔记
		"c_1083667247565705216", //尬聊数据库 tispark
		"c_1267062136230309888", //架构587
		"newbigdata",            //大数据分析与管理
		"c_1175452658096476160", //分步试存储
		"Elasticsearch",         //Elasticsearch技术研讨
		"c_1242102119706431488", // 数据库干货汇

	}

	crawlColumns(t, bigDataColumnNames[:])
}

func TestCrawlCloud(t *testing.T) {
	// 云服务专栏
	cloudColumnNames := [...]string{"c_1010851081424879616", //阿里巴巴云原生
		"TencentCloudCommunity", //腾讯云+社区
		"c_148580541",           //进击的云计算
		"tencent-TEG",           // 腾讯技术, 杂主题
		"c_1321864303444451328", //腾讯云原生
		"aliyunedu",             // 阿里云大学
		"cactus",                //云计算会议论文快报
		"c_1040269471008538624", //白话云计算
		"marketplace",           //企业上云那些事
	}

	crawlColumns(t, cloudColumnNames[:])
}

func TestCrawProgramLang(t *testing.T) {
	// 编程语言
	programLangsCN := [...]string{"c_1078248076300521472", //做一枚爱生活的Rustacean
		"c_1139487758685900800", // CPP工程师的Rust迁移之路
		"rust-lang",             //Rust编程
		"c_1186237256184029184", //Rust碎碎念
	}

	crawlColumns(t, programLangsCN[:])

}

func CrawlTest(t *testing.T) {
	//数据库专题
	dbColumnNames := [...]string{"likai"} // 分布式与存储技术-郁白

	crawlColumns(t, dbColumnNames[:])
}

// 默认测试
func TestCrawl(t *testing.T) {
	TestCrawlDB(t)
}

// 快速测试
func TestCrawl2(t *testing.T) {
	CrawlTest(t)
}
