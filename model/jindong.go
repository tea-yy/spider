package model

//评论总数结构体
type CommentsCountObj struct {
	CommentsCount []CommentsCount `json:"CommentsCount"`
}

type CommentsCount struct {
	Count     int `json:"CommentCount"`
	ShowCount int `json:"ShowCount"`
	SkuId     int `json:"SkuId"`
	ProductId int `json:"ProductId"`
}

//评论内容结构体
type CommentObj struct {
	Comments []Comments `json:"comments"`
	MaxPage  int        `json:"maxPage"`
	Score    int        `json:"score"`
}

type Comments struct {
	Id       int    `json:"id"`
	Guid     string `json:"guid"`
	Nickname string `json:"nickname"`
	Content  string `json:"content"`
}
