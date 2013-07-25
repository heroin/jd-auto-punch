package util

import (
	"entity"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

const (
	HTML_TEMPLATE = `<!DOCTYPE html><html xmlns="http://www.w3.org/1999/xhtml"><head><title></title></head><body onload="Button1.click()"><form name="form1" method="post" action="******" id="form1">%s<input name="txt_UserName" type="text" value="%s" id="txt_UserName"/><input name="txt_Password" type="password" value="%s" id="txt_Password"/><input type="submit" name="Button1" value="" id="Button1"/></form></body></html>`
)

func getHidden() string {
	response, err_con := GetUrlInUserAgent("******")
	hidden := ""
	if err_con != nil {
		ERROR("connect ERROR, %s", err_con)
		Connect()
	} else {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		re_hidden, _ := regexp.Compile(`<input type="hidden" (.*?) />`)
		result := re_hidden.FindAll(body, -1)
		hidden = fmt.Sprintf("%s", result)
	}
	return hidden
}

func HtmlFile(user *entity.User) string {
	hidden := getHidden()
	filename := fmt.Sprintf("%s-%d.html", user.UserName, user.Trigger)
	html := fmt.Sprintf(HTML_TEMPLATE, hidden, user.UserName, user.PassWord)
	out, err := os.Create(filename)
	if err != nil {
		ERROR("create file ERROR: %s", err)
	} else {
		defer out.Close()
		out.WriteString(html)
	}
	return filename
}
