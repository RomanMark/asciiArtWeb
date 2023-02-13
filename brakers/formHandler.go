package brakers

import (
	"html/template"
	"log"
	"net/http"
)

type Ascii struct {
	Art string
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the request
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		// If there is an error, redirect to error page
		http.Redirect(w, r, "/error.html", http.StatusSeeOther)
		return
	}
	// Extract the input and style from the form data
	input := r.FormValue("input")
	style := r.FormValue("styles")

	// Convert the input to ascii art with the given style
	out := PrintAscii(input, style)
	// Check if the output is empty
	if out == "" {
		http.Redirect(w, r, "/error.html", http.StatusSeeOther)
		return
	}

	// Create an Ascii struct with the output
	output := Ascii{Art: out}
	// Parse the template file
	tmpl := template.Must(template.ParseFiles("frontend/result.html"))
	// Execute the template with the output and write to the response
	err = tmpl.Execute(w, output)
	if err != nil {
		log.Println(err)
	}
}
