<h1>check the real url from a url shortener (bit.ly)</h1>

<div align="center">
<img src="https://media.discordapp.net/attachments/887003170377719840/887174274538954823/unknown.png?width=537&height=568"/>
<img src="https://media.discordapp.net/attachments/887003170377719840/887181476980985886/unknown.png"/>

**Also you can use it as an API**

</div>

example with deno

```ts

  const rawResponse = await fetch("https://anti-url-shortener.herokuapp.com/no-bitly", {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      url: `https://bit.ly/3nv15Ci`,
    }),
  });
  const content = await rawResponse.json();
  console.log(content);


```

example with curl

```sh
curl --location --request POST 'https://anti-url-shortener.herokuapp.com/no-bitly' \
--header 'Content-Type: application/json' \
--data-raw '{"url":"https://bit.ly/3nv15Ci"}'
```

example with python

```py
import requests
import json

url = "https://anti-url-shortener.herokuapp.com/no-bitly"

payload = json.dumps({
  "url": "https://bit.ly/3nv15Ci"
})
headers = {
  'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)

```

example with nodejs (axios)

```js
var axios = require('axios');
var data = JSON.stringify({
  "url": "https://bit.ly/3nv15Ci"
});

var config = {
  method: 'post',
  url: 'https://anti-url-shortener.herokuapp.com/no-bitly',
  headers: { 
    'Content-Type': 'application/json'
  },
  data : data
};

axios(config)
.then(function (response) {
  console.log(JSON.stringify(response.data));
})
.catch(function (error) {
  console.log(error);
});

```




