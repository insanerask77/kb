# How to use PostgreSQL to store and query vector embeddings

[#postgres](https://dev.to/t/postgres)[#vectorembeddings](https://dev.to/t/vectorembeddings)[#cosinesimilarity](https://dev.to/t/cosinesimilarity)[#semanticsearch](https://dev.to/t/semanticsearch)

Vector embeddings are powerful representations of data points in a high-dimensional space, and they are widely used in Natural Language Processing (NLP) and Machine Learning. Storing and efficiently querying these embeddings is crucial for building scalable and performant applications.

In this blog post, I will explain how to use [pgvector](https://github.com/pgvector/pgvector), a PostgreSQL extension, to store and query vector embeddings efficiently. All of the code for this article can be found in [the companion GitHub repository](https://github.com/stephenc222/example-postgres-vector-embeddings).

## Prerequisites

Here are the items that you need to be somewhat familiar with before continuing with this tutorial:

- Basic knowledge of Python and SQL (specifically [PostgreSQL](https://www.postgresql.org/)'s SQL).
- [psycopg2](https://pypi.org/project/psycopg2/), [PyTorch](https://pytorch.org/get-started/locally/) and [Transformers](https://huggingface.co/docs/transformers/installation) installed in your Python environment.
- [Docker](https://docs.docker.com/get-started/) installed on your system.
- [Docker Compose](https://docs.docker.com/compose/) also installed on your system.

First, I'll explain how to setup a Postgres database with the `pgvector` extension using docker and docker-compose.

## Setting Up the Database

### Step 1: Create a Dockerfile

In your project directory, create a file named `Dockerfile` and add the following content to it:

```
# Use the official Postgres image as a base image
FROM postgres:latest

# Set environment variables for Postgres
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=mydb

# Install the build dependencies
USER root
RUN apt-get update && apt-get install -y \
    build-essential \
    git \
    postgresql-server-dev-all \
    && rm -rf /var/lib/apt/lists/*

# Clone, build, and install the pgvector extension
RUN cd /tmp \
    && git clone --branch v0.5.0 https://github.com/pgvector/pgvector.git \
    && cd pgvector \
    && make \
    && make install
```



### Step 2: Create a docker-compose.yml File

In your project directory, create a file named `docker-compose.yml` and add the following content to it:

```
version: "3"

services:
  postgres:
    build: .
    ports:
      - "5432:5432"
    volumes:
      - ./data:/var/lib/postgresql/data
      - ./init_pgvector.sql:/docker-entrypoint-initdb.d/init_pgvector.sql
    environment:
      POSTGRES_USER: myuser
      POSTGRES_PASSWORD: mypassword
      POSTGRES_DB: mydb
```



### Step 3: Create Initialization SQL Script

Next, create a file called `init_pgvector.sql` and add the following to it:

```
-- Install the extension we just compiled

CREATE EXTENSION IF NOT EXISTS vector;

/*
For simplicity, we are directly adding the content into this table as
a column containing text data. It could easily be a foreign key pointing to
another table instead that has the content you want to vectorize for
semantic search, just storing here the vectorized content in our "items" table.

"768" dimensions for our vector embedding is critical - that is the
number of dimensions our open source embeddings model output, for later in the
blog post.
*/

CREATE TABLE items (id bigserial PRIMARY KEY, content TEXT, embedding vector(768));
```



This SQL script will be used in the next section about building and running the docker container.

### Step 3: Build and Run the Docker Container

Navigate to your project directory in the terminal and run the following command to build and run the Docker container using the Dockerfile and `docker-compose.yml` you just created:

```
docker compose up
```



With our pgvector extension compiled and loaded into our database, and our database up and accepting connections, let's move onto the Python app.

## Creating Embeddings with Transformers

For the sake of keeping this blog post shorter, I won't cover in detail what vector embeddings are. I do offer a free newsletter, and I've written an [introduction to vector embeddings](https://stephencollins.tech/newsletters/understanding-vector-embeddings-secret-sauce-behind-language-models) that you are free to checkout - and then come back!

Let's move on to how we are creating embeddings.

### Introduction to embedding_util

The `embedding_util` module is how we are generating embeddings in this tutorial. This module leverages a transformer model, specifically from `thenlper/gte-base` (more info [here](https://huggingface.co/thenlper/gte-base)), to transform textual data into meaningful, semantically rich vector representations (our vector embeddings).

### Code Overview

I'll share the code that we are using to create the vector embeddings in our `embedding_util.py` file, and then explain it.

```
import os
import json
from transformers import AutoTokenizer, AutoModel
import torch
import torch.nn.functional as F

# We won't have competing threads in this example app
os.environ["TOKENIZERS_PARALLELISM"] = "false"


# Initialize tokenizer and model for GTE-base
tokenizer = AutoTokenizer.from_pretrained('thenlper/gte-base')
model = AutoModel.from_pretrained('thenlper/gte-base')


def average_pool(last_hidden_states: Tensor, attention_mask: Tensor) -> Tensor:
    last_hidden = last_hidden_states.masked_fill(
        ~attention_mask[..., None].bool(), 0.0)
    return last_hidden.sum(dim=1) / attention_mask.sum(dim=1)[..., None]


def generate_embeddings(text, metadata={}):
    combined_text = " ".join(
        [text] + [v for k, v in metadata.items() if isinstance(v, str)])

    inputs = tokenizer(combined_text, return_tensors='pt',
                       max_length=512, truncation=True)
    with torch.no_grad():
        outputs = model(**inputs)

    attention_mask = inputs['attention_mask']
    embeddings = average_pool(outputs.last_hidden_state, attention_mask)

    embeddings = F.normalize(embeddings, p=2, dim=1)

    return json.dumps(embeddings.numpy().tolist()[0])
```



#### 1. **Importing Necessary Libraries**

We start by importing necessary libraries and modules, including `json` for serialization, `AutoTokenizer` and `AutoModel` from the transformers library for tokenization and model loading, `torch` from PyTorch for tensor operations, and `torch.nn.functional` for additional functionalities like normalization.

#### 2. **Setting Environment Variable**

```
os.environ["TOKENIZERS_PARALLELISM"] = "false"
```



This line disables tokenizers parallelism to prevent any threading issues in this example app.

*NOTE: If you are using this inside an asynchronous web framework like [fastAPI](https://fastapi.tiangolo.com/), you'll want to set this to `"true"` instead.*

#### 3. **Initializing Tokenizer and Model**

```
tokenizer = AutoTokenizer.from_pretrained('thenlper/gte-base')
model = AutoModel.from_pretrained('thenlper/gte-base')
```



Here, we are initializing the tokenizer and model for `GTE-base`, enabling us to convert text into embeddings subsequently.

#### 4. **Average Pooling Function**

```
def average_pool(last_hidden_states: Tensor, attention_mask: Tensor) -> Tensor:
    last_hidden = last_hidden_states.masked_fill(
        ~attention_mask[..., None].bool(), 0.0)
    return last_hidden.sum(dim=1) / attention_mask.sum(dim=1)[..., None]
```



This function performs average pooling on the last hidden states of the transformer model, using the attention mask to handle padded tokens. It returns the average pooled tensor, which represents the semantic essence of the input text.

#### 5. **Embedding Generation Function**

```
def generate_embeddings(text, metadata={}):
    combined_text = " ".join(
        [text] + [v for k, v in metadata.items() if isinstance(v, str)])

    inputs = tokenizer(combined_text, return_tensors='pt',
                       max_length=512, truncation=True)
    with torch.no_grad():
        outputs = model(**inputs)

    attention_mask = inputs['attention_mask']
    embeddings = average_pool(outputs.last_hidden_state, attention_mask)

    embeddings = F.normalize(embeddings, p=2, dim=1)

    return json.dumps(embeddings.numpy().tolist()[0])
```



This is where the main action happens. This function takes in text and optional metadata, tokenizes the input, passes it through the transformer model, and obtains the last hidden states. It then applies the `average_pool` function to get the final embedding, normalizes it, and returns it as a JSON string, ready to be stored in the database.

Understanding the embedding generation process is pivotal for implementing effective and efficient semantic searches. This module, powered by the transformers library, facilitates the creation of rich, meaningful vector representations of text, enabling nuanced similarity searches when integrated with `pgvector` and PostgreSQL.

## Integration with PostgreSQL for Storage and Querying of Embeddings

With the explanation of the code behind how we are creating the embeddings, let's move onto the core of our Python app. Here, in `app.py`, we consolidate all of our components to insert example sentences into our PostgreSQL database, and then execute a cosine similarity search using `pgvector` to rank and retrieve the most similar items to a given query.

### Bringing it All Together: `app.py`

Now, I'll share the whole `app.py` file, then explain it part by part:

```
import psycopg2
from embedding_util import generate_embeddings

def run():

    # Establish a connection to the PostgreSQL database
    conn = psycopg2.connect(
        user="myuser",
        password="mypassword",
        host="localhost",
        port=5432,  # The port you exposed in docker-compose.yml
        database="mydb"
    )

    # Create a cursor to execute SQL commands
    cur = conn.cursor()

    try:
        sentences = [
            "A group of vibrant parrots chatter loudly, sharing stories of their tropical adventures.",
            "The mathematician found solace in numbers, deciphering the hidden patterns of the universe.",
            "The robot, with its intricate circuitry and precise movements, assembles the devices swiftly.",
            "The chef, with a sprinkle of spices and a dash of love, creates culinary masterpieces.",
            "The ancient tree, with its gnarled branches and deep roots, whispers secrets of the past.",
            "The detective, with keen observation and logical reasoning, unravels the intricate web of clues.",
            "The sunset paints the sky with shades of orange, pink, and purple, reflecting on the calm sea.",
            "In the dense forest, the howl of a lone wolf echoes, blending with the symphony of the night.",
            "The dancer, with graceful moves and expressive gestures, tells a story without uttering a word.",
            "In the quantum realm, particles flicker in and out of existence, dancing to the tunes of probability.",
        ]

        # Insert sentences into the items table
        for sentence in sentences:
            embedding = generate_embeddings(sentence)
            cur.execute(
                "INSERT INTO items (content, embedding) VALUES (%s, %s)",
                (sentence, embedding)
            )

        # Example query
        query = "Give me some content about the ocean"
        query_embedding = generate_embeddings(query)

        # Perform a cosine similarity search
        cur.execute(
            """SELECT id, content, 1 - (embedding <=> %s) AS cosine_similarity
               FROM items
               ORDER BY cosine_similarity DESC LIMIT 5""",
            (query_embedding,)
        )

        # Fetch and print the result
        print("Query:", query)
        print("Most similar sentences:")
        for row in cur.fetchall():
            print(
                f"ID: {row[0]}, CONTENT: {row[1]}, Cosine Similarity: {row[2]}")

    except Exception as e:
        print("Error executing query", str(e))
    finally:
        # Close communication with the PostgreSQL database server
        cur.close()
        conn.close()

# This check ensures that the function is only run when the script is executed directly, not when it's imported as a module.
if __name__ == "__main__":
    run()
```



#### 1. **Connection Setup**

After importing dependencies, we need to establish a connection to our Postgres and create a cursor that we can use to execute queries with:

```
    # Establish a connection to the PostgreSQL database
    conn = psycopg2.connect(
        user="myuser",
        password="mypassword",
        host="localhost",
        port=5432,  # The port you exposed in docker-compose.yml
        database="mydb"
    )

    # Create a cursor to execute SQL commands
    cur = conn.cursor()
```



#### 2. **Embedding Insertion**

Iterating over a list of sentences, generating their corresponding embeddings, and inserting them into the `items` table in the database, inside the `try..except` block:

```
sentences = [
    "A group of vibrant parrots chatter loudly, sharing stories of their tropical adventures.",
    "The mathematician found solace in numbers, deciphering the hidden patterns of the universe.",
    "The robot, with its intricate circuitry and precise movements, assembles the devices swiftly.",
    "The chef, with a sprinkle of spices and a dash of love, creates culinary masterpieces.",
    "The ancient tree, with its gnarled branches and deep roots, whispers secrets of the past.",
    "The detective, with keen observation and logical reasoning, unravels the intricate web of clues.",
    "The sunset paints the sky with shades of orange, pink, and purple, reflecting on the calm sea.",
    "In the dense forest, the howl of a lone wolf echoes, blending with the symphony of the night.",
    "The dancer, with graceful moves and expressive gestures, tells a story without uttering a word.",
    "In the quantum realm, particles flicker in and out of existence, dancing to the tunes of probability.",
]

# Insert sentences into the items table
for sentence in sentences:
    embedding = generate_embeddings(sentence)
    cur.execute(
        "INSERT INTO items (content, embedding) VALUES (%s, %s)",
        (sentence, embedding)
    )
```



#### 3. **Embedding Retrieval and Similarity Search**

Creating an embedding for a sample query and leveraging the `<=>` operator provided by `pgvector` allows us to perform a cosine similarity search.

It is crucial to note that `pgvector` actually calculates the cosine distance, which is different from cosine similarity.

Cosine distance measures the cosine of the angle between two non-zero vectors, whereas cosine similarity measures the cosine of the angle between two vectors, projecting the amount they overlap. Here's another [explanation of cosine distance and cosine similarity](https://medium.datadriveninvestor.com/cosine-similarity-cosine-distance-6571387f9bf8).

To obtain the cosine similarity from the cosine distance calculated by `pgvector`, we subtract the cosine distance from 1, returning results ordered by similarity, as illustrated below:

```
        # Example query
        query = "Give me some content about the ocean"
        query_embedding = generate_embeddings(query)

        # Perform a cosine similarity search
        cur.execute(
            """SELECT id, content, 1 - (embedding <=> %s) AS cosine_similarity
               FROM items
               ORDER BY cosine_similarity DESC LIMIT 5""",
            (query_embedding,)
        )

        # Fetch and print the result
        print("Query:", query)
        print("Most similar sentences:")
        for row in cur.fetchall():
            print(
                f"ID: {row[0]}, CONTENT: {row[1]}, Cosine Similarity: {row[2]}")
```



Our final output logged to the console should look something like this:

```
Query: Give me some content about the ocean
Most similar sentences:
ID: 7, CONTENT: The sunset paints the sky with shades of orange, pink, and purple, reflecting on the calm sea., Cosine Similarity: 0.8009044169301547
ID: 5, CONTENT: The ancient tree, with its gnarled branches and deep roots, whispers secrets of the past., Cosine Similarity: 0.760971262397107
ID: 1, CONTENT: A group of vibrant parrots chatter loudly, sharing stories of their tropical adventures., Cosine Similarity: 0.7582434704013556
ID: 8, CONTENT: In the dense forest, the howl of a lone wolf echoes, blending with the symphony of the night., Cosine Similarity: 0.7446566376294829
ID: 2, CONTENT: The mathematician found solace in numbers, deciphering the hidden patterns of the universe., Cosine Similarity: 0.7399669003526144
```



## Conclusion

This blog post illustrated how to efficiently store and query vector embeddings in PostgreSQL using pgvector, opening up advanced possibilities in semantic search and analysis for Natural Language Processing and Machine Learning.

The companion code repository, with all of the code explained in this tutorial in one place, can be found [here](https://github.com/stephenc222/example-postgres-vector-embeddings).