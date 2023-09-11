import requests

USERNAME = "exologo"
PASSWORD = "AqwXsz123987"
HOST = "geo.iproyal.com"
PORT = "32325"


def get_with_proxy(url):
    proxies = {
        'http': f'http://{USERNAME}:{PASSWORD}@{HOST}:{PORT}'
    }
    return requests.get(url, proxies=proxies)
