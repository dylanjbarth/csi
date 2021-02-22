package extract

import (
	"encoding/json"
	"fmt"
	"gopl-exercises/ex-4-12/xkcd/types"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const urlTemp = "https://xkcd.com/%d/info.0.json"
const rateLimit = 300 * time.Millisecond
const datadir = "raw_data"

// ReadAllExhaustedError returned by ReadAll generator when completed.
type ReadAllExhaustedError struct{}

func (e *ReadAllExhaustedError) Error() string {
	return fmt.Sprintf("Extract finished")
}

func init() {
	path := fullDatadir()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Data dir '%s' doesn't exist yet. Creating it.", path)
		os.Mkdir(path, 0700)
	}
}

// Extracts comic descriptions from xkcd.com to build an offline index.
// Will not run if data is already cached unless a force flag is passed.
func main() {
	// create data directory.
	// cache each comic description in a single json file by it's ID.
	// naively fetch starting from 1 until we get 3 errors (because num 404 returns a 404 😂 but isn't the last one.)
	fails := 0
	for i := 1; ; i++ {
		if cached(i) {
			log.Printf("Skipping %d because it's already cached locally.\n", i)
			continue
		}
		log.Printf("Fetching comic num %d\n", i)
		comic, err := fetch(i)
		if err != nil {
			fails++
			if fails > 2 {
				log.Fatalf("Failed more than 3 times. Assuming we have reached the end. %s", err)
			} else {
				log.Printf("Failed to fetch %d. %s. Total fail count is %d. Continuing to try.", i, err, fails)
			}
			continue
		}
		write(&comic)
		time.Sleep(rateLimit)
	}
}

// fetch returns the json response of a comic or error
func fetch(id int) (types.Comic, error) {
	var comic types.Comic
	resp, err := http.Get(fmt.Sprintf(urlTemp, id))
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("HTTP fetch failed. Code: %s", resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, fmt.Errorf("JSON decoding failed: %s", err)
	}
	return comic, nil
}

func fullDatadir() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to fetch pwd %s", err)
	}
	return filepath.Join(pwd, "..", datadir)
}

func fname(id int) string {
	return fmt.Sprintf("%d.json", id)
}

func write(comic *types.Comic) {
	b, err := json.Marshal(comic)
	if err != nil {
		log.Fatal(err)
	}
	fullpath := filepath.Join(fullDatadir(), fname(comic.Num))
	log.Printf("Writing comic metadata to disk => %s\n", fullpath)
	if err = ioutil.WriteFile(fullpath, b, 0600); err != nil {
		log.Fatalf("Failed to write file: %s", err)
	}
}

// Read returns the comic with the num specified.
func Read(path string) (types.Comic, error) {
	bytes, err := ioutil.ReadFile(path)
	var comic types.Comic
	if err != nil {
		return comic, err
	}
	err = json.Unmarshal(bytes, &comic)
	if err != nil {
		return comic, err
	}
	return comic, err
}

// ReadAll returns an generator that will return all the comics in the raw data cache one at a time
func ReadAll() func() (types.Comic, error) {
	d, err := ioutil.ReadDir(fullDatadir())
	if err != nil {
		log.Fatalf("Failed to read raw data dir. %s", err)
	}
	i := 0
	return func() (types.Comic, error) {
		var comic types.Comic
		if i >= len(d) {
			return comic, new(ReadAllExhaustedError)
		}
		path := filepath.Join(fullDatadir(), d[i].Name())
		i++
		return Read(path)
	}
}

func cached(id int) bool {
	path := filepath.Join(fullDatadir(), fname(id))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
