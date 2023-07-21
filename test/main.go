package main

//
//import (
//	"context"
//	"net/http"
//
//	"entgo.io/ent/dialect"
//	"entgo.io/ent/entc/gen"
//	"entgo.io/ent/entc/integration/enttest"
//	"github.com/labstack/echo/v4"
//	"github.com/labstack/echo/v4/middleware"
//	_ "github.com/mattn/go-sqlite3"
//)
//
//type Candidate struct {
//	Name  string `json:"name"`
//	Votes int    `json:"votes"`
//}
//
//type Poll struct {
//	ID         int         `json:"id"`
//	Title      string      `json:"title"`
//	Candidates []Candidate `json:"candidates"`
//}
//
//var client *gen.Client
//
//func main() {
//	client = enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
//	defer client.Close()
//
//	e := echo.New()
//	e.Use(middleware.Logger())
//	e.Use(middleware.Recover())
//
//	e.POST("/polls", createPoll)
//	e.GET("/polls/:id", getPoll)
//	e.PUT("/polls/:id", vote)
//
//	e.Logger.Fatal(e.Start(":1323"))
//}
//
//func createPoll(c echo.Context) error {
//	p := new(Poll)
//	if err := c.Bind(p); err != nil {
//		return err
//	}
//
//	poll, err := client.Poll.
//		Create().
//		SetTitle(p.Title).
//		Save(context.Background())
//	if err != nil {
//		return err
//	}
//
//	for _, candidate := range p.Candidates {
//		_, err := poll.Update().
//			AddCandidateIDs(
//				client.Candidate.Create().
//					SetName(candidate.Name).
//					SetVotes(candidate.Votes).
//					SaveX(context.Background()).ID,
//			).
//			Save(context.Background())
//		if err != nil {
//			return err
//		}
//	}
//
//	return c.JSON(http.StatusCreated, poll)
//}
//
//func getPoll(c echo.Context) error {
//	id := c.Param("id")
//
//	poll, err := client.Poll.
//		Query().
//		Where(gen.PollID(id)).
//		QueryCandidates().
//		All(context.Background())
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, poll)
//}
//
//func vote(c echo.Context) error {
//	id := c.Param("id")
//	candidateName := c.QueryParam("candidate")
//
//	candidate, err := client.Candidate.
//		Query().
//		Where(
//			gen.CandidateName(candidateName),
//			gen.CandidateHasPollsWith(gen.PollID(id)),
//		).
//		Only(context.Background())
//	if err != nil {
//		return err
//	}
//
//	_, err = candidate.Update().
//		AddVotes(1).
//		Save(context.Background())
//	if err != nil {
//		return err
//	}
//
//	return c.NoContent(http.StatusOK)
//}
