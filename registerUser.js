'use strict';

var utils = require('fabric-client/lib/utils.js');
//utils.setConfigSetting('hfc-logging', '{"debug":"console"}');
var logger = utils.getLogger('E2E create-channel');

var tape = require('tape');
var _test = require('tape-promise');
var test = _test(tape);

var hfc = require('fabric-client');
var util = require('util');
var fs = require('fs');
var path = require('path');

var testUtil = require('./util.js');
var parameters = require('./parameters.json');

var the_user = null;

hfc.addConfigFile(path.join(__dirname, './config.json'));
var ORGS = hfc.getConfigSetting('test-network');

//
//Attempt to send a request to the orderer with the sendCreateChain method
//
test('\n\n***** End-to-end flow: create channel *****\n\n', function(t) {
	//
	// Create and configure the test chain
	//
	var client = new hfc();

	var caRootsPath = ORGS.orderer.tls_cacerts;
	let data = fs.readFileSync(path.join(__dirname, caRootsPath));
	let caroots = Buffer.from(data).toString();

	var orderer = client.newOrderer(
		ORGS.orderer.url,
		{
			'pem': caroots,
			'ssl-target-name-override': ORGS.orderer['server-hostname']
		}
	);

	// Acting as a client in org1 when creating the channel
	var org = ORGS.org1.name;

	utils.setConfigSetting('key-value-store', 'fabric-client/lib/impl/FileKeyValueStore.js');
	return hfc.newDefaultKeyValueStore({
		path: testUtil.storePathForOrg(org)
	}).then((store) => {
		client.setStateStore(store);
		return testUtil.getRegistrar(parameters.registerUser,client, t,false, 'org1');
	})
	.then((user) => {
		t.pass('Successfully enrolled user \'sujay\'');
  }, (err) => {
		t.fail('Failed due to error: ' + err.stack ? err.stack : err);
		t.end();
	});
});
