package seed

import (
	"errors"
	"github.com/samithiwat/samithiwat-backend/src/database"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Seed struct {
	db database.Database
}

func seed(s Seed, seedMethodName string) error {
	m := reflect.ValueOf(s).MethodByName(seedMethodName)

	if !m.IsValid() {
		return errors.New("Invalid seed")
	}

	log.Println("Seeding", seedMethodName, "...")

	m.Call(nil)

	log.Println("Seed", seedMethodName, "succeed")

	return nil
}

func Execute(db database.Database, seedMethodNames ...string) {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	// Execute all
	if len(seedMethodNames) == 0 {
		seeds := []reflect.Method{}

		log.Println("Running all seeder...")
		for i := 0; i < seedType.NumMethod(); i++ {
			seeds = append(seeds, seedType.Method(i))

		}

		sort.Slice(seeds, func(p, q int) bool {
			name1 := strings.Split(seeds[p].Name, "_")
			name2 := strings.Split(seeds[q].Name, "_")

			name1Timestamp, err := strconv.Atoi(name1[1])
			if err != nil {
				log.Fatalln(err)
			}

			name2Timestamp, err := strconv.Atoi(name2[1])
			if err != nil {
				log.Fatalln(err)
			}
			return name1Timestamp < name2Timestamp
		})

		for _, method := range seeds {
			seed(s, method.Name)
		}
	}

	// Execute only the given names
	for _, item := range seedMethodNames {
		seed(s, item)
	}
}
