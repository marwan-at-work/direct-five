package directfive

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// Hello there
func Hello() {
	fmt.Println("hello from directfive")
	jwt.GetSigningMethod("")
}
