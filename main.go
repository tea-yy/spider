package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/tea-yy/spider/model"
	"github.com/tea-yy/spider/util"
)

//获取评论总条数
func getPageCount() (int, error) {
	page := 0
	content, err := util.GetHttpResponse("https://club.jd.com/comment/productCommentSummaries.action?referenceIds=49938064797&callback=jQuery5037676&_=1589162096858")
	if err != nil {
		return page, err
	}

	start := strings.Index(content, "(")
	end := strings.LastIndex(content, ")")
	if start <= 0 || end <= 0 {
		return page, errors.New("response data error")
	}
	content = content[start+1 : end]

	//反序列化json对象
	var result model.CommentsCountObj
	if err := json.Unmarshal([]byte(content), &result); err != nil {
		return page, err
	}

	if len(result.CommentsCount) == 0 {
		return page, errors.New("comments length is 0")
	}

	//计算分页总页数，一页10条
	pageSize := 10
	page = result.CommentsCount[0].Count / pageSize
	if result.CommentsCount[0].Count%pageSize > 0 {
		page = page + 1
	}

	return page, nil
}

//获取评论内容
func getComments(page int) ([]string, error) {
	total := 0
	comments := make([]string, 0)
	for i := 0; i < page; i++ {
		gate := fmt.Sprintf("https://club.jd.com/comment/productPageComments.action?callback=fetchJSON_comment98&productId=49938064797&score=0&sortType=5&page=%d&pageSize=10&isShadowSku=0&fold=1", i)
		fmt.Println(fmt.Sprintf("============== 正在抓取第【%d】页数据，抓取地址：%s ==============", i+1, gate))
		content, err := util.GetHttpResponse(gate)
		if err != nil {
			return comments, err
		}

		start := strings.Index(content, "(")
		end := strings.LastIndex(content, ")")
		if start <= 0 || end <= 0 {
			continue
		}
		content = content[start+1 : end]

		//反序列化json对象
		var result model.CommentObj
		if err := json.Unmarshal([]byte(content), &result); err != nil {
			return comments, err
		}

		if len(result.Comments) == 0 {
			break
		}

		for j, v := range result.Comments {
			total = (i + 1) * (j + 1)
			fmt.Println(fmt.Sprintf("%d:%s", total, v.Content))
			if strings.Index(v.Content, "好评") > 0 {
				comments = append(comments, v.Content)
			}
		}
	}

	msg := fmt.Sprintf("抓取数据完毕，共%d条评论，其中包含关键字【好评】的有%d条\n", total, len(comments))
	fmt.Println(msg)
	return comments, nil
}

//输出满足条件评论内容并保存到spider
func saveComments(comments []string) error {
	file, err := os.OpenFile("spider.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	for i, v := range comments {
		line := fmt.Sprintf("%d:%v\n", i+1, v)
		file.WriteString(line)
		fmt.Println(line)
	}

	return nil
}

func main() {
	fmt.Println("==============spider begin==============")

	//获取评论总条数
	page, err := getPageCount()
	if err != nil {
		fmt.Println("getPageCount fail:", err)
		return
	}

	//获取评论内容
	comments, err := getComments(page)
	if err != nil {
		fmt.Println("getComments fail:", err)
		return
	}

	//输出满足条件评论内容并保存到spider.txt文件
	err = saveComments(comments)
	if err != nil {
		fmt.Println("saveComments fail:", err)
		return
	}

	fmt.Println("==============spider end==============")
}
