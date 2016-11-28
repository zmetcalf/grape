package handlers

import (
	"fmt"
	"net/http"
	"github.com/zmetcalf/grape/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %d!", models.GetAll())
}
