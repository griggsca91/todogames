package controllers

import (
  "aahframework.org/aah.v0"
  "todogames/app/models"
)

// App struct application controller
type App struct {
  *aah.Context
}

// Index method is application home page.
func (a *App) Index() {
  data := aah.Data{
    "Greet": models.Greet{
      Message: "Welcome to aah framework - Web Application",
    },
  }

  a.Reply().Ok().HTML(data)
}
