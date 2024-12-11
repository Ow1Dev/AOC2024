# Advent of Code 2024

Welcome to my Advent of Code 2024 solutions! This repository contains my solutions to each day's challenges. Below are the commands available to help you interact with the project.

## Available Commands

### List Available Commands
To list all available commands, run:

```bash
just --list
```

## Run Solution for a Given Day
To run the solution for a specific day (replace {{day}} with the day number), use the following command:

```bash
just run day={{day}}
```

Example for Day 1:

```bash
just run day_1
```

### Run Tests for a Given Day
To run the tests for a specific day, use the following command (replace {{day}} with the day number):

```bash
just test day={{day}}
```

Example for Day 1:

```bash
just test day_1
```

### Run Tests for All Solutions
To run the tests for all solutions, use:

```bash
just test-all
```

# Format the Code
To format the Go code in the project, use:

``` bash
just fmt
```

## Getting Started
Using Nix for Development Environment
This project includes a Nix flake for setting up a development environment with the necessary tools like Go and Just. Follow these steps to get started:

Clone the repository:

```bash
git clone https://github.com/yourusername/advent-of-code-2024.git
cd advent-of-code-2024
```

Ensure you have Nix installed and enabled for flakes. You can check by running:

```bash
nix --version
```

Enter the development shell:

```bash
nix develop
```

This will set up a shell with go and just installed.

Once inside the shell, you can run any of the commands using just (e.g., just run day=1 for Day 1 solution).

## Manually Installing Dependencies
Alternatively, you can install the required dependencies manually:

Ensure you have Go installed (version go1.20+ recommended). You can check by running:

```bash
go version
```
Install just (the task runner). You can follow the installation instructions at https://github.com/casey/just.

Run the desired command using just (e.g., just run day=1 for Day 1 solution).

License
This repository is licensed under the MIT License. See LICENSE for more information.
