# basic-blockchain-app

**Update**
Works with hyperledger/fabric commit 118f82f04a8f3227480904ea6b7c3ed961c0d8c1.
Works with hyperledger/fabric-ca commit 855036cc931b20c0019f9f67e4b7f8acb9ec868d.

Blockchain application developed on Hyperledger Fabric using the Fabric-Node-SDK. Simple modifications to the end to end example provided to exercise the entire functionality of SDK.

Instructions to run:

1. Download hyperledger/fabric and hyperledger/fabric-ca to /home/$USER/directory_name/src/github.com/
2. Ensure that export GOPATH=/home/$USER/directory_name
3. Go to fabric directory and run 'make docker' and then 'make couchdb'.
4. Go to fabric-ca directory and run 'make docker'.
5. Go to the directory you cloned this repository and run:
	a. docker-compose up -d
	b. export GOPATH=$PWD && npm install
	c. node create-channel.js
	d. node join-channel.js
	e. node install-chaincode.js
	f. node instantiate-chaincode.js
	g. Go to e2eutils.js and modify the invokeTransaction parameters to any of the ones mentioned below.
	h. node invoke-transaction.js
	i. Go to e2eutils.js and modify the queryTransaction parameters to any of the ones mentioned below.
	j. node query.js


Invoke Transaction parameters:
1. "create", "purchase_order_number" (Example : ['create','PO156897'])
2. "delete", "purchase_order_number" (Example : ['delete','PO156897'])
3. "updateStatus", "purchase_order_number", "status_update" (Example : ['updateStatus','PO156897','Delivered'])
4. "updateQuantity", "purchase_order_number", "new_quantity" (Example : ['updateQuantity','PO156897','150'])
5. "updateCustomer", "purchase_order_number", "customer_name" (Example : ['updateStatus','PO156897','Company XYZ'])
6. "updateSupplier", "purchase_order_number", "supplier_name" (Example : ['updateStatus','PO156897','Company ABC'])
7. "updatePartName", "purchase_order_number", "part_name" (Example : ['updateStatus','PO156897','Razor Gaming Keyboard'])


Query Transaction parameters:
1. "queryPO", "purchase_order_number" (Example : ['queryPO','PO156897'])
2. "queryPOIds", "" (Example : ['queryPOIds'])

If there are any issues it's usually because the development environment is not setup properly. Make sure to follow all the setup instructions only, from https://github.com/hyperledger/fabric-sdk-node and the 'Prerequisites' from http://hyperledger-fabric.readthedocs.io/en/latest/dev-setup/devenv.html?highlight=development.


