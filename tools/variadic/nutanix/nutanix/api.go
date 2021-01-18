package nutanix

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// API network info
type API struct {
	Username string
	Password string
	URL      string
	Host     string
	APIv3    string
	Port     string
	Path     string
	Payload  *strings.Reader
}

// NewRequest creates a new api request
func NewRequest(u, p string) API {
	ret := API{
		Username: u,
		Password: p,
		Host:     "192.168.216.193",
		Port:     "9440",
		APIv3:    "api/nutanix/v3",
		Payload:  strings.NewReader(`{}`),
	}
	return ret
}

// Some basic setup
func init() {
	fmt.Println
	customTransport := http.DefaultTransport.(*http.Transport)
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

// Request create a new request
func (a *API) Request(query string) string {

	a.Path = fmt.Sprintf("https://%s:%s/%s/%s", a.Host, a.Port, a.APIv3, query)

	req, _ := http.NewRequest("POST", a.Path, a.Payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth(a.Username, a.Password))
	// req.BasicAuth()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error on post", err)
	}

	defer func() {
		// fmt.Println("We are in defer")
		res.Body.Close()
	}()

	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	return string(body)

}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
