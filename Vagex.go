package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	var url, sid, length = "", "", "0"
	for true {
		url, sid, length = dofupdate(url, sid)
		//		fmt.Println(url)
		//		fmt.Println(sid)
		var lengthInt, _ = strconv.Atoi(length)
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "暂停" + strconv.Itoa(lengthInt+5) + "秒")
		time.Sleep(time.Second*time.Duration(lengthInt) + 5)
	}

}

func dofupdate(url, sid string) (string, string, string) {
	email := "814495656@qq.com"
	logged_userid := "400913"
	version := "2.3.5"
	logged_secid := "ENZHIbHqto"
	model := "Firefox"
	dataToSend := ""
	dataToSend = "userid=" + logged_userid + "&version=" + version + "&videoid=" + url + "&model=" + model + "&sid=" + sid + "&secid=" + logged_secid + "&speed=" + "1804.54" + "&youtubeUserId=" + "" + "&youtubeChannelId=" + "" + "&youTubeVideoDuration=" + "0" + "&cid=" + "undefined" + "&subed=" + "false"
	return httpDo(email, dataToSend)
}

func httpDo(userId, data string) (string, string, string) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://vagex.com/fupdater2.php", strings.NewReader("userid="+userId+"&data="+base64.StdEncoding.EncodeToString([]byte(data))))
	if err != nil {
		// handle error
	}
	//	req.PostForm()["userid"] = userId
	//	req.PostForm()["data"] = base64.StdEncoding.EncodeToString(data)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("error")
		return "", "", "0"
	}
	response := string(body)
	var url = string([]rune(response)[strings.Index(response, "<url>")+5 : strings.Index(response, "</url>")])
	var sid = string([]rune(response)[strings.Index(response, "<sid>")+5 : strings.Index(response, "</sid>")])
	var credits = string([]rune(response)[strings.Index(response, "<credits>")+9 : strings.Index(response, "</credits>")])
	var length = string([]rune(response)[strings.Index(response, "<length>")+8 : strings.Index(response, "</length>")])
	var error = string([]rune(response)[strings.Index(response, "<error>")+7 : strings.Index(response, "</error>")])
	if error != "" {
		fmt.Println(error)
	} else {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "赚到点数：" + credits)
	}
	fmt.Println(string(body))
	return url, sid, length
}
