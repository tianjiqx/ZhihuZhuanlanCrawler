package ZhihuZhuanlanCrawler

import (
	"fmt"
	"strings"
)

// 获取置顶文章pid以及作者信息
func GetPinnedArticlePidAndAuthor(columnName string) (*PinnedArticleAndAuthor, error) {
	if columnName == "" {
		return nil, ColumnNameCanNotBeEmpty
	}
	u := fmt.Sprintf("https://zhuanlan.zhihu.com/api/columns/%s/pinned-article", columnName)
	res, err := sendNewZhihuRequest(u)
	if err != nil {
		return nil, err
	}

	pinnedArticleAndAuthor := PinnedArticleAndAuthor{}
	err = res.ToJSON(&pinnedArticleAndAuthor)
	if err != nil {
		return nil, err
	}

	return &pinnedArticleAndAuthor, nil
}

// 获取单个文章
func GetSingleArticle(pid int) (*Article, error) {
	if pid == 0 {
		return nil, PidCanNotBeEmpty
	}
	u := fmt.Sprintf("https://api.zhihu.com/articles/%d", pid)
	res, err := sendNewZhihuRequest(u)
	if err != nil {
		return nil, err
	}

	article := Article{}
	err = res.ToJSON(&article)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

// 获取文章pid列表
func GetArticlesListPids(columnName string) ([]int, error) {
	if columnName == "" {
		return nil, ColumnNameCanNotBeEmpty
	}

	var limit = 20
	var offset = 0

	u := fmt.Sprintf("https://zhuanlan.zhihu.com/api/columns/%s/articles?limit=%d&offset=%d", columnName, limit, offset)
	res, err := sendNewZhihuRequest(u)
	if err != nil {
		return nil, err
	}

	articleList := ArticleList{}
	err = res.ToJSON(&articleList)
	if err != nil {
		return nil, err
	}

	var articleIds = []int{}

	for _, entry := range articleList.Data {
		articleIds = append(articleIds, entry.ID)
	}

	for offset = offset + limit; offset < articleList.Paging.Totals; offset = offset + limit {
		u := fmt.Sprintf("https://zhuanlan.zhihu.com/api/columns/%s/articles?limit=%d&offset=%d", columnName, limit, offset)
		res, err := sendNewZhihuRequest(u)
		if err != nil {
			return nil, err
		}

		articleList := ArticleList{}
		err = res.ToJSON(&articleList)
		if err != nil {
			return nil, err
		}
		for _, entry := range articleList.Data {
			articleIds = append(articleIds, entry.ID)
		}
	}

	return articleIds, nil
}

// 获取文章简介信息,标题与url连接
func GetArticlesListHyperlinks(columnName string) ([]string, error) {
	if columnName == "" {
		return nil, ColumnNameCanNotBeEmpty
	}

	var limit = 20
	var offset = 0

	u := fmt.Sprintf("https://zhuanlan.zhihu.com/api/columns/%s/articles?limit=%d&offset=%d", columnName, limit, offset)
	res, err := sendNewZhihuRequest(u)
	if err != nil {
		return nil, err
	}

	articleList := ArticleList{}
	err = res.ToJSON(&articleList)
	if err != nil {
		return nil, err
	}

	var articleHyperlinks = []string{}

	for _, entry := range articleList.Data {
		article, err := GetSingleArticle(entry.ID)
		if err != nil {
			// 忽略错误
			continue
		}
		// 当前URL 格式是https://api.zhihu.com/articles/34311102
		// 需要替换为https://zhuanlan.zhihu.com/p/34311102
		url := strings.Replace(article.URL, "https://api.zhihu.com/articles", "https://zhuanlan.zhihu.com/p", 1)
		hyperLink := fmt.Sprintf(" - [%s](%s) \n", article.Title, url)
		articleHyperlinks = append(articleHyperlinks, hyperLink)
	}

	for offset = offset + limit; offset < articleList.Paging.Totals; offset = offset + limit {
		u := fmt.Sprintf("https://zhuanlan.zhihu.com/api/columns/%s/articles?limit=%d&offset=%d", columnName, limit, offset)
		res, err := sendNewZhihuRequest(u)
		if err != nil {
			return nil, err
		}

		articleList := ArticleList{}
		err = res.ToJSON(&articleList)
		if err != nil {
			return nil, err
		}
		for _, entry := range articleList.Data {
			article, err := GetSingleArticle(entry.ID)
			if err != nil {
				// 忽略错误
				continue
			}
			url := strings.Replace(article.URL, "https://api.zhihu.com/articles", "https://zhuanlan.zhihu.com/p", 1)
			hyperLink := fmt.Sprintf(" - [%s](%s) \n", article.Title, url)
			articleHyperlinks = append(articleHyperlinks, hyperLink)
		}
	}

	return articleHyperlinks, nil
}
