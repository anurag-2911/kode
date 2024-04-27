## Interview rounds
 - Recruiter screening/ Hiring Manager screening: Whether the skill set matches the job requirements and soft skills check for the current team culture: 30 mins to 1 hr
 - Live Coding(sharing desktop and coding): 30 mins to 1 hr
    -  initial coding
    -  handling edge cases
    -  UTs
 - Or Online coding by hacker rank, codility, etc (Live coding is easier as the approach can be demonstrated even if the problem is not exactly solved)
 - Assignments: typically 1 week is given for assignments
    - Readable, modular, production-grade code
    - Edge cases
    - Secure code: Like handling of configuration of secrets
    - UTs
    - Docker file
    - Folder structure and maintainable code
    - Formatting
    - 
 - Conversational interview(no coding or design): Technical questions about skillset and projects: 1 hr 
 - System Design (1 or 2 rounds): 1 hr
 - Executive interview(no coding or design): technical conversational interview by higher management: 1 hr
 - Culture fit or behavioral interview:30 mins to 1 hr
 - Recruiter round for salary discussion
## System Design
  - Approach
    - Steps:
      - Understand the problem and clarify all the requirements
      - Write high-level design and high-level diagrams
      - Implement all the components of high-level diagrams in detail
    - Resources:
      - https://bytebytego.com/courses/system-design-interview/foreword
      - https://www.educative.io/courses/grokking-modern-system-design-interview-for-engineers-managers
      - Implement small projects for each design problem and refine them to perfection.
      - Tool: https://excalidraw.com/: For drawing System Design problems
        
  - Concepts needed for System Design:
    - API Gateway | Load Balancer | Vertical Scaling and Horizontal Scaling
    - Web Server
      - Stateful and Stateless architecture
    - Application Server
    - Database : SQL | NoSQL | GraphDB | DocumentDB | Replication | Partitioning(Sharding) 
    - Cache: Redis | MemCache
    - CDN
    - Message Queues |  Event-driven architecture
    - Microservices and communication patterns: REST | gRPC | GraphQL | Message Queues
    - Logging, metrics, automation
    - Back of the envelope Calculation: QPS | Storage | Number of Servers
    - 
  - High-Level Design Problems (HLDs)
     - URL Shortener
     - Unique ID Generator in Distributed system:
        - multi-master replication, UUID(GUID), ticket server, and Twitter snowflake-like unique ID generator
     - Rate Limiter
        - Token Bucket Algorithm
     - Storage based Design
        - S3-like storage system
        - CDN
        - Google Drive
     - Content streaming-based design
         - You Tube
         - Netflix
     -  E-commerce
         - Design Hotel reservation
         - Design Shopping websites like Amazon
     -  Graph-Based design:
        - Web crawler
        - Google Map
        - Search Engine
        - Nearby Friends
        - News Feed

## Low-Level Design (LLDs): actual functions and code
  - Resource: Solve 1 problem every week
  - Rate Limiter
        - Token Bucket Algorithm (Example code: kode\lld\lld_test.go)
  - Meeting Scheduler
  - HashTable
  - LRU Cache

## DSA Problems
   - Resource: Solving 1-2 problems every day
      - Attempt for 30 mins to 45 mins then look for help
      - Eventually, solve all the problems without any help
   - Arrays | Strings | HashTable | Queues | LinkList | Stack | Search | Sort | Graph | Recursion | Dynamic Programming
   - DP: https://www.youtube.com/watch?v=oBt53YbR9Kk&t=6s&ab_channel=freeCodeCamp.org
   - Recursion: https://www.youtube.com/watch?v=IJDJ0kBx2LM&t=4403s&ab_channel=freeCodeCamp.org
   - Arrays common problems
     - Find the Maximum and Minimum Element in an Array-
          Write a function that returns the maximum and minimum numbers in an array.
          Reverse an Array - Implement a program to reverse an array or a string.
     - Sort an Array of 0s, 1s, and 2s - Given an array consisting only of 0s, 1s, and 2s,
         sort the array without using any sorting algorithm.(Dutch National Flag problem)
     - Find the "Kth" Max and Min Element of an Array -
        Given an array and a number k, find both the kth largest and the kth smallest elements in the array.
     - Move all the Negative Elements to One Side of the Array -
        Rearrange the array elements so that all negative numbers appear on the left, and positive numbers appear on the right.
     - Find the Duplicate in an Array of N+1 Integers -
       Assume there is exactly one duplicate number in the array, find the duplicate one.
     - Merge Two Sorted Arrays Without Using Extra Space -
        Merge two sorted arrays into a single sorted array without using extra space for another array.
     - Kadane's Algorithm - Write a function that returns the maximum sum contiguous subarray within a one-dimensional numeric array.
     - Find the Union and Intersection of Two Sorted Arrays - Given two sorted arrays, find their union and intersection.
     - Cyclically Rotate an Array by One - Given an array, cyclically rotate the array clockwise by one index.
     - 
   - String Common problems
     - Reverse a String | Check for Palindrome | Find Duplicate Characters | Anagrams | First Non-Repeated Character | Most Repeated Character | String to Integer | Longest Substring Without Repeating Characters
     - Longest Palindromic Substring | Substring Check | String Compression | Valid Parentheses
   - HashTables
       - Two Sum Problem
          (Given an array of integers and a target sum, find two numbers that add up to the target sum.)
       - Longest Consecutive Sequence
          (Given an unsorted array of integers, find the longest consecutive element sequence length.)
       - First Unique Character in a String
          (Find the first non-repeating character in a string and return its index.)
       - Subarray Sum Equals K
          (Given an array of integers and an integer k, find the total number of continuous subarrays whose sum equals k.)
       - Group Anagrams
          (Given an array of strings, group anagrams together.)
       - Contains Duplicate
          (Given an array of integers, find if the array contains any duplicates.)
       - Find All Anagrams in a String
          (Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.)
       - Longest Substring Without Repeating Characters
          (Given a string, find the length of the longest substring without repeating characters.)
       - Check if Array Pairs are Divisible by k
          (Given an array of integers and an integer k, check if it's possible to divide this array into pairs such that the sum of every pair is divisible by k.)
       - Design a Hashmap
          (Design a hashmap without using any built-in hash table libraries.)
       - Implement LRU Cache
          (Design and implement a data structure for the Least Recently Used (LRU) cache.)


## Misc
  - How you do code review, which things you consider in a review.
  - How do you mentor a junior?
  - How much unit test coverage do you have, and how do you follow TDD?
  - How much you interact with the stakeholders.

## Go Concepts
  - Resource: Go official tutorial(https://go.dev/learn/#tutorials) and then doing assignments
  - Loops: for and for range
  - Channels and Go Routines | Buffered and unbuffered channels | synchronization and locks 
  - Structs and interfaces
  - Channel patterns: fan-in/ fan-out, worker pattern
  - Packages like go-gin for REST, gorm for database interactions
  - error handling in go
  - pointers
  - slices, arrays, and maps
  - Which OOPs concepts are supported in Go
  - init function | type assertion
  - Go modules | Go sum file | packages 
  - Profiling tools: Memory and CPU profiling
  - UT and mocks: testing package and mock packages: simple tests,table-based testing, benchmark testing, etc.
## Behavioural
  - Use Star Technique: Situation, Task, Action and Result: https://www.indeed.com/career-advice/interviewing/how-to-use-the-star-interview-response-technique
  - Why do you want to join our company?
  - Explain the most challenging task you have done in your career.
  - Could you tell me a situation when you got constructive feedback?
  - Could you tell me when you had a difficult interaction with a team-mate?
  - Explain about your team and interaction with stakeholders in day-to-day tasks.
  - Could you explain your role in the team?
  - How will you approach the situation if you are interacting with other team members and the task is not done by that team member?
  - How do team members describe you, what they like about you, and what they don't like about you?
  - 
## CI/CD
  - Resource: Implementing on projects
  - Explain Continuous Integration, Continuous Deployment, and Continuous Delivery done in your project.
  - Can you describe the CI/CD pipeline and its key stages
  - What tools are you familiar with for implementing CI/CD?
  - How do you handle database schema changes in CI/CD pipelines?
  - What are some common problems in CI/CD, and how can they be mitigated
  - What is containerization, and how does it integrate with CI/CD pipelines?
  - How do you ensure security within your CI/CD pipeline?
  - What is blue-green deployment, and why is it used in CD?
  - Can you explain how monitoring and logging are handled in CI/CD environments?
  - 


## UTs and CTs
  - UTs
  - Mocking
  - Test coverage
  - TDD practice
  - 
## Microservices
  - Resource
    - Implement small projects.
  - What are microservices? What are the challenges in microservices architecture?
  - Can you explain the concept of domain-driven design (DDD) in the context of microservices?
  - How do microservices communicate with each other?
    (synchronous and asynchronous communication methods between microservices, such as RESTful APIs, messaging queues, and event streams)
  - What are some common patterns used in microservices architecture?
    (API Gateway, Circuit Breaker, Service Discovery, SAGA pattern, etc)
  - What challenges might you face while implementing microservices and how would you address them?
  - How do you handle data management in a microservices architecture?
  - What role do containers and container orchestration tools play in microservices?
  - How do you ensure the security of microservices?
  - How do you monitor and troubleshoot a microservices architecture?
  - What is the role of DevOps in microservices environments?
  - 


## Kubernetes
  - Resource:
    - Implementing K8s concepts by doing small projects
    - https://www.youtube.com/playlist?list=PLdpzxOOAlwvJdsW6A0jCz_3VaANuFMLpc
    - https://www.youtube.com/playlist?list=PLyBW7UHmEXgylLwxdVbrBQJ-fJ_jMvh8h
    - https://www.youtube.com/watch?v=EGgtJUjky8w&ab_channel=PrometheusMonitoringwithJulius%7CPromLabs
    - 
  - Kubernetes Architecture: Kubernetes cluster, nodes, pods, services, etcd, kubelet, kube-proxy, and the control plane (API server, scheduler, controller manager).
  - How Kubernetes manages containerized applications and what components are involved in deploying an application.
  - Pods and Containers
  - Services and Networking
  - The concept of services and how they enable communication between different parts of a cluster.Types of services (ClusterIP, NodePort, LoadBalancer).
  - Network policies and their role in controlling pod communication within a cluster.
  - Volumes and Storage
  - Understanding persistent volumes, persistent volume claims, and the role of the StorageClass.
  - Different types of volumes supported by Kubernetes and their use cases.
  - Deployment and Scaling
  - How to manage application deployment using Kubernetes objects like Deployments, StatefulSets, and DaemonSets.
  - Scaling applications in Kubernetes and the role of the Horizontal Pod Autoscaler.
  - Configuration and Secrets Management
   - Using ConfigMaps and Secrets to manage configuration and sensitive data in Kubernetes.
   - Best practices for securely managing secrets and configurations.
   - Kubernetes API and CLI Tools
     - Familiarity with `kubectl` commands for managing resources.
     - Understanding how to interact with the Kubernetes API.
  - Security Practices
    - Role-based access control (RBAC), Security Contexts, and Network Policies.
    - Best practices for securing a Kubernetes cluster.
  - Monitoring and Logging
   - Tools and strategies for monitoring the health of a Kubernetes cluster.
   - Understanding of logging practices within Kubernetes using tools like Fluentd, Prometheus, and Grafana.
  - Troubleshooting
    - Common issues that might arise in a Kubernetes environment and how to troubleshoot them.
    - Debugging pods and services.
  - Understanding of Helm charts for package management.
  - Kubernetes Operators and Custom Resource Definitions (CRDs).
  - Service Meshes (like Istio) and their role in a Kubernetes environment.
  - IaaC(Infrastructure as a Code) : Terraform


## Event Driven Design
  - Resource
    - Small projects
  - Kafka: https://www.youtube.com/watch?v=j4bqyAMMb7o&list=PLa7VYi0yPIH0KbnJQcMv5N9iW8HkZHztH&ab_channel=Confluent
  - rabbitmq
  - 

## Caching
  - Resource
     - Small projects
  - Redis : https://www.youtube.com/watch?v=8sHCdz_tOjk&list=PL4cUxeGkcC9h3V2eqhi8rRdIDJshP-b4P&ab_channel=NetNinja
  - MemCache

## Object-Oriented Programming and Design Patterns
  - What is an interface and how is it different from an abstract class?
  - Can you explain polymorphism with an example?
  - What are the four basic principles of Object-Oriented Programming?
    (encapsulation, inheritance, polymorphism, and abstraction)
  - How does inheritance work in OOP?
  - What is encapsulation? and What is an abstraction in OOP?
  - Can you explain the concept of a constructor? What is a destructor?
  - What are access specifiers? Provide examples.
  - What is method overriding and method overloading?
  - What is the 'this' keyword and how is it used?
  - Can you describe what composition and aggregation are in OOP?
  - What is multiple inheritance and what problems can it introduce?
  - What are virtual functions?
  - How do you handle exceptions in OOP?
  - What is the difference between deep copy and shallow copy?
  - What are SOLID principles? Can you explain the Liskov Substitution Principle?
  - What are coupling and cohesion
  - How do you ensure that a class is immutable?
  - What are design patterns? Could you give examples of some common design patterns used in your project?
  - 



## C# and .NET
  - Resource
    - Jeffrey Richter clr via c# ebook
  - Data types, variables, and operators | Control structures  | Exception handling | nullable types
  - Principles of OOP: Encapsulation, Inheritance, Polymorphism, and Abstraction | Interfaces vs Abstract classes
  - Constructors, including static constructors and private constructors
  - Overloading and overriding methods
  - Delegates, events, and lambdas
  - Extension methods | LINQ
  - Threads | Async/Await and asynchronous programming
  - Reflection
  - Differences between .NET Framework, .NET Core
  - CLR | Memory management and garbage collection
  - Unit testing with frameworks like NUnit or xUnit | Integration testing and mocking frameworks like Moq.
  - Debugging techniques and tools in Visual Studio

## Database
   - Resource:
     - Implement Small projects
   - PostGresSQL | Cassandra | MongoDB 
   - Relational database design | Normalization and denormalization | Primary, foreign keys.
   - SQL queries involving multiple tables (JOINs, subqueries, unions).
   - Aggregate functions and how to use GROUP BY and HAVING clauses.
   - SQL functions for data manipulation and conversion.
   - ACID properties (Atomicity, Consistency, Isolation, Durability).
   - Transactions, locks, and concurrency control.
   - Deadlocks
   - Indexing strategies and types of indexes.
   - Query optimization techniques.
   - Query optimizer and query execution plans.
   - Performance implications of different database designs.
   - Differences between SQL and NoSQL databases.
   - Types of NoSQL databases (document, key-value, column-family, graph) and their use cases.
   - When to choose NoSQL over a traditional relational database.
   - Methods for securing database access (authentication and authorization).
   - SQL injection and other common security threats.
   - Data encryption techniques both at rest and in transit.
   - Strategies for data backup, and types of backups (full, incremental, differential).
   - Disaster recovery plans and techniques.
   - Point-in-time recovery.
   - Cloud database services (e.g., AWS RDS, Azure SQL Database, Google Cloud SQL).
   - Scaling databases horizontally (sharding) vs. vertically.
   - Managing and maintaining databases in a cloud environment.


