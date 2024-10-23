# WordCounter Project

### My plan to get things done :

```mermaid
graph TD
   A[My Solution plan] --> C[Design System Architecture]:::green

   C --> D1[Concurrent Data Fetching]:::green
   C --> D2[Word Processing & Validation]:::green
   C --> D3[Concurrency & Rate Limiting]:::green
   C --> D4[Top 10 Words Aggregation]:::green

   D1 --> E1[Set Up Concurrency for Fetching URLs]:::green
   E1 --> F1[Implement Error Handling & Retry Logic]:::green
   F1 --> G1[Extract Text Content from URLs]:::green

   D2 --> E2[Load Word Bank & Preprocess Data]:::green
   E2 --> F2[Tokenize & Validate Words]:::green
   F2 --> G2[Count Word Frequencies]:::green

   D3 --> E3[Implement Concurrency with Thread Safety]:::green
   E3 --> F3[Handle Rate Limiting Time-based]:::green

   D4 --> E4[Sort Word Frequencies]:::green
   E4 --> F4[Extract Top 10 Words]:::green
   F4 --> G4[Format Output as Pretty JSON]:::green

   G1 --> H1[Concurrency Completed]:::green
   G2 --> H2[Word Processing Completed]:::green
   F3 --> H3[Rate Limiting Completed]:::green
   G4 --> H4[Output Aggregation Completed]:::green

   H1 --> I[Combine Results]:::green
   H2 --> I
   H3 --> I
   H4 --> I

   I --> J[Engineering Standards]:::green
   J --> K1[Unit Testing & Code Quality]:::yellow
   J --> K2[Documentation & Instructions]:::green
   J --> K3[Logging & Error Handling]:::yellow
   J --> K4[Containerization & Dockerization]:::green

   K1 --> L[Prepare Version Control Setup]:::green
   K2 --> L
   K3 --> L
   K4 --> L

   L --> M[Setup Continuous Integration CI]:::green
   M --> N[Deploy & Release Standards]:::green
   N --> O[End]:::green

%% Color Definitions
   classDef green fill:#000,color:#fff,stroke:#097969,stroke-width:3px;
   classDef yellow fill:#000,color:#fff,stroke:#FFC300,stroke-width:1px;

   ZZ[Done]:::green
   ZF[To-do]:::yellow
```


## How to Run

1. **Install Go**: Make sure you have Go installed. [Download Go](https://golang.org/dl/)

## Using Pre-Built Docker Image

You can pull and run the pre-built Docker image from Docker Hub:

1. **Pull the image from Docker Hub**:
   ```bash
   docker pull seferovramin7/wordcounterproject
   ```

2. **Run the container**:
   ```bash
   docker run --rm seferovramin7/wordcounterproject
   ```

### Option 2 : To build the project locally : 


1. **Build the project**:
   ```bash
   go build -o app
   ```
2. **Run the project**:
   ```bash
   go run main.go
   ```
3. **Build and run with Docker**:
   ```bash
   docker build -t wordcounterproject .
   docker run --rm wordcounterproject
   ```


## Project Structure

- `main.go`: Main application logic
- `utils/`: Utility functions for fetching, processing, and word validation
- `wordbank.txt`: File containing the list of valid words
