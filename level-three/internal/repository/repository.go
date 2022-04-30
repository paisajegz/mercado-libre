package repository

import (
	db "mercado-level-1/internal/database"
	"sync"
)

type RepositoryXmen struct{}

var repository *RepositoryXmen
var once sync.Once

func (r *RepositoryXmen) GenerateStats() *[]db.Stat {
	database := db.GetBD()
	var stat []db.Stat
	database.Table("dna_registers").Select("is_mutant , count(is_mutant) as count ").Group("is_mutant").Find(&stat)
	return &stat
}

func (r *RepositoryXmen) InsertDna(dna *db.DNARegister) {
	database := db.GetBD()
	database.Create(dna)
}

func GetRepositoryXmen() *RepositoryXmen {
	once.Do(func() {
		repository = new(RepositoryXmen)
	})
	return repository
}
