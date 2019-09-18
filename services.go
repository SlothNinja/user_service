package main

import (
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/SlothNinja/log"
	"github.com/SlothNinja/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/errors"
	"golang.org/x/exp/errors/fmt"
)

const (
	userNewPath  = "/#/new"
	userShowPath = "/#/show/"
	enterMsg     = "Entering"
	exitMsg      = "Exiting"
	msgKey       = "msg"
	cuKey        = "cu"
	uKey         = "u"
)

var (
	errValidation   = errors.New("validation error")
	errUnexpected   = errors.New("unexpected error")
	errHaveAccount  = fmt.Errorf("you already have an account: %w", errValidation)
	errMissingToken = errors.New("missing session token")
)

func jerr(c *gin.Context, err error) {
	switch {
	case errors.Is(err, errValidation):
		c.JSON(http.StatusOK, gin.H{msgKey: err.Error()})
	default:
		log.Debugf(err.Error())
		c.JSON(http.StatusOK, gin.H{msgKey: errUnexpected.Error()})
	}
}

func newAction(prefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debugf(enterMsg)
		defer log.Debugf(exitMsg)

		cu := user.Current(c)
		if cu != user.None {
			jerr(c, errHaveAccount)
			return
		}

		s := sessions.Default(c)
		token, ok := user.SessionTokenFrom(s)
		if !ok {
			jerr(c, errMissingToken)
			return
		}

		u := user.New(token.ID)
		u.Name = user.NameFrom(token.Email)
		u.Email = token.Email
		u.EmailReminders = true
		u.EmailNotifications = true
		u.GravType = user.GTMonster
		c.JSON(http.StatusOK, gin.H{uKey: u})
	}
}

func current(c *gin.Context) {
	log.Debugf(enterMsg)
	defer log.Debugf(exitMsg)

	cu := user.Current(c)
	if cu == user.None {
		jerr(c, user.ErrNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{cuKey: cu})
}

func json(prefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debugf(enterMsg)
		defer log.Debugf(exitMsg)

		id := c.Param("id")

		client, err := user.Client(c)
		if err != nil {
			jerr(c, fmt.Errorf("%s: %w", err, user.ErrClient))
			return
		}

		u := user.New(id)
		err = client.Get(c, u.Key, &u)
		if err != nil {
			jerr(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{uKey: u})
	}
}

func create(prefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debugf(enterMsg)
		defer log.Debugf(exitMsg)

		cu := user.Current(c)
		if cu != user.None {
			jerr(c, errHaveAccount)
			return
		}

		session := sessions.Default(c)
		token, ok := user.SessionTokenFrom(session)
		if !ok {
			jerr(c, errMissingToken)
			return
		}

		client, err := user.Client(c)
		if err != nil {
			jerr(c, fmt.Errorf("%s: %w", err, user.ErrClient))
			return
		}

		u := user.New(token.ID)
		err = client.Get(c, u.Key, &u)
		if err == nil {
			jerr(c, errHaveAccount)
			return
		}

		if err != datastore.ErrNoSuchEntity {
			jerr(c, err)
			return
		}

		u, err = fromJSON(c, token.ID, token.Email)
		if err != nil {
			jerr(c, err)
			return
		}

		t := time.Now()
		u.UpdatedAt, u.CreatedAt = t, t
		if u.JoinedAt.IsZero() {
			u.JoinedAt = t
		}
		_, err = client.Put(c, u.Key, &u)
		if err != nil {
			jerr(c, err)
			return
		}

		log.Debugf("put user: %#v", u)

		err = u.SaveTo(session)
		if err != nil {
			jerr(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{uKey: u})
	}
}

func fromJSON(c *gin.Context, id, email string) (user.User, error) {
	log.Debugf(enterMsg)
	defer log.Debugf(exitMsg)

	jData := struct {
		Name               string `json:"name"`
		EmailNotifications bool   `json:"emailnotifications"`
		GravType           string `json:"gravtype"`
	}{}

	err := c.ShouldBindJSON(&jData)
	if err != nil {
		return user.User{}, err
	}

	u := user.New(id)
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
		log.Debugf(enterMsg)
		defer log.Debugf(exitMsg)

		cu := user.Current(c)
		if cu == user.None {
			jerr(c, fmt.Errorf("must be logged-in to edit account: %w", errValidation))
			return
		}

		uid := c.Param("uid")
		if uid != cu.ID() && !cu.Admin {
			jerr(c, fmt.Errorf("you can only edit your own account: %w", errValidation))
			return
		}

		client, err := user.Client(c)
		if err != nil {
			jerr(c, fmt.Errorf("%s: %w", err, user.ErrClient))
			return
		}

		u := user.New(uid)
		err = client.Get(c, u.Key, &u)
		if err != nil {
			jerr(c, err)
			return
		}

		u2, err := fromJSON(c, u.ID(), u.Email)
		if err != nil {
			jerr(c, err)
			return
		}

		if u2.Name != u.Name {
			uniq, err := uniqueName(c, u2)
			if err != nil {
				jerr(c, err)
				return
			}

			if !uniq {
				jerr(c, fmt.Errorf("screen name: %s already in use: %w", u2.Name, errValidation))
				return
			}
		}

		_, err = client.Put(c, u2.Key, &u2)
		if err != nil {
			jerr(c, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{cuKey: u2})
	}
}

func uniqueName(c *gin.Context, u1 user.User) (bool, error) {
	log.Debugf("Entering")
	defer log.Debugf("Exiting")

	client, err := user.Client(c)
	if err != nil {
		return false, err
	}

	var us []user.User
	ks, err := client.GetAll(c, datastore.NewQuery(user.Kind).Filter("LCName =", u1.LCName), &us)
	if err != nil {
		return false, err
	}

	switch l := len(ks); l {
	case 0:
		return true, nil
	case 1:
		return u1.LCName == us[0].LCName && u1.ID() != us[0].ID(), nil
	default:
		return false, nil
	}
}
