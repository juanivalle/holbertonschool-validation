# Development Lifecycle with Makefile
This project aims at showing use cases where a DevOps mindset is bringing value to a software project by automating it, which decreases the amount of manual work and increases the development speed. It focuses on why automation is useful and helps speeding a development lifecycle.

## Prerequisites:
- GNU Make in version 3.81+
- Git (command line) in version 2+
- Go Hugo v0.80+

## Lifecycle:
- **build**: Generate the website from the markdown and configuration files in the directory ./dist
- **clean**: Cleanup the content of the directory ./dist
- **post**: Creates a post with given POST_NAME and POST_TITLE
- **help**: Provides information about all Makefile targets defined