package controllers

import (
	"fmt"
	"github.com/kamilyrb/bookstore_items-api/domain/items"
	"github.com/kamilyrb/bookstore_items-api/services"
	"github.com/kamilyrb/bookstore_oauth-go/oauth"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//	TODO Return error to the user,
		return
	}
	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemsService.Create(item)
	if err != nil {
		//TODO: Return error json to the user
	}
	fmt.Println(result)
	// TODO Return created item as json with HTTP Status  201 Created
}

func (*itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
