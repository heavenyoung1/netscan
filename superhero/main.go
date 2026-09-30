package main 

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
)

const (
	accessToken = "5eab2a01945dd127e9304d2dc2e1a431"
	urlApi = "https://superheroapi.com/api/"
)

type Biography struct {
	FullName string `json:"full-name"`
	AlterEgo string `json:"alter-egos"` // Тут возможно "No alter egos found." Как это обработать?
	Aliases []string `json:"aliases"`
	PlaceOfBirht string `json:"place-of-birth"`
	Publisher string `json:"publisher"`
	Alignment string `json:"alignment"`
}

type Hero struct {
	Id string `json:"id"`
	Name string `json:"name"`
}


type LookedHero struct {
	Response string `json:"response"`
	Results []Hero `json:"results"`
}

type Appearance struct {
	Gender string `json:"gender"`
	Race string `json:"race"`
	Height []string `json:"height"`
	Weight []string  `json:"weight"`
}


func searchHero(name string) (Hero, error) {
	// Найти героя по его имени
	urlSearchHero := fmt.Sprintf("%s/%s/search/%s", urlApi, accessToken, name)

	// 1. Выполняем GET-запрос
	response, err := http.Get(urlSearchHero)
	if err != nil {
		return Hero{}, fmt.Errorf("Ошибка запроса: %v", err)
	}
	defer response.Body.Close()

	var data LookedHero
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return Hero{}, fmt.Errorf("Ошибка декодирования: %v", err)
	}
	//fmt.Printf("%+v", hero)

	return data.Results[0], nil
}

func getBiographySuperHero(id string) (Biography, error) {
	urlBiography := fmt.Sprintf("%s/%s/%s/biography", urlApi, accessToken, id)

	// 1. Выполняем GET-запрос
	response, err := http.Get(urlBiography)
	if err != nil {
		return Biography{}, fmt.Errorf("Ошибка запроса: %v", err)
	}
	defer response.Body.Close()

	var bio Biography
	err = json.NewDecoder(response.Body).Decode(&bio)
	if err != nil {
		return Biography{}, fmt.Errorf("Ошибка декодирования: %v", err)
	}
	// fmt.Printf("%+v\n", bio)
	return bio, nil

}

func getAppearance(id string) (Appearance, error) {
	urlAppearance := fmt.Sprintf("%s/%s/%s/appearance", urlApi, accessToken, id)

	// 1. Выполняем GET-запрос
	response, err := http.Get(urlAppearance)
	if err != nil {
		return Appearance{}, fmt.Errorf("Ошибка запроса: %v", err)
	}
	defer response.Body.Close()

	var appearance Appearance
	if err := json.NewDecoder(response.Body).Decode(&appearance); err != nil {
		return Appearance{}, fmt.Errorf("Ошибка декодирования: %v", err)
	}
	return appearance, nil
}


func main() {
	hero, err := searchHero("batman")

	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	fmt.Printf("%+v\n", hero)
	ID := hero.Id

	bio, err := getBiographySuperHero(ID)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	fmt.Printf("%+v\n", bio)

	appearance, error := getAppearance(ID)
	if error != nil {
		log.Fatalf("Ошибка: %v", error)
	}
	fmt.Printf("%+v\n", appearance)
}
	



