# Solution : Kindly Read Me!!!

Really enjoyed working on the project. It was an interesting assessment to take. Given the fact that I had just started with GoLang and had written my first "Hello World" in Go just a week ago. after receiving this, I have learnt a lot in a week. Lets jump right into the solution. The entire project is designed with keeping modularity in mind. Also to enable faster changes and more controll in the hands of business/Analyst teams.

# Components

- Pool of checks
- Strategy Loader and Strategy File
- Check Constructor
- API Controllers
- Cache
- Properties File

## Pool Of Checks

Every eligibility check is implemented as a standone check. Thus all modular checks are kept under one package (hence Pool of checks) and implement a common interface. Creating a pool of checks has two distinct advantages

- **Code Reusability**
- **Mix-N-Match**

The attributes/parameters of these checks


> age in case of AgeCheck, income in case of IncomeCheck

are configurable and thus each check in this pool just holds the logical part of it, logic around threshold.  The threashold of that logic is completely configurable.

## Strategy File


A **Strategy** is a set of rules/checks that determine the eligibility of the customer. There can be more than one strategies. Each strategy contains a set of checks from the **pool of checks** along with the configurable thresholds . A strategy is defined in a file known as **Strategy File**.
in our solution we have just one strategy V1.
**It is encouraged to create more than one strategy of diffeerent checks or paramter thresholds to undertand the rejection/acceptance rate. Also this promotes A/B Testing.**

V1.json

    {
    "IncomeCheck": {
    "Income": "100000"},
    "AgeCheck": {
    "Age": "18"},
    "CreditRiskCheck": {
    "NoOfCards": "3"},
    "PoliticallyExposed": {
    "Exposed": "false"},
    "AcceptedAreaCodes": {
    "AreaCodes": "0,2,5,8"}
    }
It is not necessary to maintain these strategies in a file, we can put these up in a Document Oriented Db as well and read it from there. For the sake and simplicity of this solution I have decided to read from a file.

## Strategy Loader

As the name suggests, The task of the Strategy loader is to load a strategy based on the active version of strategy which is passed to it. This active version of strategy is configured in properties file. **The Strategy loader then informs the Check Constructor, what are all the check objects in needs to create.**
This happens only once during the system startup time.

## Check Constructor

Every Check present in the system has to be configured with Check Constructor. **The task of this component is to create objects at runime and to initialize there threshold parameters as well**. It is fed by the Strategy Loader about the objects to create.Again Once  these objects are created at the system startup, we keep using them untill and unless there is a change in the strategy, which would anyway require an application restart.

## API Controllers

To Interact with the Decision Engine and to meet the business requirements, there are 2 controllers provided.
1.

`POST /process`

Body:

`{
"income": 213383,
"number_of_credit_cards": 3,
"age": 36,
"politically_exposed": false,
"job_industry_code": "10-200 - Louvers and Vents",
"phone_number": "114-183-5760"
}`

This is the main controller  which takes in customer paramters and returns the application status.

2.

`PATCH /ApproveList`

Controller to update the list of Pre-Approved phone numbers and add to it for overriding decision

`DELETE /ApproveList`

Controller to update the list of Pre-Approved phone numbers and add to remove from it
The body of both these controllers is same

`{"phone_numbers":["41-759-8127","114-183-5760"]}`

You can add or remove multiple phone numbers in one go from the list.

Regarding the security of this api, Along with  authenticating it using an authentication token, we can also configure the Gateway server using ngnix.conf file to allow access to this api only for requests comming from within VPC
# Cache

To hold the list of pre-approved phone numbers and to provide the ability to update it at runtime without restarting the server, we are using an in-memory cache. This is thread safe implementation of map using RW locks to support Concurrency. We could have used a DB or a third party in-memory cache, but being a new gopher , I decided to keep it simple. Plus it does the trick.


## Properties File

Used to determin the port No and active strategy version as explained in the Strategy section. 

# IMP INSTRUCTION FOR STARTUP
It is crucial for the application to have a well formated strategy file with checks defined in pool of checks for the application to startup.

Also Kinldy determine the active strategy version in the properties file

