'use strict';

var sleep = require('system-sleep');

for(let i=0;i<10000;i++) {
  if(i%1000 == 0) {
    console.log("Process is executing in background"+(i/1000));
    sleep(2000);
  }
}

const spawn = require('child_process').spawn;
spawn('node', ['invoke-updateStatus.js'], {
  //stdio: [ 'ignore', out, err ],
  stdio: 'inherit',
  detached: true
}).unref();
