package internal

import "pokedex/api"

var Pokedex = make(map[string]api.Pokemon)
var CurrLocationAreasUrl = api.PokeApiBaseUrl + api.LocationAreaPath
var Previous string
