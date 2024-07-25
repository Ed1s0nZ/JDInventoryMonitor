package main

import (
	"encoding/json"
	"fmt"
)

type StockInfo struct {
	FreshEdi       interface{} `json:"freshEdi"`
	SidDely        string      `json:"sidDely"`
	Channel        int         `json:"channel"`
	Rid            string      `json:"rid"`
	Sid            string      `json:"sid"`
	DcId           string      `json:"dcId"`
	IsPurchase     bool        `json:"IsPurchase"`
	Eb             string      `json:"eb"`
	Ec             string      `json:"ec"`
	StockState     int         `json:"StockState"`
	Ab             string      `json:"ab"`
	CanAddCart     string      `json:"canAddCart"`
	Ac             string      `json:"ac"`
	Ad             string      `json:"ad"`
	Ae             string      `json:"ae"`
	SkuState       int         `json:"skuState"`
	PopType        int         `json:"PopType"`
	Af             string      `json:"af"`
	Ag             string      `json:"ag"`
	StockStateName string      `json:"StockStateName"`
	M              string      `json:"m"`
	Rfg            int         `json:"rfg"`
	ArrivalDate    string      `json:"ArrivalDate"`
	V              string      `json:"v"`
	Rn             int         `json:"rn"`
	Dc             string      `json:"dc"`
}

type OutputInfo struct {
	Description    string `json:"Description"`
	StockStateName string `json:"StockStateName"`
}

func finder(Data string) string {
	data := Data

	var stockData map[string]StockInfo
	err := json.Unmarshal([]byte(data), &stockData)
	if err != nil {
		fmt.Println(err)
		return string(err.Error())
	}

	// SKU descriptions
	skuDescriptions := map[string]string{
		"10086628263328": "20张白边相纸+大耳狗盛满樱桃相册",
		"1566641027":     "40张白边相纸",
		"10040203309050": "20张白边相纸",
		"10086628263327": "20张白边相纸+大耳狗海绵星空相册",
	}

	outputData := make(map[string]OutputInfo)

	for sku, info := range stockData {
		if description, exists := skuDescriptions[sku]; exists {
			outputData[sku] = OutputInfo{
				Description:    description,
				StockStateName: info.StockStateName,
			}
		}
	}

	outputJSON, err := json.MarshalIndent(outputData, "", "  ")
	if err != nil {
		fmt.Println(err)
		return string(err.Error())
	}

	return string(outputJSON)
}

func jsonToString(jsonData string) string {
	var stockData map[string]struct {
		Description    string `json:"Description"`
		StockStateName string `json:"StockStateName"`
	}

	err := json.Unmarshal([]byte(jsonData), &stockData)
	if err != nil {
		fmt.Println("Error:", err)
		return err.Error()
	}

	outputString := ""
	fmt.Println("stockData:", stockData)
	for _, v := range stockData {
		outputString += fmt.Sprintf("%s 【%s】\n", v.Description, v.StockStateName)
		fmt.Println(v.Description)
	}

	fmt.Println(outputString)
	return outputString
}

// 检查是否全部为无货
func checkAllOutOfStock(jsonData string) (bool, error) {
	var stockData map[string]struct {
		Description    string `json:"Description"`
		StockStateName string `json:"StockStateName"`
	}

	err := json.Unmarshal([]byte(jsonData), &stockData)
	if err != nil {
		return false, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	for _, v := range stockData {
		if v.StockStateName != "无货" {
			return false, nil
		}
	}

	return true, nil
}
