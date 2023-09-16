import os

import requests
from bs4 import BeautifulSoup
import network
import utils
from utils import convert_to_epoch
import functions_framework
import multiprocessing

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

    print(event_slug)
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

@functions_framework.http
def scrape_parties(request):
    os.environ["TOTAL_PAGE"] = "10"
    os.environ["SCRAPE_URL"] ="https://shotgun.live/paris"
    os.environ["GATEWAY_URL"] = "https://web-gateway-r3a2dr4u4a-nw.a.run.app"
    os.environ["PROXY_USERNAME"] = "exologo"
    os.environ["PROXY_PASSWORD"] = "AqwXsz123987"
    os.environ["PROXY_HOST"] ="geo.iproyal.com"
    os.environ["PROXY_PORT"] ="32325"

    scrape_url = os.environ.get("SCRAPE_URL")
    gateway_url = os.environ.get("GATEWAY_URL")
    total_page = os.environ.get("TOTAL_PAGE")
    eventsUrl = []
    events = []


    for page in range(1, int(total_page)):
        url = f"{scrape_url}?page={page}"
        eventsUrl.extend(get_all_events(url))

    success_count = 0

     # Create a pool of workers
    with multiprocessing.Pool(len(eventsUrl)) as pool:
        events = pool.map(scrape_party, eventsUrl)

    print(events)

    for event in events:
        try:
            x = requests.put(
                f"{gateway_url}/v1/parties",
                json={
                    "party": event,
                },
                timeout=3,
            )
            print(x)
            success_count += 1
        except:
            print("An exception occurred")

        print(f"{success_count}/{len(eventsUrl)} successful updates")

    return f"Done {success_count}/{len(eventsUrl)} successful updates"


scrape_parties("")