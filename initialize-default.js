/**
 * Copyright 2017 IBM All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

// This is an end-to-end test that focuses on exercising all parts of the fabric APIs
// in a happy-path scenario
'use strict';
var utils = require('fabric-client/lib/utils.js');
//utils.setConfigSetting('hfc-logging', '{"debug":"console"}');
var logger = utils.getLogger('E2E testing');
var tape = require('tape');
var _test = require('tape-promise');
var test = _test(tape);
var e2eUtils = require('./e2eUtils.js');
var orders = require('./initialscript/initial_po.json');
var parameters = require('./parameters.json');
var testUtil = require('./util.js');


var async = require('async');

var q = async.queue(function(i, callback) {

  var args = [];
  args.push('createCompletePO');
  var PoId = orders.test_data[i].PoId;
  args.push(PoId);
  var Quantity = orders.test_data[i].Quantity;
  args.push(Quantity);
  var Part_Name = orders.test_data[i].Part_Name;
  args.push(Part_Name);
  var Customer = orders.test_data[i].Customer;
  args.push(Customer);
  var Supplier = orders.test_data[i].Supplier;
  args.push(Supplier);
  var Status = orders.test_data[i].Status;
  args.push(Status);
  var Price = orders.test_data[i].Price;
  args.push(Price);
  logger.info("the args are "+args);

  test('\n\n***** End-to-end flow: invoke transaction *****', (t) => {

      e2eUtils.invokeChaincode('org1', parameters.properties.chaincodeVersion, t, args)
    	.then((result) => {
    		if(result){
    			t.pass('Successfully added default purchase order'+i+' on channel');
          //logger.info("Waiting for 5 seconds before next transaction");
    			t.end();
          callback();
    		}
    		else {
    			t.fail('Failed to invoke transaction chaincode ');
    			t.end();
    		}
    	}, (err) => {
    		t.fail('Failed to invoke transaction chaincode on channel. ' + err.stack ? err.stack : err);
    		t.end();
    	}).catch((err) => {
    		t.fail('Test failed due to unexpected reasons. ' + err.stack ? err.stack : err);
    		t.end();
    	});
  });
}, 1);

// assign a callback
q.drain = function() {
    console.log('all items have been processed');
};



for(let i=0;i<orders.test_data.length;i++) {

q.push(i, function (err) {
    console.log('finished processing '+i);
});
}
