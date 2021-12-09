package core

import (
	"fmt"
	"time"

	"github.com/Vardan1995/filter_notifyer/entity"
	"github.com/Vardan1995/filter_notifyer/service"
	"github.com/Vardan1995/filter_notifyer/utils"
	"github.com/gocolly/colly/v2"
)

var ArchiveExist = service.NewArchiveService().CreateIfNotExists
var filterService = service.NewFilterService()
var userService = service.NewUserService()

func Run(filters []entity.Filter, email string) {

	var result string
	c := colly.NewCollector()
	c.OnHTML(`a[href^="/item/"]`, func(e *colly.HTMLElement) {
		product_id := e.Attr("href")
		exist := ArchiveExist(email, product_id)
		if !exist {
			result += "\n" + e.Text + "\n" + "https://www.list.am" + product_id
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting link from ", email)
	})
	for _, v := range filters {
		c.Visit(v.Link)
	}
	if result != "" {
		utils.SendEmail(result, email)

	}

}

func Track() {
	for {
		var users []entity.User
		if err := userService.GetUsers(&users); err != nil {
			fmt.Printf("err: %v\n", err)
		}

		for _, user := range users {
			var filters []entity.Filter
			err := filterService.GetUserFilter(&filters, user.ID)
			if err != nil {
				fmt.Printf("err: %v\n", err)

			}

			Run(filters, user.Email)

		}
		time.Sleep(time.Second * 10)
	}
}
