package webserver

type WebServerStarter struct {
	WebServer WebServer
}

func NewWebServerStarter(webServer WebServer) *WebServerStarter {
	return &WebServerStarter{
		WebServer: webServer,
	}
}
