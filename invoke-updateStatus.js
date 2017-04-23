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

var tape = require('tape');
var _test = require('tape-promise');
var test = _test(tape);
var e2eUtils = require('./e2eUtils.js');

var parameters = require('./parameters.json');
var testUtil = require('./util.js');
var args = testUtil.getArgs(parameters.query_and_execute.updateStatus);


test('\n\n***** End-to-end flow: invoke transaction to move money *****', (t) => {
	e2eUtils.invokeChaincode('org2', parameters.properties.chaincodeVersion, t, args)
	.then((result) => {
		if(result){
			t.pass('Successfully invoke transaction chaincode on channel');
			const spawn = require('child_process').spawn;
			spawn('node', ['query.js'], {
			  //stdio: [ 'ignore', out, err ],
			  stdio: 'inherit',
			  detached: true
			}).unref();
			t.end();
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
