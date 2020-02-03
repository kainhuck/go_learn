import requests

url = "http://movie.douban.com/top250"

headers = {
   "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36"
}
r = requests.get(url, headers=headers)

print(r.text)
