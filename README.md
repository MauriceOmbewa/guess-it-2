# GUESS-IT-1
## Description
The guess-it-1 Program is a command-line application designed to predict a range for the next number based on a series of input numbers. The program reads numbers from standard input, processes them using statistical methods, and outputs a predicted range for the next number in the sequence. The range is determined by calculating the mean and standard deviation of the recent input numbers, ensuring a balanced prediction with a focus on accuracy and performance.

## Features
### Real-time Number Processing: 
Reads and processes numbers from standard input, predicting the range for the next number.
### Statistical Analysis: 
Utilizes statistical calculations, including mean and standard deviation, to determine prediction ranges.
### Performance Optimized: 
Designed to handle large datasets efficiently, ensuring quick and accurate predictions.
### Flexible Input Handling: 
Can be used with various data sources, making it adaptable for different use cases.

## Usage

1.  Clone this repository to your local machine.
```go
    git clone https://learn.zone01kisumu.ke/git/mombewa/guess-it-1.git
```
2.  Change directory to student directory
```go
    cd student
```
3.  Run the program
```go
    go run .
```
4.  Key in the data set. The program needs atleast two numbers to make a prediction and so for the first round you'll key in two numbers then the program prints the prediction range starting with the lowest point then a space and finally the uppermost point. For example:
```go
120
132
114 138
78
63 156
56
34 158
146
38 174
99
35 168

```

## Authors
- [@mombewa](https://learn.zone01kisumu.ke/git/mombewa)