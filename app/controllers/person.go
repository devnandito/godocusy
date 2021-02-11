package controllers

import (
    "fmt"
    models "docusys/app/models"
)

func getPerson(personPK int) (models.Person, error) {
    res := models.Person{}
    var pk int
    var first_name string
    err := DB.QueryRow(`SELECT pk, first_name FROM person WHERE pk = $1`, personPK).Scan(&pk, &first_name)
    if err != nil {
        res = models.Person{PK: pk, First_name: first_name}
    }
    return res, err
}

func allPerson() ([]models.Person, error) {
    people := []models.Person{}
    person := models.Person{}
    rows, err := DB.Query(`SELECT pk, first_name FROM person ORDER BY pk`)
    defer rows.Close()
    if err == nil {
        for rows.Next() {
            var pk int
            var first_name string
            err = rows.Scan(&pk, &first_name)
            if err != nil {
                panic(err.Error())
            }
            person.PK = pk
            person.First_name = first_name
            people = append(people, person)
            return people, err
        }
    }
    return people, err
}

func insertPerson(first_name string) (int, error) {
    var personPK int
    err := DB.QueryRow(`INSERT INTO person(first_name) VALUES($1) RETURNING pk`, first_name).Scan(&personPK)
    if err != nil {
        return 0, err
    }
    fmt.Printf("Last inserted PK: %v\n", personPK)
    return personPK, err
}
