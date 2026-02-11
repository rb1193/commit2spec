# TDD Development Methodology

This project follows a strict TDD methodology. We write the tests first, then the minimum amount of code needed to pass the tests, then refactor. We start our work from the top down, and only add tests at lower levels of the test pyramid where branching or looping makes adding test coverage at the higher level expensive.

Use dependency injection to enable us to inject test doubles to represent IO resources such as files and HTTP request and responses.

You have a number of subagents available which specialise in the three activities of TDD. Orchestrate them to deliver code in a TDD fashion.
