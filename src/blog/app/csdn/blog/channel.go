package blog

import (
	"github.com/astaxie/beego/httplib"
	"blog/app/csdn"
	"time"
	"blog/fox"
)
//获取博客系统分类
type Channel struct {
	AccessToken string `json:"access_token" 是 OAuth授权后获得`
	ClientId    string `json:"client_id" App申请的Key`
}
func (t *Channel)Check() (error) {
	if len(t.AccessToken) < 1 {
		return fox.Error{Msg:"access_token 不能为空"}
	}
	if len(t.ClientId) < 1 {
		return fox.Error{Msg:"client_id 不能为空"}
	}
	return nil
}
//发送
func (t *Channel)Post() (string, error) {
	err := t.Check()
	if err != nil {
		return "", err
	}
	req := httplib.Post(csdn.BLOG_CHANNEL_URL)
	//超时
	req.SetTimeout(100 * time.Second, 30 * time.Second)
	//参数
	req.Param("access_token", t.AccessToken)
	req.Param("client_id", t.ClientId)
	//返回
	s, err := req.String()
	if err != nil {
		return "", err
	}
	return s, nil
}