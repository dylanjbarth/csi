# Lesson 2
https://github.com/fastai/fastbook/blob/master/02_production.ipynb 

https://course.fast.ai/videos/?lesson=2

Model to production 

What's deep learning actually good at? 
- Vision: detection, classification
- Text classification and translation. It's not good at conversation (yet). Not good at providing accurate information, but is good at stuff that "sounds accurate" 
- Tabular data - stuff with high cardinality (eg lots of discrete levels like zip code, product ID)
- Recommender systems: predictions don't necessarily = recommendations 
- Multi-modal: labeling, captioning (but not accurate, eg "two birds" but it's actually 3 birds.. – wonder why this is?); human in the loop type problems. 
- NLP and "protein analysis". Often about being creative with the data. 

You just have to try it and see. 

How to determine if there really is a relationship in data?

1. Pick a null hypothesis
2. Gather data of independent & dependent variables 
3. What % of the time would we see a relationship between these things? You can simulate this, but there is an equation for it – calculate the P-value (just indicating that you are toward one side of a bell curve). 

P-values are terrible – they do not measure the probability that the studied hypothesis is true. 

Drivetrain approach (example: setting insurance prices): 

1. What am I trying to achieve? (maximize 5 year profit)
2. What levers can we pull? (what price can I set)
3. Whata data can we collect? (what happens when I increase or decrease costs, eg increase premium, less costs but less customers)
4. How do levers influence the objective? via a model.

data augmentation: rotating, flipping, perspective warping, changing brightness, changing contrast, etc. This gives you lots more data and makes your model more robust. 

confusion matrix: plotting correct vs incorrect classifications 

look at top losses to see "most incorrect" or "most unsure"

tensors? seem like vector of predictions for the classes 

why pickl file as the export? is that really the standard output / serialization format for machine learned models? why?

Further reading: Building Machine Learning Powered Applications by Emmanuel Ameisen 

Watch out for data drive or out of domain data 

Questionnaire
Provide an example of where the bear classification model might work poorly in production, due to structural or style differences in the training data.

- if you were actually going to use the bear detector, you might try and deploy it in the wild (eg connected to a camera feed at a campsite that would sound an alarm if it spotted a bear) The input data there would be a video feed vs static images, which when broken into stills would probably still not align well with the training data because the pictures of bears from Bing contain mostly head on shots, less in situ, very few pictures are obstructed like they likely would be in the wild, the resolution is probably higher etc. 

Where do text models currently have a major deficiency?

- they are often inaccurate, eg they can generate realistic sounding but inaccurate text. 

What are possible negative societal implications of text generation models?

- they could be used by spammers or other bad actors to spead disinformation at scale

In situations where a model might make mistakes, and those mistakes could be harmful, what is a good alternative to automating a process?

- putting a human in the loop, AI can still make that person much more productive than they would be on their own and they can act as a safeguard / sanity check to the output of the models. 

What kind of tabular data is deep learning particularly good at?

- making predictions about data with high cardinality (high numbers of discrete categories) 

What's a key downside of directly using a deep learning model for recommendation systems?

- it's more likely to give recommendations that match previously seen behavior from similar users, vs the experience you'd get with a human expert where they would recommend something based on core characteristics and similarities of what you already like.

What are the steps of the Drivetrain Approach?

- figure out what levers you can pull, and figure out what you are actually trying to achive, figure out what data you have, build models that use the data you have to pull the levers in a way that leads you to your objective. 

How do the steps of the Drivetrain Approach map to a recommendation system?

- objective: drive sales, data: what users have bought in the past, levers: what products we show/highlight to users, model the chances that users buy something based on whether or not you show it to them. 

TODO Create an image recognition model using data you curate, and deploy it on the web.

What is DataLoaders?

- a class exported by the fastai package that wraps DataLoader classes (compatible with Pytorch dataloader) which gives you all kinds of helpers to load data, turn it into a test, train, validation sets. 

What four things do we need to tell fastai to create DataLoaders?

- the type of data, how to get it, how to label it, and how to create the validation set. 

What does the splitter parameter to DataBlock do?

- gives you control over how to split your blocks into validation set and training data. 

How do we ensure a random split always gives the same validation set?

- by setting seed, makes it deterministic

What letters are often used to signify the independent and dependent variables?

- independent = X, dependent = Y
> The independent variable is the thing we are using to make predictions from, and the dependent variable is our target. 

What's the difference between the crop, pad, and squish resize approaches? When might you choose one over the others?

- crop loses data from the edges, pad adds blank/black space, squish distorts the image.. 

What is data augmentation? Why is it needed?

- taking your initial set of training data and creating slightly distorted/modified copies / random variations of it to increase the size and variety of your training set (without changing the input so much you change the actual meaning of the data). This can help make your model more robust, especially in cases where your training data set is small. 

What is the difference between item_tfms and batch_tfms?

- item_tfms - Item Transforms - is a function to apply to each item in the input training data. 
- batch_tfms - Batch Transforms - similar idea but applying them in bulk to save time. 

What is a confusion matrix?

- a square chart comparing predictions to actual categories. This helps you identify incorrect classifications. 

What does export save?

- a pickl file which is serialized python objects. 

What is it called when we use a model for getting predictions, instead of training?

- it's called "inference" 

What are IPython widgets?

- mini-apps that can be loaded inside a notebook used for interactivity

When might you want to use CPU for deployment? When might GPU be better?

- you want to use a CPU most of the time because it's more cost effective for the workload, GPUs are only useful when you need to do many similar calculations in parallel. 

What are the downsides of deploying your app to a server, instead of to a client (or edge) device such as a phone or PC?

- the roundtrip time to make the prediction, storage and server costs. 

What are three examples of problems that could occur when rolling out a bear warning system in practice?

- basically everything comes back to the training data not matching the real data - medium differences (photos vs videos), lighting differences, in situ differences like bears hiding behind trees, etc. 

What is "out-of-domain data"?

- data that is used as input but not actually relevant to the problem that you are solving for. 

What is "domain shift"?

- the type of data the model sees changes over time so the original training is no longer relevant. 

What are the three steps in the deployment process?

- ship it with humans in the loop first, then roll out to a small group, then larger roll out. 

Further Research
Consider how the Drivetrain Approach maps to a project or problem you're interested in.
When might it be best to avoid certain types of data augmentation?
For a project you're interested in applying deep learning to, consider the thought experiment "What would happen if it went really, really well?"
Start a blog, and write your first blog post. For instance, write about what you think deep learning might be useful for in a domain you're interested in.