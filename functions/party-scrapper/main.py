import json

import requests
from bs4 import BeautifulSoup
import firebase_admin
import network
import utils
from utils import convert_to_epoch

def update_parties(request):
    total_page = 10
    events = []
    for page in range(1, total_page):
        url = f"https://shotgun.live/paris?page={page}"
        print(url)
        events.extend(get_all_events(url))
    print(events)

    token = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjYzODBlZjEyZjk1ZjkxNmNhZDdhNGNlMzg4ZDJjMmMzYzIzMDJmZGUiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiSHVnbyBTYWxhemFyIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FBY0hUdGVwQmJRMjA0SGtTM0VmemNLUFdjVk41SVZfUlkzS19HY2RaTVpPeGR6Ui1IVT1zOTYtYyIsImFkbWluIjp0cnVlLCJtb2RlcmF0b3IiOnRydWUsInZlcmlmaWVkIjp0cnVlLCJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vY2hlZXJzLWEyNzVlIiwiYXVkIjoiY2hlZXJzLWEyNzVlIiwiYXV0aF90aW1lIjoxNjkyODkyMjc1LCJ1c2VyX2lkIjoibUlGclZwQ0NKcGJjU3dNTFo0OVJCWDEyRThtMSIsInN1YiI6Im1JRnJWcENDSnBiY1N3TUxaNDlSQlgxMkU4bTEiLCJpYXQiOjE2OTI4OTIyNzUsImV4cCI6MTY5Mjg5NTg3NSwiZW1haWwiOiJodWdvYnJvY2s3NEBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJnb29nbGUuY29tIjpbIjEwNjgxNTcyOTkwNzU5NzQxNDk0MiJdLCJlbWFpbCI6WyJodWdvYnJvY2s3NEBnbWFpbC5jb20iXX0sInNpZ25faW5fcHJvdmlkZXIiOiJnb29nbGUuY29tIn19.i3brKKb5OrR5kZo0ltATG06p1D6a5II9LWXrL2JI59pJ93bNzPdIxSM8K5y5978KlVQECZQ93MTZHnmtPwPoQutWiqVrHoEtYG9hAnJoTmYPiXZNzelQil-a1pzTCDKOgT_X8qK23FcW2Is56AjA-B5FPHkiDdq6O3dc_wIoN6TesGlBqpi--yVO_n-tnbHEg_KSGnLoYjy_i7aFuRfs3Shz3XwSfvDTwuCe2YtuQYA-SKiDG4QDxEF31xGO5fhNCsBmlvl6abGffpVtbgZ9Voh0wa8SLNEBybWJaktxfZF34vl1B_fFFghVoUhG7ztv_fTvUBaIAm99iIiJ4SLNjw"

    for event in events:
        res = scrape_party(event)
        print(json.dumps(res, indent=4))
        x = requests.patch(
            "https://android-gateway-clzdlli7.nw.gateway.dev/v1/parties",
            json={
                "party": res,
            },
            headers={
                "Authorization": "Bearer " + token,
            }
        )
        print(x)


def safe_list_get(l, idx, default):
    try:
        return l[idx]['src']
    except IndexError:
        return default


def get_all_events(url):
    response = network.get_with_proxy(url)
    soup = BeautifulSoup(response.content, 'html.parser')
    links = []

    for link in soup.find_all('a'):
        href = link.get('href')
        if "https" not in href:
            continue
        if "events" not in href:
            continue
        links.append(href)

    return links


def scrape_party(url):
    response = network.get_with_proxy(url)
    soup = BeautifulSoup(response.content, "html.parser")
    event_slug = url.split("/").pop()
    event_name = soup.find("h1").text
    event_organizer = soup.select("div span")[2].text
    event_description = soup.select("div.css-1sjbeek div")[1].text
    event_banner = safe_list_get(soup.select("div.r-1hfyk0a img"), 0, "")

    svgs = soup.select("div.r-k200y svg")
    event_location = svgs[-1].parent.parent.findChildren("span")[0].text  # Dernier svg = location
    (latitude, longitude) = utils.get_latitude_longitude_from_address(event_location)
    container = svgs[0].parent.parent.findChildren("div")[1].findChildren("div")

    event_end_date = 0
    event_start_date = 0

    if len(container) == 1:
        (start_date, end_date) = convert_to_epoch(container[0].text)
        event_start_date = start_date
        event_end_date = end_date
    else:
        (start_date, _) = convert_to_epoch(container[0].text)
        (end_date, _) = convert_to_epoch(container[1].text)
        event_start_date = start_date
        event_end_date = end_date

    return {
        "slug": event_slug,
        "name": event_name,
        "organizer": event_organizer,
        "location_name": event_location,
        "start_date": event_start_date,
        "end_date": event_end_date,
        "privacy": "PUBLIC",
        "description": event_description,
        "banner_url": event_banner,
        "latitude": latitude,
        "longitude": longitude,
    }


update_parties("")
