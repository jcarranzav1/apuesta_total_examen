package entity

type Product struct {
	ID       uint
	Name     string
	Price    float64
	Quantity int
}

type Payment struct {
	ID            uint
	Amount        float64 // Monto del pago
	Currency      string  // Moneda del pago
	PaymentMethod string  // Método de pago (tarjeta de crédito, PayPal, etc.)
	Products      []Product
	//Status        string    // Estado del pago (pendiente, completado, fallido)
}
