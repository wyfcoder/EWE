package WebTool

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//一些爬虫 工具

//输入网址 输入 post的文本 返回网页的string 代码 输入的是键值编码
func GetWebPost(webside string , text string) (string,error){
	value :=url.Values{}
    post :=postStringToList(text)
	for i:=0 ;i<len(post);i+=2{
		value.Add(post[i],post[i+1])
	}
	resp, err := http.PostForm(webside, value)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	fmt.Println(string(body))

	return string(body),nil

}

//以Get获取网页
func GetWebGet(webside string)(string,error){
	resp, err := http.Get(webside)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}

	return string(body),err
}

//输入一个 string 转化为 一个序列
//解码
//保证 数据Post数据如下格式
/*
*  start=100.
   stop=2000.
   geo_flag=0.
   profile=1
   step=0.5
   sun_n=
 */
func postStringToList(text string) []string{
	post :=[]string{}

	content :=""  //内容
	isStart :=false

	for i:=0 ;i < len(text) ;i++{
		if text[i] == ' ' || text[i]=='\n' || text[i] == '\r' || text[i]=='\t'{
			if isStart{
				post=append(post,content)
				content=""
				isStart=false
			}
			continue
		}
		if text[i] == '=' || text[i] == ':'{
			post=append(post,content)
			isStart=true
			content=""
			continue
		}
		isStart=true
		content+=string(text[i])
	}
	if isStart{
		post=append(post, content)
	}
	return post
}


//标志 和 网页 结尾为 "
func GetWebside(start string , text *string) string{


	webside :=""
	k:=strings.Index(*text,start)
	for i:=k ; i< len(*text) ;i++{
		if  (*text)[i] == '"'{
			break
		}
		webside += string((*text)[i])
	}
     return webside
}