## reducer package

This package holds the description and the data for reudcing the data coming into a stream and send the data thorugh a RPC to a key-value handler and which reduces the data stream into key-value store.


1. The mapreduce program takes the data present in the *.urb files and process them to put them into seperate key-value pairs seperated by frequency and send them fast.


2. PROTOCOL:
   
   COMMAND NAME         DESCRIPTION
1. ADD                  ADD OBJECT TO THE STAGING TABLE BEFORE BEING COMMITED.
2. REMOVE               REMOVE OBJECT FROM THE STAGING TABLE.
3. COMMIT               MOVE FROM STAGING TABLE TO THE REMOTE SERVER (HANDLED BY A AUTOMATION SERVER LIKE JENKINS).
4. VIEW                 VIEW THE UPLOADED OBJECT UPLOADED ON  THE STAGING TABLE.


