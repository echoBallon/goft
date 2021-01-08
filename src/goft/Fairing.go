package goft

import "github.com/gin-gonic/gin"

/**
middleware interface
*/
type Fairing interface {
	OnRequest(*gin.Context) error
}
