package main

type CataloguePage struct {

  Total int `json:"total"`
  Items []Item `json:"items"`
}

type Item struct {
  Icon string
  IconLarge string
  Id int
  CatType string `json:"type"` //reserved keyword
  TypeIcon string
  Name string
  Description string
  Current ItemPrice
  Today ItemPrice
  Members bool
}

type ItemPrice struct {
  Trend string
  Price int
}
