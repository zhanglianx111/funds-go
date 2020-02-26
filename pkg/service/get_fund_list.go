package service

import (
	"fmt"
	"github.com/zhanglianx111/funds-go/pkg/model"
	"io/ioutil"
	"net/http"
)


var (
	FundListUrl = "http://fund.eastmoney.com/js/fundcode_search.js"
)

type FList struct {
	Count int
	Fund []*model.FundsList
}



func GetFundsList() error {
	_, _ = httpGet(FundListUrl)
	return nil
}

func httpGet(url string) ([]*model.FundsList, error){
	resp, err :=  http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}


	//fmt.Println(string(body))
	l := len(body) - 1
	b := body[7:l]
	fmt.Println(string(b))
	return nil, nil
}

func format(fl []byte) (*FList){

}