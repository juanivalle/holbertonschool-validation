build:	## Generate the website from the markdown and configuration files in the directory ./dist
	hugo > /dev/null 2>&1


clean:	## Cleanup the content of the directory ./dist
	rm -fr dist


post:	## Creates a post with given POST_NAME and POST_TITLE
	hugo new posts/$(POST_NAME).md
	sed -i '/title/c\title: "$(POST_TITLE)"' content/posts/$(POST_NAME).md


help:	## Provides information about all Makefile targets defined
	@fgrep "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/##//'
