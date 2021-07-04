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

func crawlDB(t *testing.T) {
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
	}

	crawlColumns(t, dbColumnNames[:])
}

func crawlTest(t *testing.T) {
	//数据库专题
	dbColumnNames := [...]string{"likai"} // 分布式与存储技术-郁白

	crawlColumns(t, dbColumnNames[:])
}

// 默认测试
func TestCrawl(t *testing.T) {
	crawlDB(t)
}

// 快速测试
func TestCrawl2(t *testing.T) {
	crawlTest(t)
}
