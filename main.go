package main

import (
	"fmt"
	"parsechina/addr"
)

func main(){
	//正常数据
	addrStr := `四川省成都市武侯区天府新谷9999号楼99-1`
	fmt.Println(addr.Parse(addrStr))
	addrStr = `浙江省杭州市余杭区文一西路969号`
	fmt.Println(addr.Parse(addrStr))

	//不写 "省"字 或者 "市" 的数据
	addrStr = `四川成都市武侯区天府新谷9999号楼99-1`
	fmt.Println(addr.Parse(addrStr))
	addrStr = `四川成都武侯区天府新谷9999号楼99-1`
	fmt.Println(addr.Parse(addrStr))

	//直辖市数据
	addrStr = `北京市朝阳区定福庄西街1号`
	fmt.Println(addr.Parse(addrStr))
	addrStr = `重庆市九龙坡区杨家坪前进支路1号跃华新都26层`
	fmt.Println(addr.Parse(addrStr))

	//自治区问题
	addrStr = `广西壮族自治区南宁市青秀区桃源路6号`
	fmt.Println(addr.Parse(addrStr))

	//手滑写错的数据
	addrStr = `广西省南宁市青秀区桃源路6号`
	fmt.Println(addr.Parse(addrStr))

	//假数据,异常数据
	addrStr = `日本省韩国市中心区桃源路98号`
	fmt.Println(addr.Parse(addrStr))
}
