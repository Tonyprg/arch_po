package main

import (
    "html/template"
    "net/http"
    "os"
    "database/sql"
    "time"
    _ "github.com/lib/pq"
)

type Client struct {
    id int
    name string
    last_name string
    surname string
    birthdate time.Time
    issue_plase string
    phone_number int
}

type Agent struct {
    id int
    serial_number int
}

type Service struct {
    id int
    name string
    description string
    cost int
    compensation int
}

type Bid struct {
    id int
    status string
    client_id int
    agent_id int
    service_id int
}

type DataBaseEntry interface {
    insert()
    delete()
}

func (self Client) insert () {
    db, _ := sql.Open("postgres", "user=tony password='1' sslmode=disable");
    db.Exec(`INSERT INTO client (name, last_name, surname, birthdate, issue_plase, phone_number)
             VALUES ($1, $2, $3, $4, $5, $6); `,
        self.name, self.last_name, self.surname,
        self.birthdate, self.issue_plase,
        self.phone_number)
}

func (self Client) delete () {
    db, _ := sql.Open("postgres", "user=tony password='1' sslmode=disable");
    db.Exec(`
    DELETE FROM client WHERE phone_number = $1;
    `, self.phone_number)
}

func (self Agent) insert() {
    db, _ := sql.Open("postgres", "user=tony password='1' sslmode=disable");
    db.Exec(`INSERT INTO agent (serial_number)
             VALUES ($1);`, self.serial_number)
}

func (self Agent) delete() {
    db, _ := sql.Open("postgres", "user=tony password='1' sslmode=disable");
    db.Exec("DELETE FROM agent WHERE serial_number = $1;", self.serial_number)
}

func (self Service) insert() {
    db, _ := sql.Open("postgres", "user=tony password='1' sslmode=disable")
    db.Exec(`INSERT INTO service (name, description, cost, compensation)
             VALUES ($1, $2, $3, $4); `,
    self.name, self.description, self.cost, self.compensation)
}

func (self Service) delete () {
    db, _ := sql.Open("postgres", "user=tony password='1' sslmode=disable")
    db.Exec("DELETE FROM service WHERE name = $1", self.name)
}

func (self Bid) insert () {
    db, _ := sql.Open("postgres", "user=tony password='1' sslmode=disable")
    db.Exec(`INSERT INTO bid (status, client_id, agent_id, service_id)
             VALUES ($1, $2, $3, $4);`,
        self.status, self.client_id, self.agent_id, self.service_id)
}

func (self Bid) delete () {
    db, _ := sql.Open("postgres", "user=tony password='1' sslmode=disable")
    db.Exec("DELETE FROM bid WHERE id = $1", self.id)
}

func main () {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fs := http.FileServer(http.Dir("./ui/static"))
    mux := http.NewServeMux()
    mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        ts, _ := template.ParseFiles("./ui/html/index.html")
        ts.Execute(w, nil)
    })
    mux.Handle("/ui/static/", http.StripPrefix("/ui/static/", fs))
    
    mux.HandleFunc("/bingo", func (w http.ResponseWriter, r *http.Request){
        w.Write([]byte("{\"name\" : \"tony\"}"))
    })

    http.ListenAndServe(":" + port, mux)
}
