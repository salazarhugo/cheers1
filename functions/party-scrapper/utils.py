import re
from datetime import datetime, timedelta

import requests


def convert_to_epoch(input_str):
    date_pattern = r'(Mon|Tue|Wed|Thu|Fri|Sat|Sun) (\d+) (Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)'
    time_pattern = r'(\d+:\d+ [APap][Mm])'

    date_match = re.search(date_pattern, input_str)
    time_match = re.findall(time_pattern, input_str)

    day_name = date_match.group(1)
    day = int(date_match.group(2))
    month_name = date_match.group(3)

    year = 2023
    two_hours_in_seconds = 2 * 60 * 60

    day_dict = {
        'Mon': 0, 'Tue': 1, 'Wed': 2, 'Thu': 3, 'Fri': 4, 'Sat': 5, 'Sun': 6
    }

    month_dict = {
        'Jan': 1, 'Feb': 2, 'Mar': 3, 'Apr': 4, 'May': 5, 'Jun': 6,
        'Jul': 7, 'Aug': 8, 'Sep': 9, 'Oct': 10, 'Nov': 11, 'Dec': 12
    }

    day_num = day_dict[day_name]
    month_num = month_dict[month_name]

    date_obj = datetime(year, month_num, day)

    start_time_str = time_match[0]
    start_time_obj = datetime.strptime(start_time_str, '%I:%M %p').time()

    if len(time_match) > 1:
        end_time_str = time_match[1]
        end_time_obj = datetime.strptime(end_time_str, '%I:%M %p').time()
        if end_time_obj < start_time_obj:  # Handle overnight events
            end_time_obj = datetime.combine(date_obj + timedelta(days=1), end_time_obj)
        else:
            end_time_obj = datetime.combine(date_obj, end_time_obj)
        start_epoch = (datetime.combine(date_obj, start_time_obj) - datetime(1970, 1, 1)).total_seconds() - two_hours_in_seconds
        end_epoch = (end_time_obj - datetime(1970, 1, 1)).total_seconds() - two_hours_in_seconds
        return int(start_epoch), int(end_epoch)
    else:
        start_epoch = (datetime.combine(date_obj, start_time_obj) - datetime(1970, 1, 1)).total_seconds() - two_hours_in_seconds
        return int(start_epoch), None


input_str1 = "Sat 2 Sep from 2:00 PM to 12:30 AM"
input_str2 = "From Sun 20 Aug at 7:30 PM"
input_str3 = "To Sun 20 Aug at 7:30 PM"

# print(convert_to_epoch(input_str1))
# print(convert_to_epoch(input_str2))
# print(convert_to_epoch(input_str3))


def get_latitude_longitude_from_address(address):
    """
    Get the latitude and longitude from an address using the Mapbox API.

    Args:
        address: The address to geocode.
        mapbox_access_token: Your Mapbox access token.

    Returns:
        A tuple of latitude and longitude, or None if the address could not be geocoded.
    """

    url = "https://api.mapbox.com/geocoding/v5/mapbox.places/" + address + ".json?access_token=" + "pk.eyJ1Ijoic2FsYXphcmJyb2NrIiwiYSI6ImNqd3hxaDBleTFobGQ0Y2sxc3lpdnNwYXMifQ.TgYFgyclmL6pGXXANhnlsw"
    response = requests.get(url)

    if response.status_code == 200:
        data = response.json()
        if data["features"]:
            feature = data["features"][0]
            return feature["geometry"]["coordinates"]
        else:
            return None
    else:
        return None