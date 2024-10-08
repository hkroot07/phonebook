package book

import (
	"fmt"
	"time"
)

type PhoneBook map[string]PhoneNumber

type PhoneNumber struct {
	Number       string
	LastUpdateAt int64
}

func (book *PhoneBook) Add(name, phoneNum string) error {
	if _, exists := (*book)[name]; exists {
		return fmt.Errorf("name %s already exists", name)
	}

	// n := new(PhoneNumber)
	// n.Number = phoneNum
	// n.LastUpdateAt = time.Now().Unix()
	// (*book)[name] = *n
	(*book)[name] = PhoneNumber{
		Number:       phoneNum,
		LastUpdateAt: time.Now().Unix(),
	}
	return nil
}

func (book *PhoneBook) Get(name string) (PhoneNumber, error) {
	if numberData, exists := (*book)[name]; exists {
		return numberData, nil
	}

	return PhoneNumber{}, fmt.Errorf("no entry found for %s", name)
}

func (book *PhoneBook) Update(name, newPhoneNum string) error {
	if _, exists := (*book)[name]; !exists {
		return fmt.Errorf("name %s does not exist", name)
	}

	(*book)[name] = PhoneNumber{
		Number:       newPhoneNum,
		LastUpdateAt: time.Now().Unix(),
	}
	return nil
}

func (book *PhoneBook) Delete(name string) error {
	if _, exists := (*book)[name]; !exists {
		return fmt.Errorf("name %s does not exist", name)
	}
	delete(*book, name)
	return nil
}
