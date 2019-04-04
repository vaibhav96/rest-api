package main


type Post struct{
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Role      string  `json:"role"`
 }

type Posts []Post