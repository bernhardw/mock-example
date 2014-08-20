# Mocking Packages

While working on Go Examples I stumbled while writing tests for the Example model. Those functions and methods use an internal package to retrieve the examples from Github to further process them.
But I didn't want to hit the Github API every time I ran the tests.

This repository shows a simple example of the solution we ended up with.

* Branch `#basic` shows a basic, untestable example.
* Branch `#master` shows the solution.

## The Example

The example gets all public repositories from Ongoing.io and prints them to the console. To fetch the repositories from Github, we will use the `service` package that we want to mock when testing the `GetRepos()` function in `mock.go`.
