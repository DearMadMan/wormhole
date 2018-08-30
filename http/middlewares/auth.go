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

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/primasio/wormhole/cache"
	"log"
	"strconv"
)

const AuthorizedUserId = "UserId"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		reqToken := c.Request.Header.Get("Authorization")

		if reqToken == "" {
			c.AbortWithStatus(401)
			return
		}

		// Check token validity

		if err, userId := cache.SessionGet(reqToken); err != nil {

			log.Println("token not exist", err)
			c.AbortWithStatus(500)

		} else {

			if userId == "" {
				c.AbortWithStatus(401)
			} else {

				userIdNum, err := strconv.Atoi(userId)

				if err != nil {
					log.Println(err)
					c.AbortWithStatus(500)
				}

				c.Set(AuthorizedUserId, uint(userIdNum))
				c.Next()
			}
		}
	}
}
