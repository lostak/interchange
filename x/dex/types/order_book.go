// x/dex/types/order_book.go
package types

import (
	"errors"
	"sort"
)

func NewOrderBook() OrderBook {
	return OrderBook{
		IdCount: 0,
	}
}

const (
	MaxAmount = int32(100000)
	MaxPrice  = int32(100000)
)

type Ordering int

const (
	Increasing Ordering = iota
	Decreasing
)

var (
	ErrMaxAmount		= errors.New("max amount reached")
	ErrMaxPrice		= errors.New("max price reached")
	ErrZeroAmount		= errors.New("amount is zero")
	ErrZeroPrice		= errors.New("price is zero")
	ErrOrderNotFound	= errors.New("order not found")
)

func (book *OrderBook) appendOrder(creator string, amount int32, price int32, ordering Ordering) (int32, error){
	if err := checkAmountAndPrice(amount, price); err != nill {
		return 0, err
	}

	// initialize order
	var order Order
	order.Id = book.GetNextOrderID()
	order.Creator = creator
	order.Amount = amount
	order.Price = price
	// increment id tracker
	book.IncrementNextOrderID()
	// insert the order
	book.inserOrder(order, ordering)
	return orderId, nil
}

func checkAmountAndPrice(amount int32, price int32) error {
	if amount == int32(0) {
		return ErrZeroAmount
	}
	if amount > MaxAmount {
		return ErrMaxAmount
	}
	if price == int32(0) {
		return ErrZeroPrice
	}
	if price > MaxPrice {
		return ErrMaxPrice
	}
	return nil
}

func (book OrderBook) GetNextOrderID() int32 {
	return book.IdCount
}

func (book *OrderBook) IncremenetNextOrderID() {
	book.IdCount++
}

func (book *OrderBook) inserOrder(order Order, ordering Ordering) {
	if len(book.Orders) > 0 {
		var i int
		// ge tindex of the new order depending on the provided ordering
		if ordering == Increasing {
			i = sort.Search(len(book.Orders), func(i int) bool { return book.Orders[i].Price > order.Price })
		} else {
			i = sort.Search(len(book.Orders), func(i int) bool { return book.Orders[i].Price < order.Price })
		}
		// inser order
		orders := append(book.Orders, &order)
		copy(orders[i+1:], orders[i:])
		orders[i] = &order
		book.Orders = orders
	} else {
		book.Orders = append(book.Orders, &order)
	}

