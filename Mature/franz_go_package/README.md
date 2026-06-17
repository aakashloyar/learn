# Kafka

-> We will learn about Kafka in this module

# client 
-> by this we talk to kafka
cl, err := kgo.NewClient(
    kgo.SeedBrokers(
        "localhost:9092", // Broker 1
        "localhost:9093", // Broker 2
        "localhost:9094", // Broker 3
    ),
)
if err != nil {
    panic(err)
}
defer cl.Close()

# 3 terminologies
1. Broker
2. Topic 
3. Partition

1. Broker 
-> Server on which it runs
-> broker can be multiple based on you give
-> For Brokers declaration please refer to client declaration

2. Topic
-> Basically this will tell this message belong to which category
-> suppose order and payment 
-> these are different and you want separated queue for this then you will use topic
-> with each message you must send topic

3. Partition
-> Basically to make kafka fast 
-> the topic is divided into partitions
-> each partition you can assume as queue
-> this will help in parallel execution

# Replica
-> now each partition will have a Leader broker 
-> supppose that broker down then how to handle that
-> so here comes replica that makes other broker to store the copies of partition
-> suppose leader dies then another will take it's place

# clientID 
-> Basically which service
-> helps to debug which service produce or consume 

# Consumer Group
-> suppose i have 2 consumers and they have same consumergroup then they will split the partition between themselves that i will execute this partition and you execute this one
-> they help use to scale by multiple consumers
-> this is how to declare it in the code
kgo.SeedBrokers("b0", "b1", "b2")
kgo.ConsumerGroup("order-service-group")
kgo.ConsumeTopics("orders", "payments")


# About 
1. Kafka is not FIFO
-> As is partitioned for parallel processing for scalability so not possible
-> But inside partitioned it is FIFO

2. Kafka is not delete after read 
-> as suppose another service analytics and log service need this
-> then deletion is not the best practice after read by 1 service
-> basically it works using offset
-> it works like the last executed data then move to next using offset
-> But data must removed after read
-> so 2 strategies TTL vs Data Limit
-> st1. Data Limit executed data must not exceeds 100GB 
-> But not a nice idea as suppose heavy load then without executing it is removed
-> But we can manually set the limit according to load
-> Not a good idea
-> st2. Time -> 7days
-> but suppose it is not executed till now and removed then 
-> but it has very less chance so we should opt this

# Strategies
-> 1 cluster vs multiple clusters
-> 1 cluster is fine as it is easy for devops
-> multiple are required suppose payment is critical service and we want separate it from other 


