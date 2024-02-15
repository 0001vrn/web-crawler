# Web crawler 
This repository demonstrates high level design and low level implementation of web crawler application. 

# Functional requirements
Your task is to write a web crawler (also known as a spider bot), a robot that
copies the Internet.

To keep things simple write an implementation that will run on a single machine and
focus only on crawling aspect. To get you started you have 2 components ready:
- "fetch" function that retrieves HTML content of a page at given URL as a string.
- "parse" function that returns a list of links found on the website.

We need to write a "startCrawling" function that starts the copying the internet given the seed URL.

# Design proposals

## BFS (Breadth-first search)
### Pros
1. BFS is simple and easy to implement, and it can cover a large portion of the web in a relatively short time
2. BFS is suitable for general-purpose web crawling such as building a search engine index

### Cons
1. Consuming a lot of memory, bandwidth, and disk space
2. Being prone to getting stuck in low-quality or irrelevant pages
3. Not optimal for specific or focused web crawling tasks


## DFS (Depth-first search)
### Pros
1. DFS is more efficient and less resource-intensive than BFS, as it does not need to store all the URLs in memory or disk
2. DFS is suitable for focused web crawling, such as collecting data from a specific domain or topic

### Cons
1. Being slow to cover a large portion of the web
2. Being susceptible to getting trapped in loops or dead ends
3. It may not be ideal for general-purpose web crawling tasks

## Best-first search
### Pros
1. BFS is more intelligent and adaptive than BFS or DFS, as it can dynamically adjust the crawling order according to the relevance, quality, or freshness of the pages
2.  BFS is suitable for adaptive web crawling, such as finding new or updated pages, or crawling pages that match a specific query or interest
### Cons
1. Designing and evaluating the heuristic function, as well as balancing the exploration and exploitation trade-off

# Preferred Solution
However, for this specific problem, it is preferable to traverse all links of a particular page in a natural order(explore the neighbor nodes first, before moving to the next-level neighbors). 
Also, as per the problem statement we need a simple web crawler. Thus, `BFS (Breadth-first search)` or `Level order traversal` is ideal for this case.