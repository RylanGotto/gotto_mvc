package tonic

import "net/http"

func (c *Tonic) SessionLoad(next http.Handler) http.Handler {
	c.InfoLog.Println("SessionLoad called")
	return c.Session.LoadAndSave(next)
}
