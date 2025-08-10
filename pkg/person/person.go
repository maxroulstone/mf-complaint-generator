package person

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type FakePerson struct {
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	Address     string
	City        string
	PostCode    string
	DateOfBirth string
}

func Generate() FakePerson {
	firstNames := []string{"James", "Sarah", "Michael", "Emma", "David", "Jessica", "Robert", "Lisa", "John", "Amanda", "Chris", "Laura", "Mark", "Helen", "Paul", "Karen", "Steve", "Rachel", "Tom", "Sophie"}
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis", "Rodriguez", "Martinez", "Taylor", "Anderson", "Thomas", "Jackson", "White", "Harris", "Martin", "Thompson", "Wilson", "Moore"}
	streets := []string{"Oak Street", "Main Road", "High Street", "Church Lane", "Mill Close", "Victoria Road", "King Street", "Queen Avenue", "Park Lane", "Station Road", "The Green", "Elm Drive", "Maple Close", "Cedar Avenue"}
	cities := []string{"Manchester", "Birmingham", "Liverpool", "Leeds", "Sheffield", "Bristol", "Cardiff", "Edinburgh", "Glasgow", "Newcastle", "Nottingham", "Southampton", "Portsmouth", "Brighton"}
	
	firstName := firstNames[randomInt(len(firstNames))]
	lastName := lastNames[randomInt(len(lastNames))]
	email := fmt.Sprintf("%s.%s@email.com", firstName, lastName)
	
	return FakePerson{
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		Phone:       fmt.Sprintf("07%d%d%d %d%d%d%d%d%d", randomInt(10), randomInt(10), randomInt(10), randomInt(10), randomInt(10), randomInt(10), randomInt(10), randomInt(10), randomInt(10)),
		Address:     fmt.Sprintf("%d %s", randomInt(200)+1, streets[randomInt(len(streets))]),
		City:        cities[randomInt(len(cities))],
		PostCode:    fmt.Sprintf("%s%d %d%s%s", string(rune('A'+randomInt(26))), randomInt(10), randomInt(10), string(rune('A'+randomInt(26))), string(rune('A'+randomInt(26)))),
		DateOfBirth: fmt.Sprintf("%02d/%02d/%d", randomInt(28)+1, randomInt(12)+1, 1960+randomInt(40)),
	}
}

func (p FakePerson) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func (p FakePerson) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s", p.Address, p.City, p.PostCode)
}

func randomInt(max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(n.Int64())
}
