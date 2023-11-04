package test

import (
	"encoding/json"
	"fmt"
	"github.com/itihey/tikuAdapter/internal/search"
	"github.com/itihey/tikuAdapter/pkg/model"
	"os"
	"testing"
)

func TestSearchIcodefClient_SearchAnswer(t *testing.T) {
	var client = search.IcodefClient{}
	testRequest := model.SearchRequest{
		Question: "下面选项中,属于男性在青春期生理变化的内容有?()",
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		t.Errorf("请求icodef题库异常: %v", err)
		return
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试icodef题库 成功搜题", string(marshal))
}

func TestSearchWannengClient_SearchAnswer(t *testing.T) {
	var client = search.WannengClient{
		Disable: false,
	}
	testRequest := model.SearchRequest{
		Question: "大学生缓解性欲的常见途径是()",
		Options:  []string{"自然缓解", "升华", "转移", "克制"},
		Type:     0,
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		t.Errorf("请求万能题库异常: %v", err)
		return
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试万能题库 成功搜题", string(marshal))
}

func TestSearchEnncyClient_SearchAnswer(t *testing.T) {
	var client = search.EnncyClient{
		Disable: false,
		Token:   os.Getenv("ENNCY_TOKEN"),
	}

	testRequest := model.SearchRequest{
		Question: "《红楼梦》传世抄本中，惟一一个可以确定抄录年代的，原为吴晓玲先生旧藏、今藏于首都图书馆的是？",
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		fmt.Printf("请求enncy题库异常: %v", err)
		return
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试enncy题库 成功搜题", string(marshal))
}

func TestBuguake_SearchAnswer(t *testing.T) {
	var client = search.BuguakeClient{
		Disable: false,
	}

	testRequest := model.SearchRequest{
		Question: "苏轼的《西江月》是其在哪里所创作的?( )",
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		fmt.Printf("请求不挂科题库异常: %v", err)
		return
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试不挂科题库 成功搜题", string(marshal))
}

func TestAidian_SearchAnswer(t *testing.T) {

	var client = search.AidianClient{
		Disable: false,
	}

	testRequest := model.SearchRequest{
		Question: "9.十进制数(9)10比十六进制数(9)16小.( )",
	}

	// 调用被测试的方法
	response, err := client.SearchAnswer(testRequest)

	if err != nil {
		fmt.Printf("请求爱点题库异常: %v", err)
		return
	}
	marshal, _ := json.Marshal(response)

	fmt.Println("测试爱点题库 成功搜题", string(marshal))
}
