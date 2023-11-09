import requests
from bs4 import BeautifulSoup
import math

target_url = "https://www.linkedin.com/jobs-guest/jobs/api/jobPosting/3698543802"
number_of_loops = math.ceil(400/25)
print("Number of loops: ", number_of_loops)