package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Shipment struct {
	NewestVersion int64
	Product       Product
	Services      []Service
}

func (s *Shipment) GetLatestVersionCopy() Shipment {
	shipment := Shipment{
		NewestVersion: s.NewestVersion,
		Product:       s.NewestProduct(),
	}

	return shipment
}

func (s *Shipment) NewestServices() []Value {
	serviceValues := make([]Value, 0)

	for _, s := range s.Services {
		for _, val := range s.Values {
			if val.Name != "" && val.Value != nil {
				if s.NewestVersion == val.Version {
					serviceValues = append(serviceValues, val)
				}
			}
		}
	}

	return serviceValues
}

func (s *Shipment) NewestProduct() Product {
	var newestTime time.Time
	var newestProductValue Value

	for _, val := range s.Product.Values {
		if val.Created.After(newestTime) {
			newestTime = val.Created
			newestProductValue = val
		}
	}

	newestProduct := Product{
		NewestVersion: s.Product.NewestVersion,
		Values:        []Value{newestProductValue},
	}

	return newestProduct
}

type Product struct {
	NewestVersion int64
	Values        []Value
}

type Service struct {
	NewestVersion int64
	Values        []Value
}

type Value struct {
	Name    string
	Version int64
	User    string
	Value   interface{}
	Created time.Time
}

func main() {

	sending := &Shipment{
		NewestVersion: 2,
		Product: Product{
			2,
			[]Value{
				Value{"Produkt", 1, "eef", "B", time.Now()},
				Value{"Produkt", 2, "eef", "Z", time.Now().Add(time.Duration(10))},
			},
		},
		Services: []Service{
			Service{
				NewestVersion: 1,
				Values: []Value{
					Value{"Sendingstjeneste", 1, "eef", "A3", time.Now()},
				},
			},
			Service{
				NewestVersion: 1,
				Values: []Value{
					Value{"Sendingstjeneste", 1, "eef", "A4", time.Now()},
				},
			},
		},
	}

	jsonBytes, _ := json.Marshal(sending)

	fmt.Println(string(jsonBytes))

	fmt.Println()
	fmt.Println()

	newestProdukt := sending.NewestProduct()

	newestProduktJsonBytes, _ := json.Marshal(newestProdukt)

	fmt.Println(string(newestProduktJsonBytes))

	fmt.Println()
	fmt.Println()

	newestServices := sending.NewestServices()
	newestServicesJsonBytes, _ := json.Marshal(newestServices)

	fmt.Println(string(newestServicesJsonBytes))

}
