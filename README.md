# Task Management Program in Go

This document contains the LaTeX source code for the Task Management Program.

```latex
\documentclass{article}
\usepackage[utf8]{inputenc}
\usepackage{hyperref}
\usepackage{listings}
\usepackage{xcolor}

% Define colors for code listings
\definecolor{codegreen}{rgb}{0,0.6,0}
\definecolor{codegray}{rgb}{0.5,0.5,0.5}
\definecolor{codepurple}{rgb}{0.58,0,0.82}
\definecolor{backcolour}{rgb}{0.95,0.95,0.92}

% Configure listings
\lstdefinestyle{mystyle}{
    backgroundcolor=\color{backcolour},   
    commentstyle=\color{codegreen},
    keywordstyle=\color{magenta},
    numberstyle=\tiny\color{codegray},
    stringstyle=\color{codepurple},
    basicstyle=\ttfamily\footnotesize,
    breakatwhitespace=false,         
    breaklines=true,                 
    captionpos=b,                    
    keepspaces=true,                 
    numbers=left,                    
    numbersep=5pt,                  
    showspaces=false,                
    showstringspaces=false,
    showtabs=false,                  
    tabsize=2
}

\lstset{style=mystyle}

\title{Task Management Program in Go}
\author{}
\date{}

\begin{document}

\maketitle

\section{Introduction}

This is a simple command-line task management program written in Go. It allows users to add, remove, list, and mark tasks as complete through an interactive menu. The program stores tasks in an in-memory list and provides basic task operations.

\section{Features}

\begin{itemize}
    \item \textbf{List Tasks}: View all tasks with their respective indices.
    \item \textbf{Add Task}: Add a new task to the list.
    \item \textbf{Delete Task}: Remove a task by its index.
    \item \textbf{Mark Task as Complete}: Mark a task as finished by its index.
    \item \textbf{Exit}: Exit the application.
\end{itemize}

\section{Usage}

To run this program, you'll need to have Go installed on your machine. After cloning or downloading the project, you can execute it using the following steps:

\subsection{Running the Program}

\begin{enumerate}
    \item \textbf{Install Go}: If you don't have Go installed, download and install it from \href{https://golang.org/dl/}{here}.
    
    \item \textbf{Clone or Download the Code}: Save the source code into a file called \texttt{main.go}.
    
    \item \textbf{Run the Program}: Open your terminal, navigate to the directory containing \texttt{main.go}, and run the following command:
    
    \begin{lstlisting}[language=bash]
    go run main.go
    \end{lstlisting}
    
    \item \textbf{Follow the Interactive Menu}: The program will present you with a list of options. Simply input the corresponding number to perform the desired action.
\end{enumerate}

\section{Example of Usage}

When prompted, select an option by typing the number:

\begin{lstlisting}
Please Input a Number Below
Option 1: List Tasks
Option 2: Add Task
Option 3: Delete Task
Option 4: Mark Task As Complete
Option 5: Exit
\end{lstlisting}

\subsection{Adding a Task}

If you choose option \textbf{2}, you'll be asked to enter a task name:

\begin{lstlisting}
Input Name:
\end{lstlisting}

\subsection{Listing Tasks}

If you choose option \textbf{1}, the program will list all current tasks:

\begin{lstlisting}
0 Buy groceries
1 Study Go
\end{lstlisting}

\subsection{Marking a Task as Complete}

If you choose option \textbf{4}, you'll be asked to input the index of the task to mark as complete:

\begin{lstlisting}
What Index?
\end{lstlisting}

\section{Code Overview}

The program consists of the following functions:

\begin{itemize}
    \item \textbf{addTask()}: Prompts the user for a task name and adds it to the task list.
    \item \textbf{removeTask()}: Prompts the user for the index of the task to be removed and deletes it from the task list.
    \item \textbf{listTask()}: Displays all the tasks currently in the task list.
    \item \textbf{completeTask()}: Marks a task as complete based on its index.
\end{itemize}

\section{License}

This project is free to use and distribute. Feel free to modify it to suit your needs.

\end{document}
```
