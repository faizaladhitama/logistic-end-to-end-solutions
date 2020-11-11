package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"regexp"
)

type OrderProduct struct {
	SKU string `json:"sku"`
	Qty float32 `json:"qty"`
	Price float32 `json:"price"`
}

type OrderLocation struct {
	Address string `json:"address"`
	Lat float32 `json:"lat"`
	Longt float32 `json:"longt"`	
}

type OrderCourier struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Order struct {
	OrderId string `json:"id"`
	OrderCode string `json:"code"`
	OrderReferenceNumber string `json:"reference_number"`
	OrderProduct []OrderProduct `json:"products"`
	OrderLocation OrderLocation `json:"location"`
	OrderCourier OrderCourier `json:"courier"` 
}

type Stock struct {
	Id int32
	SKU string
	ExpiryDate string
	InboundDate string
	InboundQuantity float32
	CurrentStock float32
}

type Product struct {
	Id int32
	Name string
	SKU string
	Expirable bool
}

type Response struct {
	Status int32 `json:"status"`
	Message string `json:"message"`
	Data map[string]string `json:"data"`
}

func main() {
	db, err := sql.Open("mysql", "root:@/swift")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()
	e.POST("/order", func(c echo.Context) error {
		m := map[string]string{}

		order := new(Order)
		if err := c.Bind(order); err != nil {
			return c.JSON(http.StatusBadRequest, order)
		}

		var stock Stock

		var orderProducts = order.OrderProduct
		var orderLocation = order.OrderLocation
		var orderReferenceNumber = order.OrderReferenceNumber
		var orderCourier = order.OrderCourier

		tx, err := db.Begin()

		for _, orderProduct := range orderProducts {
			var product Product

			// Check is expirable or not
			err = db.QueryRow(`SELECT sku, expirable, name, id from product where sku = ?`, 
				orderProduct.SKU).Scan(&product.SKU, &product.Expirable, &product.Name, &product.Id)
			if err != nil {
				response := &Response{
					Status: 400,
					Message: "SKU " + orderProduct.SKU + " is not available",
				}
				tx.Rollback()
				return c.JSON(http.StatusBadRequest, response)
			}

			if product.Expirable{
				// check stock
				err = db.QueryRow(`SELECT sku, expiry_date, inbound_date, 
					inbound_quantity, current_stock, id from stock 
					where sku = ? and expiry_date = 
					(select min(expiry_date) from stock where sku = ?)`, 
					orderProduct.SKU, orderProduct.SKU).Scan(&stock.SKU, &stock.ExpiryDate, &stock.InboundDate, 
						&stock.InboundQuantity, &stock.CurrentStock, &stock.Id)
				log.Println("Get stock FEFO")
				if err != nil {
					response := &Response{
						Status: 400,
						Message: "Stock not available for SKU " + orderProduct.SKU,
					}
					tx.Rollback()
					return c.JSON(http.StatusBadRequest, response)
				}
			}else{
				// check stock
				err = db.QueryRow(`SELECT sku, expiry_date, inbound_date, 
					inbound_quantity, current_stock, id from stock 
					where sku = ? and inbound_date = 
					(select max(inbound_date) from stock where sku = ?)`, 
					orderProduct.SKU, orderProduct.SKU).Scan(&stock.SKU, &stock.ExpiryDate, &stock.InboundDate, 
						&stock.InboundQuantity, &stock.CurrentStock, &stock.Id)

				log.Println("Get stock FIFO")
				if err != nil {
					response := &Response{
						Status: 400,
						Message: "Stock not available for SKU " + orderProduct.SKU,
					}
					tx.Rollback()
					return c.JSON(http.StatusBadRequest, response)
				}
			}

			var totalStock = stock.InboundQuantity + stock.CurrentStock

			// check stock is eligible
			if(totalStock >= orderProduct.Qty){
				// insert outbound
				 _, err := tx.Exec(`INSERT INTO outbound VALUES 
				 	(now(), ?, ?, ?, ?, ?, 'sales_order', ?, NULL)`,
				 	orderProduct.SKU, product.Name, orderProduct.Qty, 
				 	orderProduct.Price, orderProduct.Qty * orderProduct.Price,
				 	order.OrderReferenceNumber)

				if err != nil {
					log.Println("Insert outbound")
					response := &Response{
						Status: 400,
						Message: "Insert outbound failed for SKU " + orderProduct.SKU,
					}
					tx.Rollback()
					return c.JSON(http.StatusBadRequest, response)
				} 

				// update stock
				_, err = tx.Exec(`UPDATE stock SET current_stock = 
					current_stock - ? where id = ?`,
				 	orderProduct.Qty, stock.Id)

				if err != nil {
					log.Println("Update stock")
					response := &Response{
						Status: 400,
						Message: "Stock failed to be updated for SKU " + orderProduct.SKU,
					}
					tx.Rollback()
					return c.JSON(http.StatusBadRequest, response)
				}

			}else{
				response := &Response{
					Status: 400,
					Message: "Stock is less than order quantity for SKU " + orderProduct.SKU,
				}
				tx.Rollback()
				return c.JSON(http.StatusBadRequest, response)
			}
		}
		var currentTime = time.Now()
		var formattedTime = currentTime.Format("20060102150405.000000")

		log.Println(formattedTime)

		var orderCode = "ORD" + formattedTime
		
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			response := &Response{
				Status: 400,
				Message: "Failed generate order code",
			}
			tx.Rollback()
			return c.JSON(http.StatusBadRequest, response)
	    }
	    orderCode = reg.ReplaceAllString(orderCode, "")
		log.Println(orderCode) 

		// insert order
		_, err = tx.Exec(`INSERT INTO orders VALUES 
		 	(?, ?, 'Pending', ?, ?, ?, now(), NULL, ?, ?)`,
		 	orderCode, orderReferenceNumber, orderLocation.Address,
		 	orderLocation.Lat, orderLocation.Longt, 
		 	orderCourier.Name, orderCourier.Type)

		log.Println("Insert order")
		if err != nil {
			response := &Response{
				Status: 400,
				Message: "Order failed to be inserted",
			}
			tx.Rollback()
			return c.JSON(http.StatusBadRequest, response)
		}

		m["order_code"] = orderCode
		m["order_reference_number"] = orderReferenceNumber 
		response := &Response{
			Status: 201,
			Data: m,
		}
		tx.Commit()
		return c.JSON(http.StatusCreated, response)

	})
	e.Logger.Fatal(e.Start(":1323"))
}