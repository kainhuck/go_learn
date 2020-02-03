import requests
from lxml import etree


def poemSpider():
    url = "https://so.gushiwen.org/authors/"

    # 1. requests
    r = requests.get(url)
    pageHTML = r.text
    
    # 2. parse
    doc = etree.HTML(pageHTML)
    link = doc.xpath("//div[@class='cont']/a/@href")

    # 3. print
    for each in link:
        print(each)

    

if __name__ == "__main__":
    poemSpider()