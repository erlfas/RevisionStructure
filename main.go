package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Sending struct {
	NewestVersion     string
	Produkt           Produkt
	Sendingstjenester []Sendingstjeneste
}

func (s *Produkt) NewestProdukt() Produkt {
	var newestTime time.Time
	var newestProdukt Produkt

	t := time.Now()

	for _, val := range Produkt.Values {
		if val.Created > newestTime {
			newestTime = val.Created
			newestProdukt = val
		}
	}

	return newestProdukt
}

type Produkt struct {
	NewestVersion string
	Values        []Value
}

type Sendingstjeneste struct {
	NewestVersion string
	Values        []Value
}

type Value struct {
	Name    string
	Version string
	User    string
	Value   interface{}
	Created time.Time
}

func main() {

	sending := &Sending{
		Produkt: Produkt{
			"2",
			[]Value{
				Value{"Produkt", "1", "eef", "B"},
				Value{"Produkt", "2", "eef", "Z"},
			},
		},
		Sendingstjenester: []Sendingstjeneste{
			Sendingstjeneste{
				NewestVersion: "1",
				Values: []Value{
					Value{"Sendingstjeneste", "1", "eef", "A3"},
				},
			},
			Sendingstjeneste{
				NewestVersion: "1",
				Values: []Value{
					Value{"Sendingstjeneste", "1", "eef", "A4"},
				},
			},
		},
	}

	jsonBytes, _ := json.Marshal(sending)

	fmt.Println(string(jsonBytes))

}
