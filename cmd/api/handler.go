package main

import "net/http"

func (app *application) SubToNewsletter(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var n Newsletter

	err := app.ReadJSON(w, r, &n)
	if err != nil {
		app.BadRequest(w, r, err)
		return
	}

	err = app.AddNewsletter(n)
	if err != nil {
		app.BadRequest(w, r, err)
		return
	}

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"All good"`
	}

	app.WriteJSON(w, r, res, http.StatusCreated)

}

func (app *application) OrderENDPOINT(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var o Order

	err := app.ReadJSON(w, r, &o)
	if err != nil {
		app.BadRequest(w, r, err)
		return
	}

	st, err := app.GetOrderStr(o)
	if err != nil {
		app.BadRequest(w, r, err)
		return
	}

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
		Order   string `json:"order"`
	}
	res.Error = false
	res.Message = "All Good"
	res.Order = st

	app.WriteJSON(w, r, res, http.StatusOK)

}

