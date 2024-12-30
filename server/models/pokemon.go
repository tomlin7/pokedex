package models

type Pokemon struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Number          int    `json:"number"`
	ImageURL        string `json:"imageURL"`
	Description     string `json:"description"`
	BackgroundColor string `json:"backgroundColor"`
}

var PokemonList = []Pokemon{
	{
		ID:              1,
		Name:            "Bulbasaur",
		Number:          1,
		ImageURL:        "https://assets.pokemon.com/assets/cms2/img/pokedex/full/001.png",
		Description:     "Bulbasaur can be seen napping in bright sunlight. There is a seed on its back. By soaking up the sun's rays, the seed grows progressively larger.",
		BackgroundColor: "#E8F3E8",
	},
	{
		ID:              2,
		Name:            "Ivysaur",
		Number:          2,
		ImageURL:        "https://assets.pokemon.com/assets/cms2/img/pokedex/full/002.png",
		Description:     "There is a bud on this Pokémon's back. To support its weight, Ivysaur's legs and trunk grow thick and strong. If it starts spending more time lying in the sunlight, it's a sign that the bud will bloom into a large flower soon.",
		BackgroundColor: "#E8F3E8",
	},
	{
		ID:              3,
		Name:            "Venusaur",
		Number:          3,
		ImageURL:        "https://assets.pokemon.com/assets/cms2/img/pokedex/full/003.png",
		Description:     "There is a large flower on Venusaur's back. The flower is said to take on vivid colors if it gets plenty of nutrition and sunlight. The flower's aroma soothes the emotions of people.",
		BackgroundColor: "#E8F3E8",
	},
	{
		ID:              4,
		Name:            "Charmander",
		Number:          4,
		ImageURL:        "https://assets.pokemon.com/assets/cms2/img/pokedex/full/004.png",
		Description:     "The flame at the tip of Charmander's tail makes a sound as it burns. You can only hear it in quiet places.",
		BackgroundColor: "#F8D1D1",
	},
	{
		ID:              5,
		Name:            "Charmeleon",
		Number:          5,
		ImageURL:        "https://assets.pokemon.com/assets/cms2/img/pokedex/full/005.png",
		Description:     "Charmeleon mercilessly destroys its foes using its sharp claws. If it encounters a strong foe, it turns aggressive. In this excited state, the flame at the tip of its tail flares with a bluish white color.",
		BackgroundColor: "#F8D1D1",
	},
	{
		ID:              6,
		Name:            "Charizard",
		Number:          6,
		ImageURL:        "https://assets.pokemon.com/assets/cms2/img/pokedex/full/006.png",
		Description:     "Charizard flies around the sky in search of powerful opponents. It breathes fire of such great heat that it melts anything. However, it never turns its fiery breath on any opponent weaker than itself.",
		BackgroundColor: "#F8D1D1",
	},
	{
		ID:              7,
		Name:            "Squirtle",
		Number:          7,
		ImageURL:        "https://assets.pokemon.com/assets/cms2/img/pokedex/full/007.png",
		Description:     "Squirtle's shell is not merely used for protection. The shell's rounded shape and the grooves on its surface help minimize resistance in water, enabling this Pokémon to swim at high speeds.",
		BackgroundColor: "#D1E8F3",
	},
	{
		ID:              8,
		Name:            "Wartortle",
		Number:          8,
		ImageURL:        "https://assets.pokemon.com/assets/cms2/img/pokedex/full/008.png",
		Description:     "Its long, furry tail is a symbol of longevity, making it quite popular among older people.",
		BackgroundColor: "#D1E8F3",
	},
	{
		ID:              9,
		Name:            "Blastoise",
		Number:          9,
		ImageURL:        "https://assets.pokemon.com/assets/cms2/img/pokedex/full/009.png",
		Description:     "Blastoise has water spouts that protrude from its shell. The water spouts are very accurate. They can shoot bullets of water with enough accuracy to strike empty cans from a distance of over 160 feet.",
		BackgroundColor: "#D1E8F3",
	},
}
