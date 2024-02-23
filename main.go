package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

// defining a data structure to store the scraped data
type PokemonProduct struct {
	url, image, name, price string
}

func main() {
	// initializing the slice of structs that will contain the scraped data
	var pokemonProducts []PokemonProduct

	c := colly.NewCollector()
	c.Visit("https://scrapeme.live/shop/")

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		// printing all URLs associated with the a links in the page
		fmt.Printf("%v\n", e.Attr("href"))

		pokemonProduct := PokemonProduct{}

		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")
	})

}
