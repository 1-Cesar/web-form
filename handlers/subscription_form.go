package handlers

import (
	"fmt"
	"log"
	"net/http"
	"web-form/controller"
)

func SubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "Erro ao fazer parse no form: %v", err)
			return
		}

		err := controller.Create(r.PostForm.Get("name"), r.PostForm.Get("email"))
		if err != nil {
			log.Fatal(w, "Erro ao salvar os dados: %v", err)
			return
		}
	}

	http.ServeFile(w, r, "handlers/templates/subscription_form.html")
}
