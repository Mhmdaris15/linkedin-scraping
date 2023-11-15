import os
from selenium import webdriver
from selenium.webdriver.common.by import By
import requests
from bs4 import BeautifulSoup


class LinkedInBot:
    def __init__(self, username, password):
        # self.chrome_driver_path = self._get_chrome_driver_path()
        self.driver = webdriver.Chrome()
        self.username = username
        self.password = password

    def _get_chrome_driver_path(self):
        os.environ["PATH"] += os.pathsep + os.getcwd() + r"\ChromeDriver"
        return os.environ["PATH"].split(";")[-1]

    def login(self):
        self.driver.get("https://www.linkedin.com/login")
        self.driver.find_element(By.ID, "username").send_keys(self.username)
        self.driver.find_element(By.ID, "password").send_keys(self.password)
        login_button = self.driver.find_element(
            By.CSS_SELECTOR, '[data-litms-control-urn="login-submit"]'
        )
        login_button.click()
        self.driver.implicitly_wait(10)

    def visit_profile(self, profile_url):
        print("Hello")
        self.driver.get(profile_url)

    def search_job(self, job_name):
        self.driver.get(
            "https://www.linkedin.com/jobs/search/?keywords={}&location=Indonesia".format(
                job_name
            )
        )

    def soup_jobs(self, endpoint):
        response = requests.get(endpoint)
        soup = BeautifulSoup(response.content, "html.parser")
        return soup

    def array_jobs(self, soup):
        job_cards = soup.find_all("div", class_="base-card")
        for card in job_cards:
            print(card.get_text())
        return job_cards

    def print_jobs(self, soup):
        job_cards = self.array_jobs(soup)
        for job_card in job_cards:
            print(job_card.get_text())
            # job_title = job_card.find("h3", class_="base-search-card__title").text
            # company_name = job_card.find("h4", class_="base-search-card__subtitle").text
            # location = job_card.find("span", class_="job-search-card__location").text
            # job_link = job_card.find("a", class_="base-card__full-link")["href"]
            # print(job_title)
            # print(company_name)
            # print(location)
            # print(job_link)
            # print()

    def get_html(self):
        return self.driver.page_source

    def scroll_down(self, scroll_times=1):
        for i in range(scroll_times):
            self.driver.execute_script(
                "window.scrollTo(0, document.body.scrollHeight);"
            )
            self.driver.implicitly_wait(2)

    def close(self):
        self.driver.quit()


# Example usage
if __name__ == "__main__":
    username = "rafinabil39@gmail.com"
    password = "Z$jptLTvw4jd-&J"
    os.environ["PATH"] += os.pathsep + os.getcwd() + r"\ChromeDriver"
    # print(os.environ["PATH"].split(";")[-1])
    chrome_driver_path = os.environ["PATH"].split(";")[-1]

    linkedin_bot = LinkedInBot(username, password)

    try:
        linkedin_bot.login()
        # profile_url = "https://www.linkedin.com/in/"
        # linkedin_bot.visit_profile(profile_url)
        linkedin_bot.search_job("Software Engineer")
        linkedin_bot.scroll_down(10)
        soup = linkedin_bot.soup_jobs(linkedin_bot.driver.current_url)
        # print(linkedin_bot.driver.current_url)
        # print(f"Soup : %s" % soup)
        print(f"Array Jobs : %s" % linkedin_bot.array_jobs(soup))
        linkedin_bot.print_jobs(soup)
    except Exception as e:
        print(e)
    finally:
        # linkedin_bot.close()
        input("End of the program!")
