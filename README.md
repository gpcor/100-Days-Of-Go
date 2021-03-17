# 100 Days of Go

This is the log for my 100 days of Code using Golang. I will occasionally commit
finished projects here as proof of work. Most days will be a short log to show
proof of work.

### Day 13 (16MAR21)
I finished the "Hard mode Interfaces" assignment with relative ease tonight, and
did so without the tutorial telling me how which was really rewarding. I am 
officially on the last section of the course and I'm pumped. I've started learning
about Go Routines and Channels. Go routines seem super interesting and can really
speed up your program if used correctly. I could see how something like that could
be use on a web crawler or maybe even a discord bot to allow for running several 
different commands at a time. When I rain my example program with a Go routine,
however, it didn't output anything like it did before. From what I've learned, 
each go routine I rain for my siteRequest function hit a blocking call on the
"http.Get()" function. Likely, when each routine hit that blocking call they parent
routine simply stopped. I'm guessing I will figure out how to fix that when I get
into the next lesson, which is all about channels. 
I also learned the difference between concurrency and parallelism today, which was
a neat concept to dip my toes into. 

### Day 12 (15MAR21)
Today was all about wrapping up the lesson about interfaces. It's still a bit of
a wonky subject if I'm being honest, and in the small code examples we are making
up to explain them I have a hard time imagining what they might look like in a 
large scale application. But I can understand how they might be useful. Being able
to use one functon for multiple different objects (assuming those objects are 
similar in function) can help to cut down on excessive code. I look forward to
figuring out how to implement them more. 

### Day 11 (14MAR21)
I did two code katas today from Codewars, both were 7 kyu. I did one this morning
and one this evening. I had quite a bit of homework to work on so I didn't get to
do much, but I learned some new stuff about the Strings package which I'm always
happy about, such as strings.Replace, strings.ReplaceAll, and strings.TrimSpace.
That last one is intersting but I'm unsure how it works. I'm now working on naming
a project that I can work on for a bit. I think I'll finish the Udemy course this
week and start the project come this weekend. My only fear is I won't have much
time to code between homework assignments since I have two very project heavy classes
that are running concurrently. I'll be really glad when I'm done with school in
August. 

### Day 10 (13MAR21)
Tonight, after two very satisfying glasses of Glenfiddich 15, I sat down and decided
to code something simple without the aid of a tutorial or guide. So I made a simple
CLI calculator that, instead of user input, takes 3 command line arguments; Two
numbers and a string for add/subtract/multiply/divide. The numbers are converted
to Float64 and are passed to their appropriate operation function. The result is
simply return to back to the command line. It's not much, but I was proud that
I was able to figure this out. Even for the command line arguments and ParseFloat
I simply read the Golang documentation. I'm satisfied with the effort and it has 
motivated me to strike out on my own and build more programs on my own. 

### Day 9 (12MAR21)
I missed yesterdays programming session. I made a promise to the wife that I'd
stay off the computer for the afternoon since she had to go back to Memphis today,
so I stayed off and didn't code. I thought I'd be more upset about it, but it
was more than worth it. Family time is important. 

I spent some time in the Udemy course tonight learning about interfaces. It seems
like a pretty neat type that can be used to group different types together so that
they can make use of common functions, rather than writing out a different function
for every single type. I'll need more practice with it but it seems very interesting.

### Day 8 (10MAR21)
I spread my programming practice throughout the day when I'd have breaks at work. 
This seemed to be better than trying to cram my practice in to the last 2 hours
of the day. I struggled really bad with a 6 kyu problem that left me a little
disheartened, but I learned quite a bit from the solutions so I was able to
take that knowledge and apply it to a less complicated problem. I ended the night
with a better understanding of Strings and Bytes and how for loops can iterate
through the characters, as well as UTF-8 encoding (sorta). Tomorrow I want to do
more of my Udemy course since the next section is about interfaces.

### Day 7 (9MAR21)
Another night of code katas, another night of not feeling like I got anywhere.
It sucks when you feel like you're doing good in a tutorial and like things are
clicking, but then you put it into practice (on something like a code kata) and 
it doesn't seem to go as planned. That said, I'm only a week into learning Go and
taking this all seriously, so it's to be expected. No one ever got good in a week.
In the words of Cal Newport from "So Good they Can't Ignore You", you have to have
a Craftman's Mindset if you are going to build any usable career capital. Deliberate
practice and good mental strain are the only path to Mastery.

### Day 6 (8MAR2021)
Worked on a code kata that I couldn't wrap my head around. I seem to be missing
something in my fundamentals, although I think this is in part due to the time 
when I am trying to code. I need to reverse my schedule. Get up earlier, code 
before work, and then I can fit in any code practice I may want in the evening if
I can. I don't have a lot of attention economy left at this time of night.

### Day 5 (7MAR2021)
Didn't get much done today. Just tried some stuff on my own regarding the card
program. Had to do some homework and offline stuff. Now I gotta go to the airport.
Hoping for more work tomorrow.

### Day 4 (6MAR2021)
This commit is a tiny bit late because I was distracted. But I did code today.
I did more Udemy class and learned a lot more about Structs, Pointers, and Maps.
I managed to figure out how to iterate over a class before the instructor taught
it so that was cool. With each new thing I can see how a Blackjack program could
be built. I've written out the logic for it. I plan to give it a whirl soon.

### Day 3 (5MAR2021)
I did more of my Udemy course (see Day 1 for link). I learned mostly about the
structs data type (Feels like fancy Python dicts tbh but I dig it) and I've
started going over Pointers. That's gonna be a new and intersting one.

### Day 2 (4MAR2021)
I did a couple of Katas from CodeWars.com
I feel more dumb than when I initiall started. 
I don't know enough to always get a solution, but the ones I see always make sense.

### Day 1 (3MAR2021)
Still working through Go Udemy course, found :
https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7797306#overview 
Completed up through lesson 31. Still confused about when to use Receivers
