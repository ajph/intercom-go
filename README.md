intercom-go
===========

Go bindings for the [Intercom API](https://api.intercom.io/docs)

##Creating an Intercom object
```go
api := intercom.NewIntercom("appId", "apiKey")
```

##Creating/Updating a User
```go
package main

import (
	"intercom"
	"log"
) //import

func main() {
	var (
		err  error
		user *intercom.User_t = intercom.NewUser("POST")
	) //var

	api := intercom.NewIntercom("app_id", "api_key")

	user.Email = "foo@example.com"
	user.Name = "Roky Erickson"
	user.CustomAttributes["type"] = "user"

	if err = api.PostUser(user); err != nil {
		log.Fatalln(err)
	}//if
} //main
```

##Creating/Updating a Company
```go
package main

import (
	"intercom"
	"log"
) //import

func main() {
	var (
		err error
		com *intercom.Company_t = intercom.NewCompany("POST")
	) //var

	api := intercom.NewIntercom("app_id", "api_key")

	com.CompanyId = "1612"
	com.Name = "Enron"
	com.Plan = "Premium"
	com.CustomAttributes["Downloads remaining"] = 12

	if err = api.PostCompany(com); err != nil {
		log.Fatalln(err)
	}//if
} //main
```

##Submitting an event
```go
package main

import (
	"intercom"
	"log"
	"time"
) //import

func main() {
	var (
		err error
		e   intercom.Event_t
	) //var

	api := intercom.NewIntercom("app_id", "api_key")

	e.EventName = "event-name-here"
	e.CreatedAt = time.Now().Unix()
	e.UserId = "8"
	e.Metadata = map[string]interface{} {
		"Metadata": "Goes here",
	} //map

	if err = api.PostEvent(e); err != nil {
		log.Fatalln(err)
	} //if
} //main
```

##Tagging users
```go
package main

import (
	"intercom"
	"log"
) //import

func main() {
	var (
		err error
		tag intercom.Tag_t
	) //var

	api := intercom.NewIntercom("app_id", "api_key")

	tag.Name = "new-user"
	tag.Users = append(tag.Users,
		intercom.TagUser_t{
			Email: "rerickson@gmail.com",
			Untag: false,
		})

	if err = api.PostTag(tag); err != nil {
		log.Fatalln(err)
	} //if
} //main
```

#License
The MIT License (MIT)
Copyright (c) 2014 CB Insights
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
