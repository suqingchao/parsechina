package addr

import (
"encoding/json"
"errors"
"io/ioutil"
"os"
"strings"
)

// Address 通信地址
type Address struct {
	ProvinceCode string // 省
	CityCode     string // 市
	AreaCode     string // 区
	Detail       string // 街道以下具体地址
}

type provinceData struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type cityData struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	ProvinceCode string `json:"provinceCode"`
}

type areaData struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	CityCode     string `json:"cityCode"`
	ProvinceCode string `json:"provinceCode"`
}

// Parse 解析地址并规范化
func Parse(src string) (*Address, error) {

	var (
		a   = &Address{}
		err error
	)
	if err = getProvince(&src, a); err == nil {
		if err = getCity(&src, a); err == nil {
			if err = getArea(&src, a); err == nil {
			}
		}
	}

	return a, err
}

func getProvince(src *string, a *Address) (err error) {

	//广西壮族自治区
	*src = strings.Replace(*src, `广西省`, `广西壮族自治区`, -1)

	err = errors.New(`no matched province`)
	for _, v := range psObj {

		//兼容不写"省"字的情况
		if strings.Index(*src, `省`) == -1 {
			v.Name = strings.Replace(v.Name, `省`, ``, -1)
		}

		if strings.Index(*src, v.Name) != -1 {
			a.ProvinceCode = v.Code
			*src = strings.Replace(*src, v.Name, ``, -1)
			err = nil
			break
		}
	}

	return
}

var bigCity = map[string]bool{"11": true, "12": true, "50": true, "31": true}

func getCity(src *string, a *Address) (err error) {

	//处理直辖市
	if _, ok := bigCity[a.ProvinceCode]; ok {
		return
	}

	err = errors.New(`no matched city`)
	for _, v := range csObj {
		//兼容不写"市"字的情况
		if strings.Index(*src, `市`) == -1 {

			v.Name = strings.Replace(v.Name, `市`, ``, -1)
		}
		if v.ProvinceCode == a.ProvinceCode && strings.Index(*src, v.Name) != -1 {

			a.CityCode = v.Code
			*src = strings.Replace(*src, v.Name, ``, -1)
			err = nil
			break
		}
	}
	return
}

func getArea(src *string, a *Address) (err error) {
	err = errors.New(`no matched area`)
	for _, v := range asObj {
		if strings.Index(*src, v.Name) != -1 {
			a.AreaCode = v.Code
			a.Detail = strings.Replace(*src, v.Name, ``, -1)
			err = nil
			break
		}
	}
	return
}

var psObj = make([]provinceData, 0)
var csObj = make([]cityData, 0)
var asObj = make([]areaData, 0)

func init() {
	//省数据
	psFile := read(`./china/provinces.json`)             ///////
	csFile := read(`./china/cities.json`)           /////
	asFile := read(`./china/areas.json`)         ///////
	_ = json.Unmarshal(psFile, &psObj)
	_ = json.Unmarshal(csFile, &csObj)
	_ = json.Unmarshal(asFile, &asObj)

}

func read(target string) []byte {
	f, err := os.Open(target)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return data
}
