package utils

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/svcodestore/sv-resource-gin/global"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type CrudRequestData struct {
	Creates []map[string]interface{} `json:"A"`
	Updates []map[string]interface{} `json:"U"`
	Deletes []string                 `json:"D"`
}

func Post(host string, postVaules url.Values) (body []byte, err error) {
	// uri := url.URL{
	// 	Host: host,
	// }
	postString := postVaules.Encode()

	req, err := http.NewRequest("POST", host, strings.NewReader(postString))

	if err != nil {
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	return
}

func ParseBatchCrudFormData(values url.Values) (adds []map[string]interface{}, updates []map[string]interface{}, deletes []string) {
	for k, val := range values {
		data := SplitByPairSymbol(k, "[", "]")
		if strings.HasPrefix(k, "D") {
			deletes = append(deletes, val...)
		} else {
			i, e := strconv.Atoi(data[0])
			if e != nil {
				continue
			}
			v := ""
			if len(val) > 0 {
				v = val[0]
			}
			if strings.HasPrefix(k, "A") {
				if len(adds) == 0 {
					adds = append(adds, map[string]interface{}{})
				}
				if len(adds) > i {
					if v != "" {
						adds[i][data[1]] = v
					}
				} else {
					// problem
					adds = append(adds, map[string]interface{}{})
					if v != "" {
						adds[i][data[1]] = v
					}
				}
			}
			if strings.HasPrefix(k, "U") {
				if len(updates) == 0 {
					updates = append(updates, map[string]interface{}{})
				}
				if len(updates) > i {
					updates[i][data[2]] = v
					updates[i]["id"] = data[1]
				} else {
					// problem
					updates = append(updates, map[string]interface{}{})
					updates[i][data[2]] = v
					updates[i]["id"] = data[1]
				}
			}
		}
	}
	return
}

func ExecFormCrudBatch(values url.Values, addCb func(b []byte) (err error), updateCb func(b []byte) (err error), deleteCb func(ids []string) (err error)) (err error) {
	adds, updates, deletes := ParseBatchCrudFormData(values)
	if adds == nil && updates == nil && deletes == nil {
		return
	}

	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Error; err != nil {
		return err
	}

	addLen := len(adds)
	if addLen > 0 {
		for i := 0; i < addLen; i++ {
			b, e := json.Marshal(adds[i])
			if e != nil {
				err = e
				tx.Rollback()
				return
			}
			e = addCb(b)
			if e != nil {
				err = e
				tx.Rollback()
				return
			}
		}
	}

	updateLen := len(updates)
	if updateLen > 0 {
		for i := 0; i < updateLen; i++ {
			b, e := json.Marshal(updates[i])
			if e != nil {
				err = e
				tx.Rollback()
				return
			}
			e = updateCb(b)
			if e != nil {
				err = e
				tx.Rollback()
				return
			}
		}
	}

	deleteLen := len(deletes)
	if deleteLen > 0 {
		e := deleteCb(deletes)
		if e != nil {
			err = e
			tx.Rollback()
			return
		}
	}

	err = tx.Commit().Error

	return
}

func ExecJsonCrudBatch(data *CrudRequestData, addCb func(b []byte) (err error), updateCb func(b []byte) (err error), deleteCb func(ids []string) (err error)) (err error) {
	adds := data.Creates
	updates := data.Updates
	deletes := data.Deletes

	if adds == nil && updates == nil && deletes == nil {
		return
	}

	tx := global.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err = tx.Error; err != nil {
		return err
	}

	if adds != nil {
		addLen := len(adds)
		if addLen > 0 {
			for i := 0; i < addLen; i++ {
				b, e := json.Marshal(adds[i])
				if e != nil {
					err = e
					tx.Rollback()
					return
				}
				e = addCb(b)
				if e != nil {
					err = e
					tx.Rollback()
					return
				}
			}
		}
	}

	if updates != nil {
		updateLen := len(updates)
		if updateLen > 0 {
			for i := 0; i < updateLen; i++ {
				b, e := json.Marshal(updates[i])
				if e != nil {
					err = e
					return
				}
				e = updateCb(b)
				if e != nil {
					err = e
					tx.Rollback()
					return
				}
			}
		}
	}

	if deletes != nil {
		deleteLen := len(deletes)
		if deleteLen > 0 {
			e := deleteCb(deletes)
			if e != nil {
				err = e
				tx.Rollback()
				return
			}
		}
	}

	err = tx.Commit().Error

	return
}
