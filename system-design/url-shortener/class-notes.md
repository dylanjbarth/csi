Things to focus on for next time: 
- can you walkthrough lifecycle of a particular request?
- what features / SLOs are you supporting / not supporting and what tradeoffs are possible? recommendations? 
- how is data represented on physical storage? specific storage system? schema or KV data, indexes? 
- Biggest risks? Unknowns?
- Cost estimate

Learnings from reviewing peer designs: 
- think about stuff like IOPS, round trip request time, etc in justifying decision to break out into multiple nodes. 

QA Site Requirements: 

Questions I would ask (before hearing what peers asked in the recording):

- what is the nature of a question that can be asked? is it multi-media or text only? 
- how many questions get asked per day/ per year? 
- how many answers per question
- how many times is a question viewed on average?
- distribution of question / answer popularity (guessing some questions are way more interesting/popular, eg maybe 80% of views are for 20% of questions)
- do we need to think through the logic such as marking questions as answered, raking via upvoting / downvoting? authentication? security? rate limiting? 

