package config

const (
	OrderStatus           = "test status"
	OrderStatusNoOrder    = "you have no orders"
	OrderStatusNotPaid    = "your order formed but not paid"
	OrderStatusPaid       = "your order paid and accepted"
	OrderStatusProduction = "your order in production"
	OrderStatusSent       = "your order sent"
	OrderStatusCompleted  = "your order completed"

	StartMessage = "Привет!\nС моей помощью ты можешь " +
		"оформить заказ на подножки KAZA RACING.\nНиже находится список команд, которые тебе помогут.\n" +
		Start + " - Стартовое окно бота +\n" + List + " - Список доступных товаров\n" + Orders + " - Информация о твоих заказах\n" +
		Payment + "\n" + Status + " - Узнать статус твоего заказа\n" + Help +
		" - Список команд и возможность связаться, если возникли дополнительные вопросы.\n"
	Start   = "/start"
	List    = "/list"
	Orders  = "/order"
	Payment = "/payment"
	Status  = "/status"
	Help    = "/help"
	Basket  = "/basket"
)
