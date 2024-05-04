package bot

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"kaza_bot/config"
	"kaza_bot/models"
	"log"
)

//const (
//	StartMessage = "Starting message\n" +
//		Start + "\n" + List + "\n" + Orders + "\n" + Payment + "\n" + Status + "\n" + Help + "\n"
//	Start   = "/start"
//	List    = "/list"
//	Orders  = "/order"
//	Payment = "/payment"
//	Status  = "/status"
//	Help    = "/help"
//	Basket  = "/basket"
//)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil {
		HandleTextMessage(bot, update.Message)
	}
}

func HandleTextMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Text {
	case config.Start:
		HandleStartMessage(bot, message)
	case config.Orders:
		HandleOrdersMessage(bot, message)
	case config.List:
		HandleListMessage(bot, message)
	case config.Payment:
		HandlePaymentMessage(bot, message)
	case config.Basket:
		HandleBasketMessage(bot, message)
	case config.Status:
		HandleStatusMessage(bot, message)
	case config.Help:
		HandleHelpMessage(bot, message)
	default:
		HandleUnknownMessage(bot, message)
	}
}

func HandleStartMessage(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	reply := config.StartMessage
	message := tgbotapi.NewMessage(msg.Chat.ID, reply)
	_, err := bot.Send(message)
	if err != nil {
		log.Println("error sending /start reply", err)
	}
}

func HandleListMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	db, err := sql.Open("sqlite3", "./database/products.db")
	if err != nil {
		log.Println("error opening database connection:", err)
		return
	}
	defer db.Close()

	products, err := GetAllProducts(db)
	if err != nil {
		log.Println("error getting products:", err)
		return
	}

	reply := "List of products:\n"
	for _, product := range products {
		reply += fmt.Sprintf("ID: %d, Name: %s, Price: %d, Photo: %s\n",
			product.ID, product.Name, product.Price, product.Image)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, reply)
	_, err = bot.Send(msg)
	if err != nil {
		log.Println("error sending /list message:", err)
	}
}

func HandleStatusMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	db, err := sql.Open("sqlite3", "./database/orders.db")
	if err != nil {
		log.Println("error opening database of orders", err)
	}
	defer db.Close()

	user, err := GetUserByID(db, message)
	if err != nil {
		log.Println("error getting user by id", err)
		return
	}

	if user == nil || len(user.Orders) == 0 {
		reply := config.OrderStatusNoOrder
		msg := tgbotapi.NewMessage(message.Chat.ID, reply)
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("error sending /status message", err)
		}
		return
	}
	reply := "Status of your order:\n"
	for _, order := range user.Orders {
		reply += fmt.Sprintf("Order: %d\nStatus: %s\n", order.OrderID, order.Status)
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, reply)
	_, err = bot.Send(msg)
	if err != nil {
		log.Println("error sending /status message", err)
	}
}

func HandlePaymentMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	db, err := sql.Open("sqlite3", "./database/orders.db")
	if err != nil {
		log.Println("error opening database of orders", err)
	}
	defer db.Close()

	user, err := GetUserByID(db, message)
	if err != nil {
		log.Println("error getting user by id", err)
		return
	}

	if user == nil || len(user.Orders) == 0 {
		reply := config.OrderStatusNoOrder
		msg := tgbotapi.NewMessage(message.Chat.ID, reply)
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("error sending /status message", err)
		}
		return
	}
	totalAmount := 0
	for _, order := range user.Orders {
		totalAmount += order.Product.Price
	}
	reply := fmt.Sprintf("Total amount is %d", totalAmount)
	reply += "Go to payment service for payment"

	msg := tgbotapi.NewMessage(message.Chat.ID, reply)
	msg.ParseMode = "Markdown"
	_, err = bot.Send(msg)
	if err != nil {
		log.Println("error sending /payment message", err)
	}
}

func HandleHelpMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	reply := config.StartMessage

	msg := tgbotapi.NewMessage(message.Chat.ID, reply)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("error sending /help message", err)
	}
}

func HandleBasketMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	db, err := sql.Open("sqlite3", "./database/products.db")
	if err != nil {
		log.Println("error opening database connection:", err)
		return
	}
	defer db.Close()

	user, err := GetUserByID(db, message)
	if err != nil {
		log.Println("error getting user by id", err)
		return
	}

	if user == nil || len(user.Orders) == 0 {
		reply := "You have no orders"
		msg := tgbotapi.NewMessage(message.Chat.ID, reply)
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("error sending /basket reply:", err)
		}
		return
	}
	userBasket := models.GetUserBasket(user.UserID)
	total := models.GetBasketTotal(userBasket)

	reply := "Items:\n"
	for _, item := range userBasket.Items {
		reply += fmt.Sprintf("%s (%d) - %d\n", item.Product.Name, item.Product.Price, item.Product.Quantity)
	}
	reply += fmt.Sprintf("Total: %d", total)

	msg := tgbotapi.NewMessage(message.Chat.ID, reply)
	_, err = bot.Send(msg)
	if err != nil {
		log.Println("error sending /basket message", err)
	}
}

func HandleUnknownMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	reply := config.StartMessage

	msg := tgbotapi.NewMessage(message.Chat.ID, reply)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("error sending unknown message", err)
	}
}

func HandleOrdersMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	db, err := sql.Open("sqlite3", "./database/orders.db")
	if err != nil {
		log.Println("error opening database of orders", err)
	}
	defer db.Close()

	user, err := GetUserByID(db, message)
	if err != nil {
		log.Println("error getting user by id", err)
		return
	}
	if user == nil || len(user.Orders) == 0 {
		reply := "you have no orders"
		msg := tgbotapi.NewMessage(message.Chat.ID, reply)
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("error sending reply", err)
		}
		return
	}
	reply := "Order's details:"
	for _, order := range user.Orders {
		totalAmount := 0
		productTotal := make(map[string]int)

		for _, order := range user.Orders {
			productAmount := order.Product.Price * order.Product.Quantity
			totalAmount += productAmount

			productTotal[order.Product.Name] += productAmount
		}
		reply += fmt.Sprintf("Order %d\n", order.OrderID)
		for productName, productAmount := range productTotal {
			reply += fmt.Sprintf("Product: %s, Total Amount: %d\n", productName, productAmount)
		}
		reply += fmt.Sprintf("Date: %s, Status: %s\n", order.DateOfOrder, order.Status)

		msg := tgbotapi.NewMessage(message.Chat.ID, reply)
		_, err := bot.Send(msg)
		if err != nil {
			log.Println("error sending /orders message", err)
		}
	}
}

func GetUserByID(db *sql.DB, message *tgbotapi.Message) (*models.User, error) {
	row := db.QueryRow("SELECT user_id, username FROM users WHERE user_id = ?", message.From.ID)

	var user models.User
	err := row.Scan(&user.UserID, &user.Username)
	if err != nil {
		log.Println("error getting user by id", err)
	}

	rows, err := db.Query("SELECT order_id, status, date_of_order FROM orders WHERE user_id = ?", message.From.ID)
	if err != nil {
		log.Println("error getting user by id", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.OrderID, &order.Status, &order.DateOfOrder)
		if err != nil {
			log.Println("error getting user by id", err)
		}

		product, err := GetProductByID(db, order.Product.ID)
		if err != nil {
			log.Println("error getting product by id", err)
		}
		order.Product = *product

		user.Orders = append(user.Orders, order)
	}
	return &user, nil
}

func GetProductByID(db *sql.DB, productID int) (*models.Product, error) {
	row := db.QueryRow("SELECT name, price FROM products WHERE id = ?", productID)

	var product models.Product
	err := row.Scan(&product.Name, &product.Price)
	if err != nil {
		log.Println("error getting product by id", err)
	}
	return &product, nil
}

func GetAllProducts(db *sql.DB) ([]models.Product, error) {
	rows, err := db.Query("SELECT id, name, price, image FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Image)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
