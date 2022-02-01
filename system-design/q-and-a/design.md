# Problem Statement

Design website for people to post questions and answers. People should be able to upvote and downvote. We want to be able to see recent questions and popular questions. 

# Functional requirements

- must be able to post questions and answers to questions
- should be able to upvote and downvote
- must be able to view recent questions, and popular questions
- (bonus): support reddit-style editing and deleting questions / answers.

# Non-functional requirements

- minimize latency, maximize availability 
- sensisible consistency guarantees for usability purposes (eg users should read their own writes) 
- userbase is restricted to single geographic region. 

# Out of scope: 
- comments, tags, badges, hot questions, search

# SLOs

- 99.9% of the time, question and answer reads and writes should complete within 1 second. 
- 99.9% of the time, submitted answers should be readable by all users within 5 seconds. 
- 99.9999% availability, ~60 minutes of downtime per year. [Ref](https://en.wikipedia.org/wiki/High_availability#Percentage_calculation) 

# System Data

Our two main objects are going to be questions and answers. We will think through the schema below, but to get a rough sense of load we can assume that the bulk of the size will come in the qustion and answer bodies. Let's assume we limit these to be 5k characters in length, and thus assume the average size of these will fall around 5k because most will be smaller and we can account for around another 1k of metadata like the creator ID, summary, created_on and edited_on timestamps, etc. 

## question

```
created_by uuid fk 
title varchar 200
body varchar 5000
created_on timestamp
updated_on timestamp
```

## answer
```
question_id 
created_by
body varchar 5000
created_on timestamp
updated_on timestamp
```

## Load Estimation 

|Resource|Per year|Per Day|Per Second|
|-------------|------------|----------|----------------|
|Questions|5e6|13.6e3|60|
|Answers|20e6|54.8e3|236|
|Views|500e6|1.3e6|~6e3|

Other info:
- 10 M users total.
- Expect that the scale of the system may continue to increase (in terms of volume of questions, answers, and views.)

Constraints: 
- impose a reasonable limit on the question and answer size. 


Out of scope

Bonus: 
- how would you do this on your own hardware? 

SLOs
- answers should be instant ;) questions should be fetched within a second or two. after posting an answer it should be available to others within 5 seconds. would be nice for users to read their own writes. 
- no downtime allowed ;)

# Functional Requirements

# Non-Functional Requirements

# Table of Contents

- Problem Statement
- Functional Requirements
- Non-functional Requirements 
- SLOs
- Load estimations & guarantees
- API design
- Storage choice justification + data model (ERD)
- Infrastructure diagram
- Request Lifecycle 