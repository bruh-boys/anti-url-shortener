// a little test with deno
while (true) {
  const rawResponse = await fetch("http://localhost:8000/no-bitly", {
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
}
