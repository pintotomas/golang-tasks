# Task 1
## Problem description

In a project dedicated to storing live event data and processing it for metrics analysis, our system initially received incoming data in XML format from our legacy infrastructure whenever a user interacted directly with a live event on their website. 

Subsequently, a client requested direct interaction with our API to upload  live events information not created through the user interface. However, the data that arrived to our endpoint was in JSON format, containing additional fields unnecessary for our purposes, with variations in field names and not mapping 1:1 to the fields we already had in our existing data structure. 

For better understanding of where we were standing when this problem arises, I represented it with the following diagram:

![alt text.](/diagrams/task1/previous.png "Previous solution.")

The problem we had was that we were using our original model from the legacy system to feed our database model directly, and we had to add support to the data we received by the client, preferably maintaining their structure for code clarity and traceability

To address this challenge without disrupting too much other parts of the code already relying on the original structure, we implemented the adapter pattern and used an interface to access the live event data. So what we did was to replace the services that in the first place used the original object, to use an interface which the new adapter would implement. 

## Solution

A diagram representing the solution:


![alt text.](/diagrams/task1/new.png "New solution.")

As you can see, we have the 'Client Live Event' which is the adaptee, the Live event Adapter and the Live Event Interface which is the target.

The incoming data from the client had some differences:

* 'Title' field was named 'Name' in the client data
* The description was split in two fields: Description and Additional Information
* The timezone was missing (We used UTC by default)
* Instead of having the start and end date, we had the start date and the duration
* We had some extra fields like ClientTraceID we didnt require but could come handy for logging purposes for example

