# Lesson 1
Deep learning has many applications: 

![dl-applications](./dl-applications.png)

## history 

1943 - Warren McCulloch and Walter Pitts – mathematical model of artificial neuron 
1950 - Frank Rosenblatt - first perceptron 

First "AI Winter" - Minsky and Papert wrote Perceptrons. Showed a single layer of perceptrons weren't able to learn simple math functions – also showed that multiple layers could learn this but people ignored this. 

In 1986, MIT - Parallel Distributed Processing - PDP 

![pdp](./pdp.png)

In 80s, Proved that second layer could actually allow any mathematical model to be approximated, but this is too big and slow. The key is to use even more layers of neurons (made possible by more data and better computer hardware). 

## Class principles

David Perkins - best way to learn is to start by "playing the whole game" - develop a sense of understanding of the whole piece. Ie top down vs bottom up – starting with context. Next: make the game worth playing. Have a competition, keep score, motivate the problem. 

## Tools

Tensorflow – everybody was using this but got bogged down, nobody using it anymore
PyTorch - easier to use, better for researchers 20% => 80% of papers. Missing higher level APIs. 
fastai - sits on top of PyTorch

## What's happening in the first exercise
what is machine learning? a way to get a computer to complete a task when the steps aren't clear, eg recognizing objects in photos. 
Arthur Samuel - 1949 - need to spell out every minute step of each process is intractable – instead give it examples and let it figure out how to solve it itself. Built a checkers bot that beat the state champ of Connecticut. 

process for training a machine learning model = inputs + weights => model => results => measure performance and update weights 

**machine learning**: training of programs developed by allowing a computer to learn from its experience rather than through manually coding the individual steps. 

**universal approximation theorem**: neural network can solve any problem to any level of accuracy (in theory)

how to find the right weights / train the NN? 

**stochastic gradient descent / SGD** - 

model = architecture (structure of the function to parameterize)

weights = parameters to the function 

results = predictions based on independent data

loss = performance

Limitations: 
- you need data to train the model 
- model can only learn patterns based on what you show it 
- data must be labeled 

