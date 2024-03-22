# Task 1
## Problem description

In a project dedicated to storing live event data and processing it for metrics analysis, our system initially received incoming data in XML format from our legacy infrastructure whenever a user interacted directly with a live event on their website. 

Subsequently, a client requested direct interaction with our API to upload  live events information not created through the user interface. However, the data that arrived to our endpoint was in JSON format, containing additional fields unnecessary for our purposes, with variations in field names and not mapping 1:1 to the fields we already had in our existing data structure. 

For better understanding of where we were standing when this problem arises, I represented it with the following diagram:

![alt text.](/diagrams/task1/previous.png "Previous solution.")

The problem we had was that we were using our original model from the legacy system to feed our database model directly, and we had to add support to the data we received by the client, preferably maintaining their structure for code clarity and traceability

To address this challenge without disrupting too much other parts of the code already relying on the original structure, we implemented the adapter pattern and used an interface to access the live event data. So what we did was to replace the services that in the first place used the original object, to use an interface which the new adapter would implement. 

To see the state previous to applying the Adapter pattern, you can check [this commit](https://github.com/pintotomas/golang-tasks/commit/89dd356e639fdf5ee11cb10b3d811ee09074a641)

## Solution

A diagram representing the solution:

![alt text.](/diagrams/task1/new.png "New solution.")

As you can see, we have the 'Client Live Event' which is the adaptee, the Live event Adapter and the Live Event Interface which is the target.

The incoming data from the client had some differences:

* 'Title' field was named 'Name' in the client data
* The description was split in two fields: Description and Additional Information
* The timezone was missing (We used UTC by default)
* Instead of having the start and end date, we had the start date and the duration
* We had some extra fields like ClientTraceID we didn't require but could come handy for logging purposes for example

### Notes about the solution

* I am only implementing a storage service (without any connection to a db or any other storage) for keeping the task simple, other services can interact with this interface in the same way
* For the same reason, when I run the task I simulate the json/xml request bodies to avoid the overhead of creating an API and only showcase the part of the problem related to the task

### What could be improved?

* We could aswell have an adapter for the original LiveEvent struct. Although It's unlikely to change since it comes from legacy code nobody really wants to touch, it's still a possibility and by having an adapter for this one we could save time in the future 
* Maybe we don't want to expose the whole Live Event interface to all services and that could lead to multiple interfaces and adapters
* To keep it simple I used the standard library for tests, but I recommend using testify for more useful tools for testing. 
* Similar with logging, I'm using the fmt package for some basic outputs, but I would use the log package from the standard library.

# Task 2
## Problem description

During the development of a microservice aimed at managing various AWS resources, we encountered a challenge involving the AWS API rate limit quota. 
We faced limitations while using multiple workers concurrently within separate goroutines for the creation of AWS MediaPackage resources, among other requests we sent.
For example, if our account was limited to 10 requests per second, AWS would penalize for example for the next 5 seconds, making all the subsequent requests fail

To simplify the problem for the task, I will use an API limiter instead of interacting with AWS, and I will have 3 workers which will attempt to do the requests that they receive in a channel

## Solution

To solve the issue of requests being throttled by the limiter, I will implement a backoff once I get a request limited, and I will send a signal to the other workers to avoid sending more requests and to backoff for a few seconds

This solution was challenging and quite complex, but as you can see on the results we have way less requests failing. Previously, it was succeeding 20 out of 50 times, and now it succeeds 40 times

For the worker I used reading only channels for both the requests and its own backoff signal channel. And I created an array of send only channels to notify the other workers

Previously we were only ranging on the requests channel, and now we have multiple channels, so we should be careful with potential deadlocks, for that reason I add default cases in the selects to prevent deadlocking

The most important one and that was hard to detect at first, was the non-blocking send operation to notify other workers to backoff. If we simply tried to send the notification without the select-case, it could lead to a deadlock because other go routines could be attempting to notify the same worker  

### Notes about the solution

* It will take a few seconds to run (< 10)
* Once more, for simplicity I'm not sending any request or interacting with any real service, so I can locally simulate the "request rejection" problem
* There could be other reasons for a request to be rejected, for example a bad request or authentication issues, but in this solution I'm assuming the only problem is the quota limit

### What could be improved?

* A retry can be added after each back off, so we run all requests successfully
