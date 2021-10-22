# Chapter 4
https://github.com/fastai/fastbook/blob/master/04_mnist_basics.ipynb
https://course.fast.ai/videos/?lesson=3 (start at ~1:05)

MNIST - handwritten digits, Yann Lecun 1998, Lenet 5, first system to demonstrate useful hand-written digit sequences. 

Start simple and scale. 

Exercise: starting with a 2d array of uint8s from 0-255 (representing a color gradient from white to black) how would a computer be able to distinguish between array representing 3 vs 7? 

Ideas: it seems obvious that most 3s are going to have their centers shaded in as well as the bottom and top of the number curving inward. We could weigh darkness in these regions heavier. For sevens, it would be straight lines down or diagonal left. So basically creating regions of pixels that "count" toward a number and whichever has the highest score is our prediction. Another way we could do it is extending that pixel idea but zooming out a bit as well, so the combination of having darkness in two areas in different regions would be an even heavier weight. Etc. 

This exercise is "creating a baseline" - what is the simplest way to solve this problem. Eg Jeremy's idea here is to just take the average of all the categories and then compare to make predictions (simpler than my idea above). 

*L1 norm* is the average of the absolute value of differences. 
*L2 norm* is the square root of the average of the sum of the squares. (MSE - mean squared error)

Both cases force the differences to be positive, L2 norm basically just penalizes larger differences more. 

*Broadcasting*: feature of PyTorth that allows expanding the smaller element to match the size of the larger element. 

Loops make code harder to read but importantly much slower on the GPU, so take advantage of broadcasting in PyTorch! 

## Stochastic Gradient Descent 

basically the idea I had above :) come up with weights based on pixel darkness in certain areas. 

The basic algorithm is: 

1. Initialize weights randomly 
2. Make a prediction
3. Calculate the loss (how good the model is)
4. Calculate the gradient (how changing the weight would change the loss)
5. Step (change the weight based on the gradient)
6. Repeat or stop if model is good enough. 

Picking your learning rate correctly is important to ensure that you are able to solve the problem in a reasonable amount of time. 

Important jargon: 

tensor: n-dimensional array
rank: the number of axes or dimensions in a tensor, eg a 3d tensor is a rank-3 tensor
shape: is the size of each axis of a tensor. eg 6000 images, each 28x28 pixels

Important libraries to get comfortable with: 

- numpy
- pandas
- pytorch (understand the FFI to C and how some of these operations are implemented. Could be a great series of blog posts.)
- fastai

CUDA: the equivalent of C on a GPU
