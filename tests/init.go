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

package tests

import (
	"flag"
	"github.com/primasio/wormhole/cache"
	"github.com/primasio/wormhole/config"
	"github.com/primasio/wormhole/db"
	"github.com/primasio/wormhole/models"
	"log"
	"math/rand"
	"os"
	"time"
)

func InitTestEnv(configPath string) {
	environment := flag.String("e", "test", "")

	flag.Parse()

	// Init Config
	config.Init(*environment, &configPath)

	// Init Database
	if err := db.Init(); err != nil {
		log.Println("Database:", err)
		os.Exit(1)
	}

	// Init Cache
	if err := cache.InitCache(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	models.AutoMigrateModels()

	rand.Seed(time.Now().UnixNano())
}
