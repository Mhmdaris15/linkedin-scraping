import os
from selenium import webdriver
from selenium.webdriver.common.by import By


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

    def get_html(self):
        return self.driver.page_source

    def close(self):
        self.driver.quit()


# Example usage
if __name__ == "__main__":
    username = "rafinabil39@gmail.com"
    password = "Z$jptLTvw4jd-&J"
    os.environ["PATH"] += os.pathsep + os.getcwd() + r"\ChromeDriver"
    print(os.environ["PATH"].split(";")[-1])
    chrome_driver_path = os.environ["PATH"].split(";")[-1]

    linkedin_bot = LinkedInBot(username, password)

    try:
        linkedin_bot.login()
        # profile_url = "https://www.linkedin.com/in/"
        # linkedin_bot.visit_profile(profile_url)
        linkedin_bot.search_job("Software Engineer")
    finally:
        # linkedin_bot.close()
        input("End of the program!")
