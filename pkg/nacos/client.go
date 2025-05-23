package nacos

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	baseUrl = "/cs/configs"
)

// Client Nacos客户端
type Client struct {
	Config *NacosConfig
}

// getAuthInfo 获取认证信息 比较operation和c.Config
func (c *Client) getAuthInfo(operation *NacosOperation) (string, string) {
	username := operation.Username
	password := operation.Password
	if username == "" {
		username = c.Config.Username
	}
	if password == "" {
		password = c.Config.Password
	}

	return username, password
}

// Get获取配置
func (c *Client) Get(operation ConfigGetOperation) (*NacosConfigDetail, error) {

	configUrl, err := getUrl(c.Config)

	if err != nil {
		return nil, err
	}

	// 构建基础请求 URL
	requestUrl := fmt.Sprintf(configUrl+"?show=all&dataId=%s&group=%s&tenant=%s", operation.DataId, operation.Group, operation.Namespace)

	// 如果配置了用户名和密码，则添加到 URL 中
	username, password := c.getAuthInfo(operation.NacosOperation)
	if username!= "" && password!= "" {
		requestUrl = fmt.Sprintf("%s&username=%s&password=%s", requestUrl, url.QueryEscape(username), url.QueryEscape(password))
	}

	resp, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	if resp.Header.Get("Content-Length") == "0" {
		return nil, errors.New("config not exists")
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response error,status code:%d\n%s", resp.StatusCode, body)
	}

	detail := NacosConfigDetail{}

	if err = json.Unmarshal(body, &detail); err != nil {
		return nil, err
	}

	return &detail, nil
}

// AllConfig 获取所有配置
func (c *Client) AllConfig(operation ConfigGetOperation) ([]NacosPageItem, error) {

	configUrl, err := getUrl(c.Config)

	if err != nil {
		return nil, err
	}

	requestUrl := fmt.Sprintf(configUrl+"?dataId=&group=%s&tenant=%s&pageNo=1&pageSize=999&search=accurate", operation.Group, operation.Namespace)

	// 如果配置了用户名和密码，则添加到 URL 中
	username, password := c.getAuthInfo(operation.NacosOperation)
	if username!= "" && password!= "" {
		requestUrl = fmt.Sprintf("%s&username=%s&password=%s", requestUrl, url.QueryEscape(username), url.QueryEscape(password))
	}

	resp, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response error,status code:%d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		body = []byte{}
	}

	result := NacosPageResult{}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result.PageItems, nil
}

// Edit 更新配置
func (c *Client) Edit(operation ConfigEditOperation) error {

	configUrl, err := getUrl(c.Config)

	if err != nil {
		return err
	}

	formValues := url.Values{
		"dataId":  []string{operation.DataId},
		"group":   []string{operation.Group},
		"content": []string{operation.Content},
		"tenant":  []string{operation.Namespace},
		"type":    []string{operation.Type},
	}

	// 如果配置了用户名和密码，则添加到表单值中
	username, password := c.getAuthInfo(operation.NacosOperation)
	if username!= "" && password!= "" {
		formValues.Add("username", username)
		formValues.Add("password", password)
	}

	resp, err := http.PostForm(configUrl, formValues)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response error,status code:%d", resp.StatusCode)
	}

	return nil
}

func getUrl(config *NacosConfig) (string, error) {
	return url.JoinPath(config.Addr, config.ApiVersion, baseUrl)
}

func NewDefaultClient() *Client {
    addr := os.Getenv("NACOS_ADDR")
    if addr == "" {
        addr = "http://127.0.0.1:8848/nacos"
    }

    apiVersion := os.Getenv("NACOS_API_VERSION")
    if apiVersion == "" {
        apiVersion = "v1"
    }

	defaultUsername := os.Getenv("NACOS_USER")
    defaultPassword := os.Getenv("NACOS_PASSWD")

    return &Client{
        Config: &NacosConfig{
            Addr:       addr,
            ApiVersion: apiVersion,
            Username:   defaultUsername,
            Password:   defaultPassword,
        },
    }
}
