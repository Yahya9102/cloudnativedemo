package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

/*
 {
    "id": "0",
    "author": "Alejandro Escamilla",
    "width": 5000,
    "height": 3333,
    "url": "https://unsplash.com/photos/yC-Yzbqy7PY",
    "download_url": "https://picsum.photos/id/0/5000/3333"
  }

*/

// Steg 1 skapa ett struct som motsvarar ett object i APIet
type PicsumImage struct {
	ID string `json:"id"`
	Author string `json:"author"`
	Width int `json:"width"`
	Height int `json:"height"`
	URL string `json:"url"`
	DownloadUrl string `json:"download_url"` 
}



func main (){

	// Bygger vår http client med Gos inbygda biblotek 
	client := &http.Client{
		Timeout: 10 * time.Second,
	}


	//  Bygger vår URL som vi ska göra anrop till
	url := "https://picsum.photos/v2/list?page=1&limit=5"


	// Gör anropet
	resp, err := client.Get(url)
	if err != nil {
		panic(fmt.Errorf("kunde inte nå picsum api: %w", err)) // Stäng om det inte gick att nå APIet
	}


	// Stänger anslutningen när vi är klara
	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK{
		panic (fmt.Errorf("picsum svarade med errorkod %d", resp.StatusCode))
	}


	// Steg 2 avkoda json svaret till vår struct
	var images []PicsumImage

	if err := json.NewDecoder(resp.Body).Decode(&images); err != nil {
		panic(fmt.Errorf("kunde inte avkoda json: %w", err))
	}

	fmt.Printf("hämtade %d bilder från picsum: \n ", len(images))


	// Skriv ut bilderna vi fick tillbaka
	for i, img := range images {
		fmt.Printf("%d) id=%s author=%s download_url=%s\n", i, img.ID, img.Author, img.DownloadUrl)
	}












}



