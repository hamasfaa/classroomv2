package main

import (
	"be/entities"
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
)

func main() {
	stmts, err := gormschema.New("mysql").Load(&entities.AbsenMahasiswa{}, &entities.UserKelas{}, &entities.User{}, &entities.Kelas{}, &entities.AbsenDosen{}, &entities.TugasDosen{}, &entities.TugasMahasiswa{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
