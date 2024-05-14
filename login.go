/* Copyright 2024 Digilol OÜ
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client3xui

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type loginResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

func (c *Client) login(ctx context.Context) error {
	loginReq := url.Values{"username": {c.username}, "password": {c.password}}
	b := strings.NewReader(loginReq.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url+"/login", b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var loginResp loginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return err
	}
	if !loginResp.Success {
		return errors.New(loginResp.Msg)
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "session" {
			c.sessionCookie = cookie
			return nil
		}
	}
	return errors.New("session cookie not found")
}

func (c *Client) loginIfNoCookie(ctx context.Context) error {
	if c.sessionCookie != nil {
		return nil
	}
	return c.login(ctx)
}
