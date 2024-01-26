import requests
from bs4 import BeautifulSoup
import network
import utils
from utils import convert_to_epoch


def scrape_party(url):
    response = network.get_with_proxy(url)
    soup = BeautifulSoup(response.content, "html.parser")
    event_slug = url.split("/").pop()
    event_name = soup.find("h1").text
    event_organizer = soup.select("div span")[2].text
    event_description = soup.select("div.css-1sjbeek div")[1].text
    event_banner = utils.safe_list_get(soup.select("div.r-1hfyk0a img"), 0, "")

    svgs = soup.select("div.r-k200y svg")
    event_location = svgs[-1].parent.parent.findChildren("span")[0].text  # Dernier svg = location
    (longitude, latitude) = utils.get_latitude_longitude_from_address(event_location)
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

