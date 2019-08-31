package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"bitbucket.org/SlothNinja/log"
	"bitbucket.org/SlothNinja/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"cloud.google.com/go/datastore"
)

const (
	userNewPath  = "/#/new"
	userShowPath = "/#/show/"
)

func newAction(prefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debugf("Entering")
		defer log.Debugf("Exiting")

		_, found := user.Current(c)
		if !found {
			jsonMsg(c, "You already have an account.")
			return
		}

		s := sessions.Default(c)
		token, ok := user.SessionTokenFrom(s)
		if !ok {
			log.Errorf("Missing SessionToken")
			jsonMsg(c, "Unexpected error. Try again.")
			return
		}

		log.Debugf("token: %#v", token)
		u := user.New2(token.ID)
		u.Name = user.Name(token.Email)
		u.Email = token.Email
		u.EmailReminders = true
		u.EmailNotifications = true
		u.GravType = user.GTMonster
		c.JSON(http.StatusOK, gin.H{"u": u})
	}
}

func current(c *gin.Context) {
	log.Debugf("Entering")
	defer log.Debugf("Exiting")

	cu, found := user.Current(c)
	if !found {
		jsonMsg(c, "unable to find current user")
		return
	}
	c.JSON(http.StatusOK, gin.H{"cu": cu})
}

func json(prefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debugf("Entering")
		defer log.Debugf("Exiting")

		id := c.Param("id")
		u, err := user.ByID(c, id)
		if err != nil {
			log.Warningf("ByID err: %s", err)
			c.Redirect(http.StatusSeeOther, homePath)
			return
		}

		c.JSON(http.StatusOK, gin.H{"u": u})
	}
}

func create(prefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debugf("Entering")
		defer log.Debugf("Exiting")

		_, found := user.Current(c)
		if !found {
			jsonMsg(c, "must be logged-in to create account.")
			return
		}

		session := sessions.Default(c)
		token, ok := user.SessionTokenFrom(session)
		if !ok {
			log.Errorf("Missing SessionToken")
			jsonMsg(c, "Must be logged-in to create account.")
			return
		}

		u, err := user.ByID(c, token.ID)
		if err == nil {
			jsonMsg(c, "Account already exists for id: %s", token.ID)
			return
		}

		if err != datastore.ErrNoSuchEntity {
			jsonMsg(c, "Unexpected error. Try again.")
			return
		}

		u, err = fromJSON(c, token.ID, token.Email)
		if err != nil {
			log.Errorf(err.Error())
			jsonMsg(c, "Unexpected error. Try again.")
			return
		}

		id0, id1, gid, joinedAt, found, err := user.ByOldEntries(c, u.Email)
		if err != nil {
			log.Errorf("user.ByOldEntries error: %s", err)
			jsonMsg(c, "Unexpected error. Try again.")
			return
		}

		log.Debugf("found: %#v", found)
		if found {
			u.User0ID = id0
			u.User1ID = id1
			u.GoogleID = gid
			u.JoinedAt = joinedAt
		}

		client, err := user.Client(c)
		if err != nil {
			log.Errorf(err.Error())
			jsonMsg(c, "Unexpected error. Try again.")
			return
		}

		t := time.Now()
		u.UpdatedAt, u.CreatedAt = t, t
		if u.JoinedAt.IsZero() {
			u.JoinedAt = t
		}
		_, err = client.Put(c, u.Key, u)
		if err != nil {
			log.Errorf("datastore.Put error: %v", err)
			jsonMsg(c, "Unexpected error. Try again.")
			return
		}

		log.Debugf("put user: %#v", u)

		err = u.To(session)
		if err != nil {
			log.Errorf("session.Save error: %v", err)
		}

		c.JSON(http.StatusOK, gin.H{"u": u})
	}
}

func fromJSON(c *gin.Context, id, email string) (user.User2, error) {
	log.Debugf("Entering")
	defer log.Debugf("Exiting")

	jData := struct {
		Name               string `json:"name"`
		EmailNotifications bool   `json:"emailnotifications"`
		GravType           string `json:"gravtype"`
	}{}

	err := c.ShouldBindJSON(&jData)
	if err != nil {
		return user.User2{}, err
	}

	u := user.New2(id)
	u.Name = strings.TrimSpace(jData.Name)
	u.LCName = strings.ToLower(jData.Name)
	u.Email = email
	u.EmailNotifications = jData.EmailNotifications
	u.EmailReminders = true
	u.GravType = jData.GravType
	return u, nil
}

func showPath(prefix string, uid string) string {
	return fmt.Sprint(prefix, userShowPath, uid)
}

func update(prefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debugf("Entering")
		defer log.Debugf("Exiting")

		cu, found := user.Current(c)
		if !found {
			jsonMsg(c, "Must be logged-in to edit account.")
			return
		}

		uid := c.Param("uid")
		if uid != cu.ID() && !cu.Admin {
			jsonMsg(c, "You can only edit your own account.")
			return
		}

		u, err := user.ByID(c, uid)
		if err != nil {
			log.Errorf(err.Error())
			jsonMsg(c, "Unexpected error. Try again.")
			return
		}

		u2, err := fromJSON(c, u.ID(), u.Email)
		if err != nil {
			log.Errorf(err.Error())
			jsonMsg(c, "Unexpected error. Try again.")
			return
		}

		uniq, err := uniqueName(c, u2)
		if err != nil {
			log.Errorf(err.Error())
			jsonMsg(c, "Unexpected error. Try again.")
			return
		}

		if !uniq {
			jsonMsg(c, "Screen Name: %s already in use.", u2.Name)
			return
		}

		client, err := user.Client(c)
		if err != nil {
			log.Errorf(err.Error())
			return
		}

		_, err = client.Put(c, u2.Key, u2)
		if err != nil {
			log.Errorf(err.Error())
			// route := fmt.Sprintf("/user/show/%s", c.Param("uid"))
			// c.Redirect(http.StatusSeeOther, route)
			return
		}

		c.JSON(http.StatusOK, gin.H{"cu": u2})
	}
}

func jsonMsg(c *gin.Context, format string, args ...interface{}) {
	c.JSON(http.StatusOK, struct {
		Msg string `json:"msg"`
	}{
		Msg: fmt.Sprintf(format, args...),
	})
}

func uniqueName(c *gin.Context, u1 user.User2) (bool, error) {
	u2, err := user.ByLCName(c, u1.LCName)
	if err != nil {
		return false, err
	}

	if u1.LCName == u2.LCName && u1.ID() != u2.ID() {
		return false, nil
	}
	return true, nil
}
