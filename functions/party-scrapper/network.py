import os

import requests

USERNAME = os.environ.get("PROXY_USERNAME")
PASSWORD = os.environ.get("PROXY_PASSWORD")
HOST = os.environ.get("PROXY_HOST")
PORT = os.environ.get("PROXY_PORT")


def get_with_proxy(url):
    proxies = {
        'http': f'http://{USERNAME}:{PASSWORD}@{HOST}:{PORT}'
    }
    return requests.get(url, proxies=proxies)
