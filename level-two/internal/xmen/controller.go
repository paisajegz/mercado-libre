package xmen

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func VerifyMutant(c echo.Context) error {
	dna := c.Get("dna").(*DNA)

	if IsMutant(dna) {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusForbidden)

}

func IsMutant(dna *DNA) bool {
	for y, cadena := range dna.DNA {
		for x, item := range cadena {
			if verifyDna(dna, string(item), x, y) {
				return true
			}

		}
	}
	return false
}
