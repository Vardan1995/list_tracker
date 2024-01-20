package core

import (
	"fmt"
	"time"

	"github.com/Vardan1995/list_tracker/entity"
	"github.com/Vardan1995/list_tracker/service"
	"github.com/Vardan1995/list_tracker/utils"
	"github.com/gocolly/colly/v2"
)

var ArchiveExist = service.NewArchiveService().CreateIfNotExists
var preferenceService = service.NewPreferenceService()
var userService = service.NewUserService()

// Visits links based on provided preferences and sends an email if new items are found.
func Run(preferences []entity.Preference, email string) {

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
	for _, v := range preferences {
		c.Visit(v.Link)
	}
	if result != "" {
		utils.SendEmail(result, email)

	}

}

// Track continuously monitors users' preferences and runs the tracking process.
func Track() {
	for {
		var users []entity.User
		if err := userService.GetUsers(&users); err != nil {
			fmt.Printf("err: %v\n", err)
		}

		for _, user := range users {
			var preferences []entity.Preference
			err := preferenceService.GetUserPreference(&preferences, user.ID)
			if err != nil {
				fmt.Printf("err: %v\n", err)

			}

			Run(preferences, user.Email)

		}
		time.Sleep(time.Second * 10)
	}
}
