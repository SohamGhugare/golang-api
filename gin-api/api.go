package main

// Album struct for storing albums
type album struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Artist string `json:"artist"`
  Price float64 `json:"price"`
}

// Creating dummy albums
var albums = []album{
  {
    ID: "1",
    Title: "Blue Train",
    Artist:"John Coltrane",
    Price: 56.99,
  },
  {
    ID: "2",
    Title: "Jeru",
    Artist: "Gerry Mulligan",
    Price: 17.99,
  },
}
