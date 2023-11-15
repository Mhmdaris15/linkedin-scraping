from python import Python

let webdriver = Python.import_module("webdriver")
let requests = Python.import_module("requests")
let BeautifulSoup = Python.import_module("BeautifulSoup")

# Initialize Chrome driver
driver = webdriver.Chrome()

# Load environment variables
email_system: str = os.getenv("EMAIL_SYSTEM")
password_system: str = os.getenv("PASSWORD_SYSTEM")

# Define LinkedInBot class
class LinkedInBot:
  def __init__(self, username: str, password: str) -> None:
    self.driver: webdriver.Driver = driver
    self.username: str = username
    self.password: str = password

  def login(self) -> None:
    self.driver.get("https://www.linkedin.com/login")
    self.driver.find_element(webdriver.By.ID, "username").send_keys(self.username)
    self.driver.find_element(webdriver.By.ID, "password").send_keys(self.password)
    login_button = self.driver.find_element(webdriver.By.CSS_SELECTOR, '[data-litms-control-urn="login-submit"]')
    login_button.click()
    self.driver.implicitly_wait(10)

  def visit_profile(self, profile_url: str) -> None:
    self.driver.get(profile_url)

  def search_job(self, job_name: str) -> None:
    job_search_url: str = f"https://www.linkedin.com/jobs/search/?keywords={job_name}&location=Indonesia"
    self.driver.get(job_search_url)
    self.driver.implicitly_wait(5)

  def soup_jobs(self, endpoint: str) -> BeautifulSoup:
    response = requests.get(endpoint)
    return BeautifulSoup(response.content, "html.parser")

  def array_jobs(self, soup: BeautifulSoup) -> list[webdriver.WebElement]:
    job_cards: list[webdriver.WebElement] = soup.find_all("div", class_="base-card")
    return job_cards

  def print_jobs(self, soup: BeautifulSoup) -> None:
    job_cards = self.array_jobs(soup)
    for job_card in job_cards:
      print(job_card.get_text())

  def get_html(self) -> str:
    return self.driver.page_source

  def scroll_down(self, scroll_times: int = 1) -> None:
    for _ in range(scroll_times):
      self.driver.execute_script("window.scrollTo(0, document.body.scrollHeight);")
      self.driver.implicitly_wait(2)

  def close(self) -> None:
    self.driver.quit()

# Example usage
if __name__ == "__main__":
  chrome_driver_path = os.getcwd() + r"\ChromeDriver"
  linkedin_bot = LinkedInBot(email_system, password_system)

  try:
    linkedin_bot.login()
    linkedin_bot.search_job("Software Engineer")
    linkedin_bot.scroll_down(10)
    soup = linkedin_bot.soup_jobs(linkedin_bot.driver.current_url)
    print(linkedin_bot.array_jobs(soup))
    linkedin_bot.print_jobs(soup)
  except Exception as e:
    print(e)
    linkedin_bot.scroll_down()

  finally:
    input("End of the program!")
