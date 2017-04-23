

**Execution flow of invoke-transaction**:


1. Parameters are read from parameters.json (Modify only "invoke_transaction") and stored in variable args.
2. Call to the method invokeTransaction() in e2eUtils.js is made.
3. The method performs chain.sendTransactionProposal and based upon response performs chain.sendTransaction
4. Final response that the transaction has been committed to ledger is printed in the terminal.


**Execution flow of query-transaction**:

1. Parameters are read from parameters.json(Modify only "query_transaction") and stored in variable args.
2. Call to the method queryTransaction() in e2eUtils.js is made.
3. The method performs chain.queryByChaincode and gets a response which is printed to terminal.
4. For informational purposes the height of the block chain and some details of the latest block are printed.

**Execution flow of query-and-execute**:

1. Parameters are read from parameters.json(Modify "query_transaction" and "query_and_execute" so that the Ids match) and stored.
2. Call to the method queryChaincodeAndExecuteTask() in e2eUtils.js is made.
3. The method performs chain.queryByChaincode and gets a response which is printed to terminal.
4. The status of the order retrieved is checked. If it is not delivered, then a background task is executed (simulation of an external api being called to perform a task)
5. The background process performs it's tasks and calls invoke-updateStatus.js once it is done to change the status of the purchase order to 'Delivered'.
6. A query is run again using query.js to show that the value has been changed.

**Execution flow of register-User**:

1. The username to be registered is read from parameters.json(Modify "registerUser")
2. Call to method getRegistrar from util.js
3. Utilizes admin, adminpw values to get enrollment certificates of admin to register a new user
4. Calls fabric-ca methods register and enroll and sets the user context for the client.
5. New user only has access to querying and not creating or deleting assets.
