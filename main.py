from selenium import webdriver
from webdriver_manager.chrome import ChromeDriverManager

driver = webdriver.Chrome(ChromeDriverManager().install())

# Navigate to the LinkedIn login page
driver.get("https://www.linkedin.com/login")

# Enter your email address and password
driver.find_element_by_id("username").send_keys("rafinabil39@gmail.com")
driver.find_element_by_id("password").send_keys("Z$jptLTvw4jd-&J")
# Submit the login form
driver.find_element_by_css_selector(".login__form_action_container button").click()

profile_url = "https://www.linkedin.com/in/example-profile"
driver.get(profile_url)
