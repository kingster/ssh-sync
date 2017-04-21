package main

import (
	"os/exec"
	"strings"
	"regexp"
	"io/ioutil"
	"os/user"
	"os"
	"time"
	"log"
)

func main() {

	excludedFiles := []string{"authorized_keys", "known_hosts" }

	_, lookErr := exec.LookPath("lpass")
	if lookErr != nil {
		panic(lookErr)
	}

	log.Println("[INFO] LastPass SSH Sync...")
	sync()

	lsCmd := exec.Command("lpass", "ls", "--color=never", "--long", "Secure Notes\\SSH")
	currentKeys, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	entries := strings.Split(string(currentKeys), "\n")

	loc, _ := time.LoadLocation("Australia/Darwin") //lpass api returns in this timezone.
	timeFormat := "2006-01-02 15:04"

	upStreamMod := make(map[string]int64)
	upStreamId := make(map[string]string)

	re := regexp.MustCompile(`([\d- :]+)Secure Notes\\SSH\/([^ ]+) \[id: (\d+)\]`)
	for _, element := range entries {
		match := re.FindStringSubmatch(element)
		if (len(match) == 4) {
			upStreamId[match[2]] = match[3]

			t, err := time.ParseInLocation(timeFormat, strings.Trim(match[1], " "), loc)
			if err != nil {
				panic(err)
			}
			upStreamMod[match[2]] = unixMilli(t)
		}
	}

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	files, _ := ioutil.ReadDir(usr.HomeDir + "/.ssh/")
	for _, f := range files {
		name := f.Name()
		if !contains(excludedFiles, name) {
			log.Println("[DEBUG] Processing Folder File", name)
			fName := usr.HomeDir + "/.ssh/" + f.Name()
			modTime, exists := upStreamMod[name]
			if (exists) {
				info, err := os.Stat(fName)
				if err != nil {
					panic(err)
				}
				localModTime := unixMilli(info.ModTime())
				if (modTime > localModTime) {
					log.Println("[INFO] Downloading Updated File :", name, "was modified at", modTime)
					download(upStreamId[name], fName, upStreamMod[name])
				} else if (localModTime > modTime ) {
					log.Println("[INFO] Upload Updated File :", name, "was modified at", unixMilli(info.ModTime()))
					upload(name, fName)
				} else {
					log.Println("[DEBUG] No Change File :", name)
				}

			} else {
				log.Println("[INFO] Uploading New File :", name)
				upload(name, fName)
			}
			delete(upStreamId, name)
		}

	}

	for name, id := range upStreamId {
		log.Println("[INFO]  Downloading New File :", name)
		fName := usr.HomeDir + "/.ssh/" + name
		download(id, fName,upStreamMod[name])
	}

	sync()
}

func download(id string, path string, updatedAt int64) {

	uploadCmd := exec.Command("lpass", "show", id, "--notes")
	uploadOut, _ := uploadCmd.StdoutPipe()
	uploadCmd.Start()

	output, _ := ioutil.ReadAll(uploadOut)
	err := uploadCmd.Wait()
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, output, 0600)
	if err != nil {
		panic(err)
	}

	modTime := time.Unix(0, updatedAt * int64(time.Millisecond))
	touchCmd := exec.Command("touch", "-t", modTime.Format("200601021504"), path)
	_, err = touchCmd.Output()
	if err != nil {
		panic(err)
	}
	

}

func upload(name string, path string) {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	dest := "Secure Notes\\SSH/" + name
	uploadCmd := exec.Command("lpass", "edit", "--non-interactive", "--notes", dest)
	uploadIn, _ := uploadCmd.StdinPipe()
	uploadCmd.Start()
	uploadIn.Write(data)
	uploadIn.Close()
	err = uploadCmd.Wait()
	if err != nil {
		panic(err)
	}

}

func sync()  {
	log.Println("[INFO] Running Sync...")
	lsCmd := exec.Command("lpass", "sync")
	_, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func unixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}