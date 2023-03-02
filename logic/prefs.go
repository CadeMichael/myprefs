package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Pref struct {
	Name     string
	Rating   int
	Kind     string
	Location string
}

type Prefs []Pref

func (p *Prefs) New(name string, rating int, kind string, loc string) {

	pref := Pref{
		Name:     name,
		Rating:   rating,
		Kind:     kind,
		Location: loc,
	}

	*p = append(*p, pref)
}

func (p *Prefs) Read() error {
	jsonByte, err := ioutil.ReadFile("prefs.json")
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = json.Unmarshal(jsonByte, p)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(p)
	return nil
}

func (p *Prefs)SavePrefs() error {
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(string(data))
	ioutil.WriteFile("prefs.json", data, 0644)
	return nil
}
