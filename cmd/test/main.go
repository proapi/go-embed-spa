package main

import (
	"fmt"
	"time"
)

type People interface {
	SayHello()
	GetDetails()
}

type Person struct {
	name  string
	age   int
	email string
	city  string
}

func (p Person) SayHello() {
	fmt.Printf("Hi, I am %s, from %s\n", p.name, p.city)
}

func (p Person) GetDetails() {
	fmt.Printf("[Name: %s, Age: %d, City: %s, Email: %s]\n", p.name, p.age, p.city, p.email)
}

type Speaker struct {
	Person
	speaksOn   []string
	pastEvents []string
}

func (s Speaker) GetDetails() {
	s.Person.GetDetails()
	fmt.Println("Speaker talks on following technologies:")
	for _, value := range s.speaksOn {
		fmt.Println(value)
	}
	fmt.Println("Presented on the following conferences:")
	for _, value := range s.pastEvents {
		fmt.Println(value)
	}
}

type Organizer struct {
	Person
	meetups []string
}

func (o Organizer) GetDetails() {
	o.Person.GetDetails()
	fmt.Println("Organizer, conducting following Meetups:")
	for _, value := range o.meetups {
		fmt.Println(value)
	}
}

type Attendee struct {
	Person
	interests []string
}

type Meetup struct {
	location string
	city     string
	date     time.Time
	people   []People
}

func (m *Meetup) MeetupPeople() {
	for _, v := range m.people {
		v.SayHello()
		v.GetDetails()
	}
}

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	p1 := Speaker{Person{"Shiju", 35, "Kochi", "+91-94003372xx"},
		[]string{"Go", "Docker", "Azure", "AWS"},
		[]string{"FOSS", "JSFOO", "MS TechDays"}}
	p2 := Organizer{Person{"Satish", 35, "Pune", "+91-94003372xx"},
		[]string{"Gophercon", "RubyConf"}}
	p3 := Attendee{Person{"Alex", 22, "Bangalore", "+91-94003672xx"},
		[]string{"Go", "Ruby"}}

	meetup := Meetup{
		"Royal Orchid",
		"Bangalore",
		time.Date(2015, time.February, 19, 9, 0, 0, 0, time.UTC),
		[]People{p1, p2, p3},
	}
	meetup.MeetupPeople()

	for i := 0; i < 10; i++ {
		go f(fmt.Sprintf("goroutine %d", i))
		i := i
		go func(msg string) {
			fmt.Println(msg, ":", i)
		}("inside closure")
	}

	time.Sleep(time.Second)
	fmt.Println("done")
}
