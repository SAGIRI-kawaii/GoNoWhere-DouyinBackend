// Code generated by goctl. DO NOT EDIT.
package types

type Douyin_favorite_action_request struct {
	Token       string `form:"token"`
	VideoId     int    `form:"video_id"`
	VactionType int    `form:"vaction_type"`
}

type Douyin_favorite_action_response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type Douyin_favorite_list_request struct {
	UserId int    `form:"user_id"`
	Token  string `form:"token"`
}

type Douyin_favorite_list_response struct {
	StatusCode int          `json:"status_code"`
	StatusMsg  string       `json:"status_msg"`
	VideoList  Douyin_video `json:"video_list"`
}

type Douyin_comment_action_request struct {
	Token       string `form:"token"`
	VideoId     int    `form:"video_id"`
	ActionType  int    `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentId   int    `form:"comment_id"`
}

type Douyin_comment_action_response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Comment    string `json:"comment"`
}

type Douyin_comment_list_request struct {
	Token   string `form:"token"`
	VideoId int    `form:"video_id"`
}

type Douyin_comment_list_response struct {
	StatusCode  int               `json:"status_code"`
	StatusMsg   string            `json:"status_msg"`
	CommentList []*Douyin_comment `json:"comment_list"`
}

type Douyin_video struct {
	Id            int         `json:"id"`
	Author        Douyin_user `json:"author"`
	PlayUrl       string      `json:"play_url"`
	CoverUrl      string      `json:"cover_url"`
	FavoriteCount int         `json:"favorite_count"`
	CommentCount  int         `json:"comment_count"`
	IsFavorite    bool        `json:"is_favorite"`
	Titlestring   string      `json:"title"`
}

type Douyin_comment struct {
	Id         int         `json:"id"`
	User       Douyin_user `json:"user"`
	Content    string      `json:"content"`
	CreateDate string      `json:"create_date"`
}

type Douyin_user struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	FollowCount      int    `json:"follow_count"`
	FollowerCount    int    `json:"follower_count"`
	IsFollow         bool   `json:"is_follow"`
	Avatar           string `json:"avatar"`
	Background_image string `json:"background_image"`
	Signature        string `json:"signature"`
	TotalFavorited   int    `json:"total_favorited"`
	WorkCount        int    `json:"work_count = 10"`
	FavoriteCount    int    `json:"favorite_count"`
}