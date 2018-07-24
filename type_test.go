package goScrapeDmmCoJp

import (
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	url := "http://www.dmm.co.jp/digital/videoa/-/detail/=/cid=172xrw00494/"

	result := New(url)

	t.Log("ItemCode : " + result.ItemCode)
	t.Log("Title : " + result.Title)
	t.Log("PackageImageThumbURL : " + result.PackageImageThumbURL)
	t.Log("PackageImageURL : " + result.PackageImageURL)
	t.Log("ActorList :")
	for index, value := range result.ActorList {
		t.Log("\t" + strconv.Itoa(index) + " : " + value.Name + " : " + value.ListPageURL)
	}
	t.Log("SampleImageList :")
	for index, value := range result.SampleImageList {
		t.Log("\t" + strconv.Itoa(index) + " : " + value.ImageThumbURL + " : " + value.ImageURL)
	}
}
