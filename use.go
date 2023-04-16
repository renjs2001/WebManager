package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/renjs2001/WebManager/client"
	"github.com/renjs2001/WebManager/db"
	"github.com/renjs2001/WebManager/kitex_gen/meta_rpc"
	"time"
)

type effectiveTrafficPiece struct {
	Begin  int32 `json:"begin"`
	Length int32 `json:"length"`
	Flight int64 `json:"flight"`
}

type TrafficMapJson struct {
	Flights []int64                  `json:"flights"`
	Pieces  []*effectiveTrafficPiece `json:"pieces"`
}

type RuleItem struct {
	Instruments map[string][]string `json:"instruments"`
}

func main() {
	client.Init()
	db.Init()
	//pid := createProduct("token_name", "product_name", "description")
	//lid := createLayer(pid, "layer_name")
	//fmt.Println(lid)
	//createTrafficMap()
	//createFlight()
	//createVersion()
	requestVm()

}
func createProduct(token, productName, des string) int64 {

	product := db.Product{
		Token:       token,
		ProductName: productName,
		Status:      1,
		Description: des,
	}
	result := db.GetDBConn().Create(&product)
	if result.Error != nil {
		hlog.Errorf("error is create product,err = %v", result.Error)
	}
	fmt.Printf("some msg,id = %v, success line num = %v\n", product.ID, result.RowsAffected)
	return product.ID
}

func createLayer(pid int64, name string) int64 {
	layer := db.Layer{
		Name:         name,
		Status:       1,
		Pid:          1,
		Owner:        "renjunsheng",
		HashStrategy: "uid",
		HashUnit:     "normal",
	}
	result := db.GetDBConn().Create(&layer)
	if result.Error != nil {
		hlog.Errorf("error is create layer,err = %v", result.Error)
	}
	fmt.Printf("some msg,id = %v, success line num = %v\n", layer.ID, result.RowsAffected)
	return layer.ID
}

func createTrafficMap() {

	tmj := TrafficMapJson{}
	tmj.Flights = []int64{1}
	tmj.Pieces = []*effectiveTrafficPiece{
		&effectiveTrafficPiece{
			Begin:  0,
			Length: 1000,
			Flight: 1,
		},
	}
	tm, err := json.Marshal(tmj)
	if err != nil {
		fmt.Printf("json marshal error, err = %v", err)
		panic(err)
	}
	trafficMap := db.TrafficMap{
		TrafficMap:       string(tm),
		LayerID:          1,
		BucketTrafficMap: "",
		Pid:              1,
	}
	result := db.GetDBConn().Create(&trafficMap)
	if result.Error != nil {
		hlog.Errorf("error is create traffic map,err = %v", result.Error)
	}
	fmt.Printf("some msg,id = %v, success line num = %v\n", trafficMap.ID, result.RowsAffected)
}

func createFlight() {
	nowTimeInt64 := time.Now().Unix()
	nowTimeInt64 += 8640000
	rule := RuleItem{Instruments: map[string][]string{}}
	rule.Instruments["deviceType"] = []string{
		"phone",
		"pc",
		"web",
	}
	rule.Instruments["devicePlatform"] = []string{
		"ios",
		"android",
		"windows",
		"chrome",
	}
	rule.Instruments["city"] = []string{
		"beijing",
		"shanghai",
		"tianjin",
		"chongqing",
	}
	filter, err := json.Marshal(rule)
	if err != nil {
		fmt.Printf("json marshal error, err = %v", err)
		panic(err)
	}
	flight := db.Flight{
		Name:       "demo_flight1",
		Status:     3,
		StartTime:  time.Now(),
		Pid:        1,
		LayerID:    1,
		Fliter:     string(filter),
		TrafficMap: "",
		EndTime:    time.Unix(nowTimeInt64, 0),
		Kind:       1,
	}
	result := db.GetDBConn().Create(&flight)
	if result.Error != nil {
		hlog.Errorf("error is create flight,err = %v", result.Error)
	}
	fmt.Printf("some msg,id = %v, success line num = %v\n", flight.ID, result.RowsAffected)
}

func createVersion() {
	config := map[string]string{}
	config["version_config"] = "demo_version1"
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Printf("json marshal error, err = %v", err)
		panic(err)
	}
	version := db.Version{
		Name:     "demo_verison1",
		FlightID: 1,
		Status:   1,
		Summary:  "",
		Pid:      1,
		Config:   string(configStr),
	}
	result := db.GetDBConn().Create(&version)
	if result.Error != nil {
		hlog.Errorf("error is create version,err = %v", result.Error)
	}
	fmt.Printf("some msg,id = %v, success line num = %v\n", version.ID, result.RowsAffected)
}

func requestVm() {
	prop := map[string]string{}
	prop["product"] = "demo"
	prop["uid"] = "123456"
	prop["city"] = "beijing"
	prop["deviceType"] = "pc"
	prop["devicePlatform"] = "windows"
	reqbyte, err := json.Marshal(prop)
	if err != nil {
		panic(err)
	}
	req := meta_rpc.VersionReq{Req: string(reqbyte)}
	resp, err := client.VmClient.GetAbversion(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp is %v\n", *resp)
}
