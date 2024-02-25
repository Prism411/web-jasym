import requests

def buscar_google(query):
    api_key = "bruh"
    cx = "bruh"
    url = f"https://www.googleapis.com/customsearch/v1?key={api_key}&cx={cx}&q={query}"

    response = requests.get(url)
    result = response.json()

    for item in result.get("items", []):
        title = item.get("title")
        snippet = item.get("snippet")
        link = item.get("link")
        print(f"Título: {title}\nDescrição: {snippet}\nLink: {link}\n")


