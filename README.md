# LinkedIn Job Scraping

## Overview

This repository contains a Python script that utilizes the Selenium webdriver to scrape job listings from LinkedIn. The script retrieves job titles, company names, locations, and descriptions, and saves the data in a CSV file.

## Usage

1. Clone this repository and install the required dependencies using pip install -r requirements.txt.
2. Run the script by executing python `main.py` in your terminal or command prompt.
3. The script will prompt you to enter a LinkedIn search query (e.g., "software engineer").
4. Once you've entered your search query, the script will retrieve the top 100 job listings and save them to a CSV file named output.csv.
5. You can modify the script to suit your needs, such as changing the search query, increasing the number of jobs retrieved, or saving the data to a different file format.

## Dependencies

- selenium
- chromium-webdriver
- pandas
- requests

## Configuration

You can configure the script by modifying the following variables at the beginning of the `main.py` file:

- `search_query`: Enter the LinkedIn search query you want to use.
- `num_jobs`: Specify the number of jobs you want to retrieve. Default is 100.
- `output_file`: Provide the name of the output file where the scraped data should be saved. Default is output.csv.

## Troubleshooting

If you encounter any issues while running the script, check the console output for error messages. Common errors include:

- `chromium-webdriver` not installed: Install the dependency by running pip install `chromium-webdriver`.
- `selenium` not installed: Install the dependency by running `pip install selenium`.
- LinkedIn page not loading: Check your internet connection and try reloading the page.
- Too many requests: Wait for some time before running the script again.

## Contributing

Feel free to contribute to this project by opening pull requests with improvements, bug fixes, or new features.

## License

MIT License

Copyright (c) Muhammad Aris Septanugroho

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

_The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software._

**THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.**
