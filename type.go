package goScrapeDmmCoJp

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/imaizm/go_scrape_dmm-common"
)

const baseDomain = "http://www.dmm.co.jp"

// ItemOfDmmCoJp : ItemOfDmmCoJp Info Struct
type ItemOfDmmCoJp struct {
	ItemCode             string
	Title                string
	PackageImageThumbURL string
	PackageImageURL      string
	SaleStartDate        time.Time
	DistStartDate        time.Time
	ActorList            []*Actor
	DirectorList         []*Director
	SeriesList           []*Series
	KeywordList          []*Keyword
	SampleImageList      []*goScrapeDmmCommon.SampleImage
}

// Actor : Actor Info Struct
type Actor struct {
	ListPageURL string
	Name        string
}

// Director : Director Info Struct
type Director struct {
	ID   string
	Name string
}

// Series : Series Info Struct
type Series struct {
	ID   string
	Name string
}

// Keyword : Keyword Info Struct
type Keyword struct {
	ID   string
	Name string
}

// GetItemInfoFromURL : create ItemOfDmmCoJp struct from url
func GetItemInfoFromURL(url string) *ItemOfDmmCoJp {

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	result := ItemOfDmmCoJp{}

	result.ItemCode = getItemCode(url)
	result.Title = getTitle(doc)
	result.PackageImageThumbURL = getPackageImageThumbURL(doc, result.ItemCode)
	result.PackageImageURL = getPackageImageURL(doc, result.ItemCode)
	result.ActorList = getActorList(doc)
	result.SampleImageList = getSampleImageList(doc)

	return &result
}

func getItemCode(url string) string {
	return goScrapeDmmCommon.GetItemCodeFromURL(url)
}

func getTitle(doc *goquery.Document) string {
	selection := doc.Find("#title")
	title := selection.First().Text()
	return title
}

func getPackageImageThumbURL(doc *goquery.Document, itemCode string) string {
	packageImageThumbURL := ""
	doc.Find("#package-src-" + itemCode).Each(func(index int, selection *goquery.Selection) {
		imgSrc, exists := selection.Attr("src")
		if exists {
			packageImageThumbURL = imgSrc
		}
	})
	return packageImageThumbURL
}

func getPackageImageURL(doc *goquery.Document, itemCode string) string {
	packageImageURL := ""
	doc.Find("#" + itemCode).Each(func(index int, selection *goquery.Selection) {
		aHref, exists := selection.Attr("href")
		if exists {
			packageImageURL = aHref
		}
	})
	return packageImageURL
}

func getActorList(doc *goquery.Document) []*Actor {
	var actorList []*Actor

	doc.Find("#performer > a").Each(func(index int, selection *goquery.Selection) {
		actor := Actor{}

		actor.Name = selection.Text()

		href, exists := selection.Attr("href")
		if exists {
			actor.ListPageURL = baseDomain + href
		}

		actorList = append(actorList, &actor)
	})

	return actorList
}

func getSampleImageList(doc *goquery.Document) []*goScrapeDmmCommon.SampleImage {
	return goScrapeDmmCommon.GetSampleImageList(doc)
}

// Search : search by keyword
func Search(searchKeyword string) []*SearchListItem {
	url := baseDomain + "/search/=/searchstr=" + searchKeyword + "/limit=120/sort=rankprofile/"

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	listSelection := getSearchListSelection(doc)
	searchListItemList := getSearchListItemList(listSelection)

	return searchListItemList
}

func getSearchListSelection(doc *goquery.Document) *goquery.Selection {
	listSelection := doc.Find("ul#list li")
	if _, err := listSelection.Html(); err != nil {
		panic("goScrapeDmmCoJp.getSearchListSelection: search list retrive fail")
	}
	return listSelection
}

// SearchListItem : list item of search result
type SearchListItem struct {
	ItemDetailURL string
	Title         string
}

func getSearchListItemList(listSelection *goquery.Selection) []*SearchListItem {
	var searchListItemList []*SearchListItem

	listSelection.Each(func(index int, selection *goquery.Selection) {
		searchListItem := SearchListItem{}

		listItemSelectionCheck := selection.Find("div p.tmb a")
		if listItemSelectionCheck.Text() == "" {
			panic("goScrapeDmmCoJp.getSearchListItemList: list item retrive fail")
		}
		listItemSelection := listItemSelectionCheck.First()
		aHref, _ := listItemSelection.Attr("href")
		searchListItem.ItemDetailURL = aHref

		listItemImageSelectionCheck := listItemSelection.Find("span img")
		if _, err := listItemImageSelectionCheck.Html(); err != nil {
			panic("goScrapeDmmCoJp.getSearchListItemList: list item image retrive fail")
		}
		listItemImageSelection := listItemImageSelectionCheck
		imgAlt, _ := listItemImageSelection.Attr("alt")
		searchListItem.Title = imgAlt

		if filterSearchListItem(&searchListItem) {
			searchListItemList = append(searchListItemList, &searchListItem)
		}
	})

	return searchListItemList
}

func filterSearchListItem(searchListItem *SearchListItem) bool {
	if strings.Contains(searchListItem.ItemDetailURL, "/rental/") {
		return false
	}

	return true
}
