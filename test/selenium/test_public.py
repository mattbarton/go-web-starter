"""
http://selenium-python.readthedocs.org/en/latest/installation.html
pip install selenium
"""

import unittest
import json
import sys
from selenium import webdriver
from selenium.webdriver.common.keys import Keys

baseURL = 'http://localhost:8080'

class PublicPagesTest(unittest.TestCase):

    def setUp(self):
        self.driver = webdriver.Firefox()

    def load_page(self, route):
        self.driver.get(baseURL + route)
        err = False
        if "All Rights Reserved" not in self.driver.page_source:
          err = True
          print >> sys.stderr, "Footer text not found at route {0}", route        
        # check the Javascript console for errors
        for entry in self.driver.get_log('browser'):
           if entry['level'] == 'SEVERE':
             print >> sys.stderr, "Javascript errors at route {0}: {1}", route, entry['message']
             err = True
        return err

    def test_all_pages_load(self):
        routes = ['']  #, '/aboutus', '/register', '/checkout', '/contact', '/login', '/signup', '/forgot']
        errors = map(self.load_page, routes)
        self.assertFalse(any(errors))

    def tearDown(self):
        self.driver.close()

if __name__ == "__main__":
    unittest.main()