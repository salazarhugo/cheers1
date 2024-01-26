import os

import requests
import network
import utils
from utils import convert_to_epoch
import functions_framework
import scraper


@functions_framework.http
def scrape_parties(request):
    os.environ["TOTAL_PAGE"] = "10"
    os.environ["SCRAPE_URL"] = "https://shotgun.live/paris"
    os.environ["GATEWAY_URL"] = "https://web-gateway-r3a2dr4u4a-nw.a.run.app"
    os.environ["PROXY_USERNAME"] = "exologo"
    os.environ["PROXY_PASSWORD"] = "AqwXsz123987"
    os.environ["PROXY_HOST"] = "geo.iproyal.com"
    os.environ["PROXY_PORT"] = "32325"

    scrape_url = os.environ.get("SCRAPE_URL")
    gateway_url = os.environ.get("GATEWAY_URL")
    total_page = os.environ.get("TOTAL_PAGE")
    events = []

    for page in range(1, int(total_page)):
        url = f"{scrape_url}?page={page}"
        events.extend(scraper.get_all_events(url))

    success_count = 0
    for event in events:
        res = scraper.scrape_party(event)
        # print(json.dumps(res, indent=4))
        try:
            x = requests.put(
                f"{gateway_url}/v1/parties",
                json={
                    "party": res,
                },
                timeout=5,
            )
            success_count += 1
            print(x)
            print(res["name"])
            print(f"{success_count}/{len(events)} successful updates")
        except:
            print("An exception occurred")

    return f"Done {success_count}/{len(events)} successful updates"


scrape_parties("")
