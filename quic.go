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

// QUIC was deprecated in favor of XHTTP in Xray-core v24.9.30
type QuicSetting struct {
	Security string        `json:"security"`
	Key      string        `json:"key"`
	Header   HeaderSetting `json:"header"`
}

type QuicStreamSetting struct {
	Network       string      `json:"network"`
	Security      string      `json:"security"`
	ExternalProxy []string    `json:"externalProxy"`
	QuicSettings  QuicSetting `json:"quicSettings"`
}
