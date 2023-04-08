package preprocessor

// this module contains logic for replacing snippets / macros in the future, such as:
// - @include test.md => gets replaced with the markdown content of the test.md file
// - @today => gets replaced by the current date
// - @shell "command" => gets replaced by the output of the given command
