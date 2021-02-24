package comic

import (
	"encoding/json"
	"fmt"
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

type readAllExhaustedError struct{}

func (e *readAllExhaustedError) Error() string {
	return fmt.Sprintf("Extract finished")
}

func init() {
	path := fullDatadir()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Data dir '%s' doesn't exist yet. Creating it.", path)
		os.Mkdir(path, 0700)
	}
}

// FetchAll extracts all comic descriptions from xkcd.com and writes each one to disk.
// Will not run if data is already cached unless a force flag is passed.
func FetchAll(force bool) {
	// create data directory.
	// cache each comic description in a single json file by it's ID.
	// naively fetch starting from 1 until we get 3 errors (because num 404 returns a 404 😂 but isn't the last one.)
	fails := 0
	for i := 1; ; i++ {
		if cached(i) && !force {
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
		comic.write()
		time.Sleep(rateLimit)
	}
}

func read(path string) (Comic, error) {
	bytes, err := ioutil.ReadFile(path)
	var comic Comic
	if err != nil {
		return comic, err
	}
	err = json.Unmarshal(bytes, &comic)
	if err != nil {
		return comic, err
	}
	return comic, err
}

func readAll() func() (Comic, error) {
	d, err := ioutil.ReadDir(fullDatadir())
	if err != nil {
		log.Fatalf("Failed to read raw data dir. %s", err)
	}
	i := 0
	return func() (Comic, error) {
		var comic Comic
		if i >= len(d) {
			return comic, new(readAllExhaustedError)
		}
		path := filepath.Join(fullDatadir(), d[i].Name())
		i++
		return read(path)
	}
}

// fetch returns the json response of a comic or error
func fetch(id int) (Comic, error) {
	var comic Comic
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
	p, err := filepath.Abs(filepath.Join(".", datadir))
	if err != nil {
		log.Fatalf("Unable to produce data directory filepath. %s", err)
	}
	return p
}

func fname(id int) string {
	return fmt.Sprintf("%d.json", id)
}

func (comic *Comic) write() {
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

func cached(id int) bool {
	path := filepath.Join(fullDatadir(), fname(id))
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
