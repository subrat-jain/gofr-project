package main

import (
	"time"

	"gofr.dev/pkg/gofr"
)

type Car struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Car_Name  string `json:"car_name"`
	Car_Brand string `json:"car_brand"`
	Car_No    string `json:"Car_no"`
	In_Time   string `json:"in_time"`
}

type Old_Car struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Car_Name  string `json:"car_name"`
	Car_Brand string `json:"car_brand"`
	Car_No    string `json:"Car_no"`
	In_Time   string `json:"in_time"`
	Out_Time  string `json:"out_time"`
}

func main() {
	// initialise gofr object
	app := gofr.New()

	app.POST("/customer/{name}/{car_name}/{car_brand}/{car_no}", func(ctx *gofr.Context) (interface{}, error) {
		name := ctx.PathParam("name")
		car_name := ctx.PathParam("car_name")
		car_brand := ctx.PathParam("car_brand")
		car_no := ctx.PathParam("car_no")
		time := time.Now()

		// Inserting a customer row in database using SQL
		_, err := ctx.DB().ExecContext(ctx, "INSERT INTO garage (name, car_name, car_brand, car_no, in_time) VALUES (?,?,?,?,?)", name, car_name, car_brand, car_no, time.Format("2006-01-02"))

		return nil, err
	})

	app.GET("/customer", func(ctx *gofr.Context) (interface{}, error) {
		var garage []Car

		// Getting the customer from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM garage")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var customer Car
			if err := rows.Scan(&customer.ID, &customer.Name, &customer.Car_Name, &customer.Car_Brand, &customer.Car_No, customer.In_Time); err != nil {
				return nil, err
			}

			garage = append(garage, customer)
		}

		// return the customer
		return garage, nil
	})

	app.GET("/old", func(ctx *gofr.Context) (interface{}, error) {
		var garage []Old_Car

		// Getting the customer from the database using SQL
		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM complete")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var customer Old_Car
			if err := rows.Scan(&customer.ID, &customer.Name, &customer.Car_Name, &customer.Car_Brand, &customer.Car_No, customer.In_Time, customer.Out_Time); err != nil {
				return nil, err
			}

			garage = append(garage, customer)
		}

		// return the customer
		return garage, nil
	})

	app.DELETE("/customer/{id}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")
		time := time.Now()

		rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM car WHERE id=?", id)
		if err != nil {
			return nil, err
		}
		var garage []Car

		for rows.Next() {
			var customer Car
			if err := rows.Scan(&customer.ID, &customer.Name, &customer.Car_Name, &customer.Car_Brand, &customer.Car_No, customer.In_Time); err != nil {
				return nil, err
			}
			garage = append(garage, customer)
		}
		up, err := ctx.DB().ExecContext(ctx, "INSERT INTO complete (id, name, car_name, car_brand, car_no, in_time, out_time) VALUES (?,?,?,?,?,?,?)", id, garage[0].Name, garage[0].Car_Name, garage[0].Car_Brand, garage[0].Car_No, garage[0].In_Time, time.Format("2006-01-02"))
		if err != nil {
			return nil, err
		}
		_, err = ctx.DB().ExecContext(ctx, "DELETE FROM garage WHERE id=?", id)
		if err != nil {
			return nil, err
		}
		return up, nil
	})

	app.PUT("/customer/{id}/{name}/{car_name}/{car_brand}/{car_no}", func(ctx *gofr.Context) (interface{}, error) {
		id := ctx.PathParam("id")
		name := ctx.PathParam("name")
		car_name := ctx.PathParam("car_name")
		car_brand := ctx.PathParam("car_brand")
		car_no := ctx.PathParam("car_no")
		_, err := ctx.DB().ExecContext(ctx, "UPDATE garage SET name=?, car_name=?, car_brand=?, car_no=? WHERE id=?", name, car_name, car_brand, car_no, id)

		return nil, err
	})

	// it can be over-ridden through configs
	app.Start()
}
