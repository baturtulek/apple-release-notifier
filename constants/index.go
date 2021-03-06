package constants

import "github.com/baturtulek/apple-release-notifier/utils"

// APPLE_RELEASE_PAGE_URL APPLE's Release Web Page
var APPLE_RELEASE_PAGE_URL = "https://developer.apple.com/news/releases"

// LAST_CRAWL_RELEASES_FILE File Contains Last Crawled Releases
var LAST_CRAWL_RELEASES_FILE = utils.AppRootPath + "/last_release_data.json"

// MAIL_CONTACTS_FILE File Contains Mail Addresses of the Clients
var MAIL_CONTACTS_FILE = utils.AppRootPath + "/contacts"
