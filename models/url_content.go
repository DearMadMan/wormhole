/*
 * Copyright 2018 Primas Lab Foundation
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

package models

type URLContent struct {
	BaseModel
	UserId   uint   `json:"-"`
	URL      string `gorm:"type:varchar(2048);unique_index" json:"url"`
	Title    string `gorm:"type:text" json:"title"`
	Abstract string `gorm:"type:text" json:"abstract"`
	Content  string `gorm:"type:longtext" json:"content"`

	IsActive     bool `gorm:"default:false" json:"is_active"`
	Votes        uint `gorm:"default:1" json:"votes"`
	TotalComment uint `gorm:"default:0" json:"total_comment"`
}
