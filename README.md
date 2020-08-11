# webpurify-wrapper
Wrapper for WebPurify written in golang

### Usage example

```go
package main

import (
	"github.com/ypeckstadt/webpurify-wrapper/wrapper"
	"github.com/ypeckstadt/webpurify-wrapper/wrapper/request"
	"log"
)

func main() {

	 request := request.WebPurifyRequest{
		 URL:    "http://api1-ap.webpurify.com/services/rest/",
		 APIKey: "*******",
	 }

	 replaceResponse, err := request.Replace(wrapper.ENGLISH, "fuck this", "**", true, true, true, false)
	 if err != nil {
	 	log.Print(err)
	 }

	 log.Print(replaceResponse.Text)
}
```

**Output**
```bash
2020/08/11 10:24:56 ******** this
```
