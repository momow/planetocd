package server

// DefaultPort ...
const DefaultPort = 4242

// Domain ...
const Domain = "planetocd.org"

// DonateURL ...
const DonateURL = "https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=23LG7JTZSCA54"

// Email ...
const Email = "aurelien.ocd@gmail.com"

// FacebookURL ...
const FacebookURL = "https://www.facebook.com/PlanetOCDfr"

// GithubURL ...
const GithubURL = "https://github.com/momow/planetocd"

// GithubContributeURL ...
const GithubContributeURL = GithubURL + "/blob/master/README.md#contribute"

// SiteName ...
const SiteName = "Planet OCD"

// TwitterHandle ...
const TwitterHandle = "PlanetOCD_fr"

// TwitterURL ...
const TwitterURL = "https://twitter.com/" + TwitterHandle

// Constants ...
var Constants = map[string]interface{}{
	"DefaultPort":         DefaultPort,
	"Domain":              Domain,
	"DonateURL":           DonateURL,
	"Email":               Email,
	"FacebookURL":         FacebookURL,
	"GithubURL":           GithubURL,
	"GithubContributeURL": GithubContributeURL,
	"SiteName":            SiteName,
	"TwitterHandle":       TwitterHandle,
	"TwitterURL":          TwitterURL,
}
