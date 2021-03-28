# Binary representations of data

goal of the class: pick the right data types when you are coding and know their limitations. 

- ask about the binary adder for the next checkin 

signing bit exists in floating point

disadvantage with integers is that we waste some space and we have negative 0 which make comparing values harder and breaks addition. 

twos complement: 
- only single value of 0
- scheme where the adder we use works as expected. 

in 4 bits we can represent -8 => 7

starting from the highest order bit and working upward. 

so 1000 is -8 and 1001 is 7

what is bias? a way of thinking about signed integers - this is the most negative possible value – aka offset, correction, adjustment. 

