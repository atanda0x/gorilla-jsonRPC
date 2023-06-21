package main

import (
	jsonparse "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// Args holds arguments passed to JSON RPC service
type Args struct {
	Name string
}

// TwitterProfie struct holds TwitterProfile JSON structure
type TwitterProfile struct {
	Name      string `json:"name,omitempty"`
	Username  string `json:"username,omitempty"`
	Followers string `json:"followers,omitempty"`
	Following string `json:"following,omitempty"`
}

type JSONServer struct{}

// TwitterProfileDetail
func (t *JSONServer) TwitterProfileDetail(r *http.Request, args *Args, reply *TwitterProfile) error {
	var twitterprofiles []TwitterProfile
	// Read JSON file and load data
	raw, readerr := ioutil.ReadFile("./twitterprofile.json")
	if readerr != nil {
		log.Println("error:", readerr)
		os.Exit(1)
	}
	// Unmarshal JSON raw data into twitterprofiles array
	marshalerr := jsonparse.Unmarshal(raw, &twitterprofiles)
	if marshalerr != nil {
		log.Println("error:", marshalerr)
		os.Exit(1)
	}
	// Iterate over each twitterprofile to find the given twitterprofile
	for _, twitterprofile := range twitterprofiles {
		if twitterprofile.Name == args.Name {
			// If twitterprofile found, fill reply with it
			fmt.Println(twitterprofile)
			*reply = twitterprofile
			break
		}
	}
	return nil
}
func main() {
	// Create a new RPC server
	s := rpc.NewServer() // Register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")
	// Register the service by creating a new JSON server
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":9000", r)
}
