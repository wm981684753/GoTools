package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	secret = "yuxuan3507"
	url    = "http://test-platform.whgxwl.com/"
)

func main() {
	fmt.Println(`
FS站点信息生成，根据提示进行操作
-------------------------------
	`)
	getCountryList()
}

type Country struct {
	Data struct {
		CountryList struct {
			Source []struct {
				Countries_name string
				Iso_code       string
			}
		}
	}
}

type CountryInit []struct {
	Countries_name string
	Iso_code       string
}

type sitelogt struct {
	country  string
	lang     string
	currency string
}

var sitelog sitelogt

//获取国家站点列表
func getCountryList() {
	fmt.Println(`请选择对应的国家...`)
	body := get("api/website/site")
	var countryInfo Country
	json.Unmarshal(body, &countryInfo)

	//国家太多不适合选择，手动定义常用的国家
	var defindCountry CountryInit
	json.Unmarshal([]byte(`[
		{"Countries_name":"新加坡（SG）","Iso_code":"SG"},
		{"Countries_name":"美国（US）","Iso_code":"US"},
		{"Countries_name":"德国（DE）","Iso_code":"DE"},
		{"Countries_name":"俄罗斯（RU）","Iso_code":"RU"},
		{"Countries_name":"西班牙（ES）","Iso_code":"ES"},
		{"Countries_name":"日本（JP）","Iso_code":"JP"},
		{"Countries_name":"意大利（IT）","Iso_code":"IT"},
		{"Countries_name":"法国（FR）","Iso_code":"FR"},
		{"Countries_name":"澳大利亚（AU）","Iso_code":"AU"},
		{"Countries_name":"加拿大（CA）","Iso_code":"CA"}
	]`), &defindCountry)
	for v, k := range defindCountry {
		fmt.Println(v+1, "：", k.Countries_name)
	}
	var scale int
	fmt.Scanln(&scale)
	if scale > 0 && scale <= len(defindCountry) {
		sitelog.country = defindCountry[scale-1].Countries_name
		getCountrySite(defindCountry[scale-1].Iso_code)
	} else {
		fmt.Println("选择的国家不存在~_~，请重新选择")
		getCountryList()
	}
}

type Site struct {
	Data []struct {
		Currency string
		Language string
	}
}

//根据国家code获取站点信息
func getCountrySite(iso_code string) {
	body := post("api/website/filterSiteByCountryCode", "iso_code="+iso_code)
	var siteInfo Site
	json.Unmarshal(body, &siteInfo)

	fmt.Println(`请选择对应的语言|货币...`)
	for v, k := range siteInfo.Data {
		fmt.Println(v+1, "：", k.Language, " | ", k.Currency)
	}

	var scale int
	fmt.Scanln(&scale)
	if scale > 0 && scale <= len(siteInfo.Data) {
		sitelog.lang = siteInfo.Data[scale-1].Language
		sitelog.currency = siteInfo.Data[scale-1].Currency
		getWebSiteInfo(iso_code, siteInfo.Data[scale-1].Currency, siteInfo.Data[scale-1].Language)
	} else {
		fmt.Println("选择的语言不存在啊~_~，请重新选择")
		getCountrySite(iso_code)
	}
}

type Website struct {
	Data map[string]interface{}
}

//获取站点websiteinfo信息
func getWebSiteInfo(iso_code string, currency string, language string) {
	body := post("api/website/updateSiteInfo", "iso_code="+iso_code+"&language="+language+"&currency="+currency)
	var webSiteInfo Website
	json.Unmarshal(body, &webSiteInfo)
	mjson, _ := json.Marshal(webSiteInfo.Data)
	signature := base64.StdEncoding.EncodeToString(mjson)
	fmt.Println("来了来了，", sitelog.country, sitelog.lang, "|", sitelog.currency, "站点webSiteInfo玩命生成中...")
	fmt.Println(signature)
	scale := "a"
	fmt.Println(`
回车可以再次生成，快乐加倍
	`)
	fmt.Scanln(&scale)
	if scale != "" {
		getCountryList()
	}
}

//get请求
func get(api string) []byte {
	response, err := http.Get(url + api)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	return body
}

//post请求
func post(api string, param string) []byte {
	response, err := http.Post(
		url+api,
		"application/x-www-form-urlencoded",
		strings.NewReader(param),
	)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return body
}
