# Quiz CLI
This is a simple command-line quiz application written in Go. The quiz reads questions from a CSV file, presents them to the user, and times the quiz to ensure answers are provided within a specified time limit. This project is a great example of using various Go programming features, including command-line flags, goroutines, channels, and working with CSV files.

## Features
1. Load quiz questions and answers from a CSV file.
2. Present questions and options in the terminal.
3. Enforce a time limit for the quiz.
4. Calculate and display the final score.

## Getting Started

### Prerequisites
- Go 1.16 or higher installed on your machine.

### Installation
1. Clone the repository
```
git clone https://github.com/mznrasil/quiz-cli.git
cd quiz-cli
```

2. Build the project
```
go build -o quiz
```

3. Run the quiz
- To start with default parameters
  ```
  ./quiz
  ```

- To add your own CSV file (format of questions can be seen in quiz_questions.csv file)
  ```
  ./quiz -file=your_questions.csv
  ```

- To add the timer (in seconds)
  ```
  ./quiz -time=30
  ```

- To add shuffling to questions
  ```
  ./quiz -shuffle
  ```

- Combine all the parameters
  ```
  ./quiz -file=your_questions.csv -time=60 -shuffle
  ```

### Concepts Learned
1. Command Line Flags
>Command-line flags are a way to pass arguments to your Go programs. In this project, the -file flag is used to specify the path to the CSV file containing the quiz questions.

2. Go Routines
>Goroutines are lightweight threads managed by the Go runtime. They allow concurrent execution of functions. In this project, a goroutine is used to handle the quiz questions while a timer runs concurrently.

3. Channels
>Channels in Go are used to communicate between goroutines. They can be thought of as pipelines that allow data to flow between different parts of the program. In this project, a channel is used to signal when the quiz is complete.

4. Working with CSV files
>CSV (Comma-Separated Values) files are a common format for storing tabular data. In this project, questions and answers are loaded from a CSV file using Go’s encoding/csv package.
