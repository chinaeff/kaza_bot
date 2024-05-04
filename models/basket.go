package models

import "sync"

type Basket struct {
	Items []BasketItem `json:"items"`
}
type BasketItem struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

var (
	UserBasketsMutex sync.Mutex
	UserBaskets      = make(map[int64]*Basket)
)

func GetUserBasket(userID int64) *Basket {
	usersBasket := UserBaskets[userID]
	if usersBasket == nil {
		usersBasket = &Basket{}
		UserBaskets[userID] = usersBasket
	}
	return usersBasket
}
func AddUserBasket(userID int64, product Product, quantity int) {
	basket := GetUserBasket(userID)
	for i, item := range basket.Items {
		if item.Product.ID == product.ID {
			basket.Items[i].Quantity += quantity
			return
		}
	}
	basket.Items = append(basket.Items, BasketItem{Product: product, Quantity: quantity})
}

func GetBasketTotal(basket *Basket) int {
	total := 0
	for _, item := range basket.Items {
		total += item.Product.Price * item.Quantity
	}
	return total
}
