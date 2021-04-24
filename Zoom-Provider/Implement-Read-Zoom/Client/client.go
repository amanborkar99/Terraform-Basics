package client

import (
	
	"errors"
)

const URL string= "https://api.zoom.us/v2/users"

type Client struct{
	Token string
}


func NewClient( token string) (*Client, error) {
	c := Client{
		Token: token,
	}

	if (token != ""){
		
		return &c, nil
		
		}
		
	return nil, errors.New("invalid token") 
}

		




