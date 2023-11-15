from selenium import webdriver
from webdriver_manager.chrome import ChromeDriverManager
from selenium.webdriver.common.by import By
import time

import os
from dotenv import load_dotenv
load_dotenv()

EMAIL_SYSTEM = os.getenv("EMAIL_SYSTEM")
PASSWORD_SYSTEM = os.getenv("PASSWORD_SYSTEM")

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
        self.driver.get(profile_url)

    def search_job(self, job_name):
        self.driver.get(
            "https://www.linkedin.com/jobs/search/?keywords={}&location=Indonesia".format(
                job_name
            )
        )
        self.driver.implicitly_wait(5)
        
    def scroll_down(self):
        time.sleep(1)
        # Scroll down container with class name "jobs-search-results-list"
        element = self.driver.find_element(by=By.CLASS_NAME, value="jobs-search-results-list")

        # Please scroll smoothly to the bottom of the page
        self.driver.execute_script("arguments[0].scrollTop = arguments[0].scrollHeight", element)
        # self.driver.execute_script("arguments[0].scrollTop = arguments[0].scrollHeight", element)

    def scroll_down_smoothly(self):
        # Wait for 5 seconds before scrolling down to ensure the page has loaded
        time.sleep(5)

        # Find the element with the class name "jobs-search-results-list"
        element = self.driver.find_element(by=By.CLASS_NAME, value="jobs-search-results-list")

        # Execute JavaScript to scroll down the element smoothly
        self.driver.execute_script("""
            let currentScrollY = arguments[0].scrollTop;
            let targetScrollY = arguments[0].scrollHeight;

            let step = (targetScrollY - currentScrollY) / 10;

            function scrollAnimation() {
                currentScrollY += step;
                arguments[0].scrollTop = currentScrollY;

                if (currentScrollY < targetScrollY) {
                    requestAnimationFrame(scrollAnimation);
                }
            }

            requestAnimationFrame(scrollAnimation);
        """, element)


    def get_html(self):
        return self.driver.page_source

    def close(self):
        self.driver.quit()


# Example usage
if __name__ == "__main__":
    os.environ["PATH"] += os.pathsep + os.getcwd() + r"\ChromeDriver"
    print(os.environ["PATH"].split(";")[-1])
    chrome_driver_path = os.environ["PATH"].split(";")[-1]

    linkedin_bot = LinkedInBot(EMAIL_SYSTEM, PASSWORD_SYSTEM)

    try:
        linkedin_bot.login()
        # profile_url = "https://www.linkedin.com/in/"
        # linkedin_bot.visit_profile(profile_url)
        linkedin_bot.search_job("Software Engineer")
        linkedin_bot.scroll_down()
    finally:
        # linkedin_bot.close()
        input("End of the program!")
