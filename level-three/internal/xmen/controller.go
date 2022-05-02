package xmen

import (
	"encoding/json"
	db "mercado-level-1/internal/database"
	"mercado-level-1/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func VerifyMutant(c echo.Context) error {
	dna := c.Get("dna").(*DNA)
	byteDna, err := json.Marshal(dna)
	if err != nil {
		return c.NoContent(http.StatusForbidden)
	}

	if IsMutant(dna) {
		go repository.GetRepositoryXmen().InsertDna(&db.DNARegister{
			IsMutant: true,
			Dna:      string(byteDna),
		})
		return c.NoContent(http.StatusOK)
	}
	go repository.GetRepositoryXmen().InsertDna(&db.DNARegister{
		IsMutant: false,
		Dna:      string(byteDna),
	})
	return c.NoContent(http.StatusForbidden)

}

func Stats(c echo.Context) error {
	stat := repository.GetRepositoryXmen().GenerateStats()
	response := new(ResponseStat)
	for _, st := range *stat {
		if st.IsMutant {
			response.CountNutant = st.Count
			continue
		}
		response.CountHuman = st.Count
	}
	response.Ratio = float64(response.CountNutant) / float64(response.CountHuman)
	return c.JSON(http.StatusOK, response)

}

func IsMutant(dna *DNA) bool {
	for y, cadena := range dna.DNA {
		for x, item := range cadena {
			if VerifyDna(dna, string(item), x, y) {
				return true
			}

		}
	}
	return false
}
