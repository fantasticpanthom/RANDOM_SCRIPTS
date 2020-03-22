import requests
import re
from bs4 import BeautifulSoup
import sys
if(len(sys.argv)!=2):
    print("usage: %s targeturl" % (sys.argv[0]))
    sys.exit(0)
urls = []
turl = sys.argv[1]    
url = requests.get(turl)
comments = re.findall('<!--(.*)-->',url.text)
if(len(comments)!=0):
    print("Comments on page: "+turl)
    for comment in comments:
        print(comment)
else:
    print("No comments found")
soup = BeautifulSoup(url.text,features="lxml")
for line in soup.find_all('a'):
    newline = line.get('href')
    try:
        if(newline[:4] == "http"):
            if turl in newline:
                urls.append(str(newline))
            elif newline[:1] == "/":
                combine = turl+newline
                urls.append(str(combine))
    except:
        pass
        print("sad reacts") 
for uurl in urls:
    print("Comments on page: "+uurl)
    url = requests.get(uurl)
    comments = re.findall('<!--(.*)-->',url.text)
    if(len(comments)!=0):
        for comment in comments:
            print(comment)                       
    else:
        print("No comments found")
