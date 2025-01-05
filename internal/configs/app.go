package configs

type RunApp struct {
	Http func() error
	Cors string
	Log  func() error
}

func (app *RunApp) Start() error {
	app.Log.Println("Starting application...")

	if app.Cors != "" {
		app.Log.Printf("CORS configured: %s\n", app.Cors)
	}

	if app.Http != nil {
		app.Log.Println("Starting HTTP server on", app.Http.Addr)
		return app.Http.ListenAndServe()
	}

	return nil
}
