package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"
        "html/template"
        "golang.org/x/net/context"
        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/sheets/v4"
)


var tpl *template.Template

type pageData struct {
	Title     string
	FirstName string
}
func getClient(config *oauth2.Config) *http.Client {
        // The file token.json stores the user's access and refresh tokens, and is
        // created automatically when the authorization flow completes for the first
        // time.
        tokFile := "token.json"
        tok, err := tokenFromFile(tokFile)
        if err != nil {
                tok = getTokenFromWeb(config)
                saveToken(tokFile, tok)
        }
        return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
        authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
        fmt.Printf("Go to the following link in your browser then type the "+
                "authorization code: \n%v\n", authURL)

        var authCode string
        if _, err := fmt.Scan(&authCode); err != nil {
                log.Fatalf("Unable to read authorization code: %v", err)
        }

        tok, err := config.Exchange(context.TODO(), authCode)
        if err != nil {
                log.Fatalf("Unable to retrieve token from web: %v", err)
        }
        return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
        f, err := os.Open(file)
        if err != nil {
                return nil, err
        }
        defer f.Close()
        tok := &oauth2.Token{}
        err = json.NewDecoder(f).Decode(tok)
        return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
        fmt.Printf("Saving credential file to: %s\n", path)
        f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
        if err != nil {
                log.Fatalf("Unable to cache oauth token: %v", err)
        }
        defer f.Close()
        json.NewEncoder(f).Encode(token)
}


func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func getport() string{
    p := os.Getenv("PORT")

   if p!="" {
      return ":" + p
   }
   return ":8080"
}

func main() {
    port := getport()

	http.HandleFunc("/index", idx)
	http.HandleFunc("/link1", lk1)
	http.HandleFunc("/link2", lk2)
	http.HandleFunc("/link3", lk3)
	http.HandleFunc("/link4", lk4)
	http.HandleFunc("/link5", lk5)
	http.HandleFunc("/pageviews", pg)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(port, nil)

	if err != nil {
	   panic(err)
	}
}



func idx(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Index Page",
	}
	   

	err := tpl.ExecuteTemplate(w, "index.gohtml",pd)

	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal serverrrrrr error", http.StatusInternalServerError)
		return
	}
}

func pg(w http.ResponseWriter, req *http.Request) {
     var link[5]string
	 b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        client := getClient(config)

        srv, err := sheets.New(client)
        if err != nil {
                log.Fatalf("Unable to retrieve Sheets client: %v", err)
        }

        // Prints the names and majors of students in a sample spreadsheet:
        // https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
        spreadsheetId := "1WWqMBsB-iy6ZwD9V5efgVaHWQeNzI-8ideKCqhNP4VU"
        readRange := "A2:C"
        resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
        if err != nil {
                log.Fatalf("Unable to retrieve data from sheet: %v", err)
        }
        
        var count[5]string
        if len(resp.Values) == 0 {
                fmt.Println("No data found.")
        } else {
                fmt.Println("Link, Count:")
                i:=0
                for _, row := range resp.Values {
                        // Print columns A and E, which correspond to indices 0 and 4.
                        if (i>14 && i<20) {
                    
                        link[i-15]=row[0].(string)
                        count[i-15]=row[1].(string)
                        }
                     i+=1
                }
        }

         n := 5
            // set swapped to true
            swapped := true
            // loop
            for swapped {
                // set swapped to false
                swapped = false
                // iterate through all of the elements in our list
                for i := 1; i < n; i++ {
                    // if the current element is greater than the next
                    // element, swap them
                    if count[i-1] < count[i] {
                        // log that we are swapping values for posterity
                       
                        // swap values using Go's tuple assignment
                        link[i], link[i-1] = link[i-1], link[i]
                        count[i], count[i-1] = count[i-1], count[i]

                        // set swapped to true - this is important
                        // if the loop ends and swapped is still equal
                        // to false, our algorithm will assume the list is
                        // fully sorted.
                        swapped = true
                    }
                }
               
}
	
	countryCapitalMap := make(map[string]string)
   
   /* insert key-value pairs in the map*/
   countryCapitalMap["one"] = link[0]
   countryCapitalMap["two"] = link[1]
   countryCapitalMap["three"] = link[2]
   countryCapitalMap["four"] = link[3]
   countryCapitalMap["five"] = link[4]
  
	err1 := tpl.ExecuteTemplate(w, "pageviews.gohtml", countryCapitalMap)
	if err1 != nil {

		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
func lk1(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "link1 Page",
	}

	err := tpl.ExecuteTemplate(w, "link1.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
func lk2(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "link2 Page",
	}

	err := tpl.ExecuteTemplate(w, "link2.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func lk3(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "link3 Page",
	}

	err := tpl.ExecuteTemplate(w, "link3.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func lk4(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "link4 Page",
	}

	err := tpl.ExecuteTemplate(w, "link4.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func lk5(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "link5 Page",
	}

	err := tpl.ExecuteTemplate(w, "link5.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
