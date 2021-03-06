package controllers

import (
	"beego_action/helpers"
	"beego_action/models"
	"encoding/json"
	"fmt"
)

type JsonTestController struct {
	BaseController
}

type msg struct {
	Id    int
	Name  string
	Likes []string
}

//各种json数据测试
func (j *JsonTestController) Get() {
	TestSlice := []int{1, 2, 4, 5, 6}
	test_json, _ := json.Marshal(TestSlice)
	fmt.Println(string(test_json))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	msg_info := msg{Id: 1, Name: "aa", Likes: []string{"appale", "banana", "orange"}}
	msg_infos := make([]msg, 10, 10)
	for i := 0; i < 10; i++ {
		msg_infos = append(msg_infos, msg_info)
	}
	msg_info_json, _ := json.Marshal(msg_infos)
	fmt.Println(string(msg_info_json))
}

func (j *JsonTestController) ClawUrlMsg() {
	res := helpers.My_http_get("http://www.baidu.com")
	j.Ctx.WriteString(res)
}

func (j *JsonTestController) ClawResponseHeader() {
	headers := helpers.ClawResponseHeader("http://www.cnblogs.com")
	json_byte, _ := json.Marshal(headers)
	fmt.Println(string(json_byte))
}

func (g *JsonTestController) GetGameJson() {
	GameModel := &models.GameModel{}
	GameInfoJson := GameModel.GetGameJson()
	g.Data["json"] = GameInfoJson
	g.ServeJSON()
}

func (this *JsonTestController) TestXml() {
	types := this.GetString("type")
	msgData := msg{Id: 11, Name: "aa", Likes: []string{"apple", "orange", "banala"}}
	if types == "json" {
		this.Data["json"] = &msgData
		this.ServeJSON()
	} else {
		this.Data["xml"] = &msgData
		this.ServeXML()
	}
}
