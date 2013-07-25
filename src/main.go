package main

import (
	"container/list"
	"encoding/json"
	"entity"
	"fmt"
	"io/ioutil"
	"os/exec"
	"runtime"
	"time"
	"util"
)

const (
	BROWSER_PATH = "/usr/local/firefox/firefox"
)

var (
	LAST_MODIFIED time.Time
	IS_WORKER     bool
	RANGE_TIME    int64 = 600
	QUEUE               = list.New()
)

func Load(url string) *entity.Task {
	response, err_con := util.GetUrlInUserAgent(url)
	task := &entity.Task{}
	if err_con != nil {
		util.ERROR("connect ERROR, %s", err_con)
		util.Connect()
	} else {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(body, &task)
		task.Size = len(task.Users)
		las_modify, err_parse := time.Parse(time.RFC1123, response.Header.Get("Last-Modified"))
		if err_parse != nil {
			util.ERROR("Parse time is ERROR: %s", err_parse)
		} else {
			if las_modify.After(LAST_MODIFIED) {
				if LAST_MODIFIED.IsZero() {
					util.INFO("Last-Modified is NULL, program is first run, Last-Modifyed: %s", las_modify)
				} else {
					util.INFO("file is change, Last-Modifyed: %s", las_modify)
				}
				LAST_MODIFIED = las_modify
				if task.Start {
					IS_WORKER = true
					util.INFO("start worker!")
				} else {
					IS_WORKER = false
					util.INFO("worker is not start!")
				}
			}
		}
	}
	return task
}

func Jobs(task *entity.Task) {
	for _, user := range task.Users {
		user.Date = time.Unix(user.Trigger, user.Trigger)
		if user.Start && time.Now().Unix()-user.Trigger < RANGE_TIME {
			go Task(user)
		}
	}
	util.INFO("shutdown worker!")
	IS_WORKER = false
}

func Task(user *entity.User) {
	runtime.Gosched()
	util.DEBUG("add job username: %s", user.UserName)
	for {
		util.DEBUG("loop task username: %s, trigger: %d, current: %d", user.UserName, user.Trigger, time.Now().Unix())
		if time.Now().After(user.Date) && time.Now().Unix()-user.Trigger < RANGE_TIME {
			util.DEBUG("jobs username: %s, password: %s, start: %t, trigger: %d, date: %s",
				user.UserName, user.PassWord, user.Start, user.Trigger, user.Date)
			QUEUE.PushBack(util.HtmlFile(user))
			break
		} else {
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}

func OpenBrowser(filename string) {
	runtime.Gosched()
	cmd := exec.Command(BROWSER_PATH, filename)
	err_run := cmd.Run()
	if err_run != nil {
		util.ERROR("start browser file [%s] ERROR: %s", filename, err_run)
	}
}

func start() {
	runtime.Gosched()
	util.INFO("start ....")
	for {
		if QUEUE.Len() > 0 {
			task := QUEUE.Back()
			filename := fmt.Sprintf("%s", task.Value)
			util.INFO("open browser file: %s", filename)
			go OpenBrowser(filename)
			QUEUE.Remove(task)
		} else {
			time.Sleep(time.Duration(10) * time.Second)
		}
		time.Sleep(time.Duration(5) * time.Second)
	}
}

func main() {
	runtime.GOMAXPROCS(8)
	go start()
	for {
		task := Load("http://task.open-ns.org/task.json")
		if IS_WORKER {
			util.DEBUG("load user [%d] size", task.Size)
			util.INFO("worker is true, go jobs")
			Jobs(task)
		}
		time.Sleep(time.Duration(3) * time.Second)
	}
}
