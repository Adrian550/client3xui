/* Copyright 2025 İrem Kuyucu <irem@digilol.net>
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
	"fmt"
	"net/http"
)

type UpdateClientRequest struct {
	// Inbound ID
	ID uint `json:"id"`

	// See ClientSettings. This is it marshaled into an escaped JSON string.
	Settings string `json:"settings"`
}

// 				In add_client.go	 			//

// type ClientSettings struct {
// 	Clients []XrayClient `json:"clients"`
// }

// type XrayClient struct {
// 	ID         string `json:"id"`
// 	Flow       string `json:"flow"`
// 	AlterID    uint   `json:"alter_id,omitempty"`
// 	Email      string `json:"email"`
// 	LimitIP    uint   `json:"limitIp"`
// 	TotalGB    uint64 `json:"totalGB"`
// 	ExpiryTime uint64 `json:"expiryTime"`
// 	Enable     bool   `json:"enable"`
// 	TgID       uint   `json:"tgId"`
// 	SubID      string `json:"subId"`
// }

// Update client in inbound.
func (c *Client) UpdateClient(ctx context.Context, inboundId uint, clientId string, clients []XrayClient) (*ApiResponse, error) {
	settings := &ClientSettings{Clients: clients}
	settingsBytes, err := json.Marshal(settings)
	if err != nil {
		return nil, err
	}

	req := &AddClientRequest{ID: inboundId, Settings: string(settingsBytes)}
	resp := &ApiResponse{}
	err = c.Do(ctx, http.MethodPost, "/panel/api/inbounds/updateClient/"+clientId, req, resp)
	if err != nil {
		return nil, err
	}
	if !resp.Success {
		return resp, fmt.Errorf(resp.Msg)
	}
	return resp, err
}
