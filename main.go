package go_scrape_dmm_co_jp

import (
	"fmt"
	"strconv"
	"flag"
)

func main() {
	
	var url string

	flag.Parse()
	if (flag.NArg() == 0) {
		url = "http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=172xrw00494/"
	} else if (flag.NArg() == 1) {
		url = flag.Arg(0)
	} else {
		panic("invalid args")
	}

	result := New(url)

	fmt.Println(result.ItemCode)
	fmt.Println(result.Title)
	fmt.Println(result.PackageImageThumbURL)
	fmt.Println(result.PackageImageURL)
	for index, value := range result.ActorList {
		fmt.Println(strconv.Itoa(index) + " : " + value.Name + " : " + value.ListPageURL)
	}
	for index, value := range result.SampleImageList {
		fmt.Println(strconv.Itoa(index) + " : " + value.ImageThumbURL + " : " + value.ImageURL)
	}
}