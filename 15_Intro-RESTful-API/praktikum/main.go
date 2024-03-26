// Task 03
// Endpoint https://www.fruityvice.com/api/fruit/{name}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var ENDPOINT = "https://www.fruityvice.com/api/fruit/"

type Fruit struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Family     string     `json:"family"`
	Order      string     `json:"order"`
	Genus      string     `json:"genus"`
	Nutritions Nutritions `json:"nutritions"`
}

type Nutritions struct {
	Calories      float32 `json:"calories"`
	Fat           float32 `json:"fat"`
	Sugar         float32 `json:"sugar"`
	Carbohydrates float32 `json:"carbohydrates"`
	Protein       float32 `json:"protein"`
}

func getResponse(name string) {
	var fruit Fruit
	client := &http.Client{}

	request, err := http.NewRequest(http.MethodGet, ENDPOINT+name, nil)
	if err != nil {
		log.Println(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(&fruit)

	result := `== Fruit Data ==
Fruit name: %v
Fruit family: %v
Calories: %v
Fat: %v
Sugar: %v
Carbohydrates: %v
Protein: %v`
	fmt.Printf(result, fruit.Name, fruit.Family, fruit.Nutritions.Calories, fruit.Nutritions.Fat, fruit.Nutritions.Sugar, fruit.Nutritions.Carbohydrates, fruit.Nutritions.Protein)
}

func main() {
	var input string
	fmt.Print("Enter fruit name: ")
	fmt.Scan(&input)

	getResponse(input)
}
